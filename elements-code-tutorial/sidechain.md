---
layout: page
title: Elements Sidechain
permalink: /elements-code-tutorial/sidechain
redirect_from:
  - /sidechains/creating-your-own/
  - /sidechains/creating-your-own
---

# Elements code tutorial

## Elements as a Sidechain - Federated Two-Way Peg

Elements is a general purpose blockchain platform that can also be "pegged" to an existing blockchain (such as Bitcoin) in order to enable the two way transfer of assets from one chain to the other. Implementing Elements as a sidechain allows you to work around some of the inherent properties of the main chain, while retaining a good degree of the security provided by assets secured on the main chain.

While a sidechain is aware of the main chain and its transaction history, the main chain has no awareness of the sidechain and none is required for its operation. This enables sidechains to innovate without restriction or the delays associated with main chain protocol improvement proposals. Rather than trying to alter it directly, extending the main protocol allows the main chain itself to remain secure and specialised, underpinning the smooth operation of the sidechain.

An example of an Elements based sidechain in production use is Liquid. 

Liquid is an implementation of a federated sidechain, a blockchain with different features, capabilities, and benefits than the main Bitcoin blockchain. 

By extending the functionality of Bitcoin and leveraging its underlying security, the Liquid network is able to provide new functionality that was previously not available to Bitcoin users.

Liquid was designed to address the needs of exchanges, brokers and traders and enables the rapid, confidential and secure transfer of funds between participants and providing a solution to the inherent problem of delayed transaction finality on the Bitcoin network.

Every Liquid bitcoin held within the sidechain is pegged to bitcoin on the main chain using a Federated 2-Way Peg. This allows bitcoin to be deposited in the sidechain whilst retaining the ability to withdraw assets back to the main Bitcoin blockchain. Apart from peg-out, everything related to running Elements as a sidechain can be done using the Elements daemon and client. This is because peg-out requires the actions of 'Watchmen' to control the multisignature release of funds on the Bitcoin blockchain, the use of which is not covered in this tutorial.

In this section we will look at how to send bitcoin from our bitcoin regtest chain into our Elements blockchain, which will be operating as a sidechain.

First we need to wipe out the chain and wallet to get started with a new genesis block:

~~~~
rm -r ~/elementsdir1/elementsregtest/blocks
rm -r ~/elementsdir1/elementsregtest/chainstate
rm ~/elementsdir1/elementsregtest/wallets/wallet.dat
rm -r ~/elementsdir2/elementsregtest/blocks
rm -r ~/elementsdir2/elementsregtest/chainstate
rm ~/elementsdir2/elementsregtest/wallets/wallet.dat
~~~~

We will not require "n of m" block signing in this section so that we can keep the code brief, instead we will return to using "OP_TRUE" block creation. If you implement your own federated sidechain you can of course use the "n of m" signing method outlined earlier in the tutorial.

In order to enable peg-in and peg-out we need to pass a valid "fedpegscript" to our node as a startup parameter. Alternatively, this can be set within the config file of each node. 

~~~~
FEDPEGARG="-fedpegscript=5221$(echo $PUBKEY1)21$(echo $PUBKEY2)52ae"
e1-dae $FEDPEGARG
e2-dae $FEDPEGARG
~~~~

##### NOTE: The characters outside the public keys are delimiters that indicate public key and 'n of m' requirements. For example, the template for a 1-of-1 fedpegscript would be ``5121<pubkey>51ae``. When testing, you can also use the OP_TRUE script ``-fedpegscript=51`` so that you do not have to provide any pubkey values as we have above.

Create some generate receiving addresses (as we deleted the wallets associated with them above) and mature some outputs on each chain:

~~~~
ADDRGEN1=$(e1-cli getnewaddress)
ADDRGEN2=$(e2-cli getnewaddress)
e1-cli generatetoaddress 101 $ADDRGEN1
b-cli generatetoaddress 101 $ADDRGENB
~~~~

If we generate a peg-in address from Alice's daemon you'll notice the returned data contains two properties:

~~~~
e1-cli getpeginaddress
~~~~

This returns something like:

<div class="console-output">"mainchain_address": "2N5T7PThhbY3umZEjyaERrwct8rdevFCK1n",
"claim_script": "0014ba382e7ebe74d2ed16201aeb21ee83bf4448906b"
</div>

If we execute the command again you'll notice that the returned data changes:

~~~~
e1-cli getpeginaddress
~~~~

The reason is that  each time we generate a new peg-in address we are asking the daemon to create a new mainchain address, as well as a new script that will need satisfying in order to claim the peg-in.

A user would send coins from their Bitcoin wallet to the "mainchain_address" value returned from the command as shown below. Like getnewaddress, getpeginaddress adds new secrets to wallet.dat, necessitating backup on a regular basis.

With that established, let's store the data returned in some variables for use later:

~~~~
ADDRS=$(e1-cli getpeginaddress)
MAINCHAIN=$(echo $ADDRS | jq -r '.mainchain_address')
CLAIMSCRIPT=$(echo $ADDRS | jq -r '.claim_script')
~~~~

We'll be moving bitcoin to this address so we'll check existing balances in our Bitcoin wallet first:

~~~~
b-cli getwalletinfo
~~~~

That shows a current balance of 50 bitcoin:

<div class="console-output">"balance": 50.00000000,
</div>

Now we'll send funds to our unique [Watchmen]({{ site.url }}/how-it-works) P2SH address:

~~~~
TXID=$(b-cli sendtoaddress $MAINCHAIN 1)
~~~~

And check that the bitcoin have indeed left our Bitcoin wallet:

~~~~
b-cli getwalletinfo
~~~~

Which now shows a balance of:

<div class="console-output">"balance": 48.99996240,
</div>

Note that the reason the balance is not now 49 is because fees were also deducted from the wallet in order to send the transaction.

In order to claim the peg-in amount in our sidechain we need to first mature the funding transaction. This rule ensures that the funds being created in our sidechain are not prone to a reorganisation on the main chain.

##### NOTE: The 'peginconfirmationdepth' parameter can be used to override the default confirmation depth, which is 10 blocks (8 plus 2 for the wallet to avoid race conditions between nodes). This forms part of the network's consensus rules and so it must be set on chain initialization. As a guide, Liquid is a production implementation of Elements that is pegged to Bitcoin and uses 102 (100 plus 2 for the wallet).

~~~~
b-cli generatetoaddress 101 $ADDRGENB
b-cli getwalletinfo
~~~~

Get the merkel path as proof that the transaction was included in a block:

~~~~
PROOF=$(b-cli gettxoutproof '''["'''$TXID'''"]''')
RAW=$(b-cli getrawtransaction $TXID)
~~~~

We will now attempt to claim the funds within our sidechain. The claim takes the form of a transaction and contains the proof returned above:

##### NOTE: There is an optional third argument of "sidechainaddress" which can be provided to "claimpegin". This is not needed if you are calling the command from the same wallet that "owns" the address.

~~~~
CLAIMTXID=$(e1-cli claimpegin $RAW $PROOF $CLAIMSCRIPT)
~~~~

Bob's node (as well as Alice's of course) should accept the claim transaction as valid and add it to its mempool. Create a block containing the transaction:

~~~~
e2-cli generatetoaddress 1 $ADDRGEN2
~~~~

We should be able to see the confirmation:

~~~~
e1-cli getrawtransaction $CLAIMTXID 1
~~~~

The output of which can be seen below. Note that the asset hex may differ from that below:

<div class="console-output">"value": 0.99992800,
"asset": "b7c9431837115ba3b8a1753dc227311ab4480c14d97484234f984d361a00c966",
</div>

Remember that fees will also have been deducted on the sidechain from the amount received above.

As the wallet started with 21 million bitcoin it should have nearly 1 more now (1 minus sidechain fees). To check:

~~~~
e1-cli getwalletinfo
~~~~

Which returns the claimed amount in the "unconfirmed_balance" value (minus fee)::

<div class="console-output">"unconfirmed_balance": {
    "bitcoin": 0.99992800
  },
</div>

Now that we have sent assets into our sidechain (peg-in) we will now peg-out and send assets representing bitcoin back from our sidechain to the main chain:

~~~~
e1-cli sendtomainchain $(b-cli getnewaddress) 1
e1-cli generatetoaddress 1 $ADDRGEN1
e1-cli getwalletinfo
~~~~

The results of which show:

<div class="console-output">“balance”: {
    "bitcoin": 20999999.99987460
},
</div>

Which is the amount before peg-out plus the (now confirmed) "unconfirmed_balance" balance we saw above, minus the 1 we pegged out. Remember that sidechain fees have also been deducted so we return to just under the original 21 million. These fees actually show in the "immature_balance" for Alice's wallet, as she mined the block and collected her own fees:

<div class="console-output">”immature_balance”: {
    “bitcoin”: 0.00012540
},
</div>

Checking this manually:

<div class="console-output">20999999.99987460 + 0.00012540 = 21000000
</div>


We'll now shut down our daemons. 

~~~~
e1-cli stop
e2-cli stop
b-cli stop
~~~~

Now that we have run through the code step-by-step you should have a good understanding of how Elements works. 

If you would like to run through the code again there is a much easier way to do this rather than typing out or copying/pasting the code from the tutorial pages each time. How to do this is detailed in the [An easy way to run the main tutorial code]({{ site.url }}/elements-code-tutorial/easy-run-code) section. This lets you run the code line by line or automatically execute it up until a point you are interested in looking at in more detail.


[Next: Elements as a standalone Blockchain]({{ site.url }}/elements-code-tutorial/blockchain)

