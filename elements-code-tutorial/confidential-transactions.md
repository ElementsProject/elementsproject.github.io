---
layout: page
title: Confidential Transactions
permalink: /elements-code-tutorial/confidential-transactions
---

# Elements code tutorial

## Using Confidential Transactions

Let's say that our old cryptographic friends Alice and Bob are the owners of these wallets. Alice owns the e1 wallet and Bob owns e2.

We'll have Bob send some assets to himself using a blinded Elements address as the destination. All addresses in Elements are blinding using Confidential Transactions by default. Blinding is the process by which amount and the type of asset being transferred are cryptographically hidden from everyone except the sending and receiving parties. This is done using a blinding key, which we will look at later. First generate a new address and store it in a variable named "ADDR" for future use:

~~~~
ADDR=$(e2-cli getnewaddress)
~~~~

To see the new address we generated for Bob we can print out the value that was stored in the "ADDR" variable. 

~~~~
echo $ADDR
~~~~

##### NOTE: We will use the 'store the value returned from an RPC command in a variable for future use" trick throughout the rest of the tutorial. We'll print out the contents of a variable every now and again when highlighting something that has been returned but you can always just echo out the contents of any others as we go along should you want to.

After running the echo command above you should see something similar to this:

<div class="console-output">CTEmrcu15cikk88srA4Mh4REo3BeHePKNen9nFZJJQbK4f82NjRSGSwiZYkSAS69UwWanRBgQGWwvSer</div>

In the Elements protocol blinded addresses start with "CTE" and unblinded addresses start with a "2", so we can see that the address we have just generated for Bob is indeed a blinded, confidential address.

Rather than just trust the "CTE" address prefix let's look at the address in more detail to check that it is a confidential address. To do this we can use the "validateaddress" command, passing in the address that we stored in the ADDR variable as a parameter:

~~~~
e2-cli validateaddress $ADDR
~~~~

You should see a long value for the "confidential_key" property. It will look something like this:

<div class="console-output">"confidential_key": "025030f91c82297493e5dbe64aa63eef4a087a79f563d50c25c8ad0122f7547212"</div>

The confidential_key is the public blinding key, which is added to the confidential transaction address and is the reason why the address itself is so long. This is what allows the sender to send and hide the amount being sent from other people.

We'll now send an amount of 1 "bitcoin" from Bob's wallet to the new address we generated for him:

~~~~
TXID=$(e2-cli sendtoaddress $ADDR 1)
~~~~

In order to have the transaction confirm we need to generate a block that will include it. As an aside we can now query the mempool of each of our Elements nodes to see the transaction waiting to be added to a block and the current block count of each node's blockchain:

~~~~
e1-cli getrawmempool
e2-cli getrawmempool
e1-cli getinfo
e2-cli getinfo
~~~~

Both should display just one transaction with the same ID as that stored in TXID and a block value of 202. If it does not, wait a few seconds and try the calls again as it may take a moment for the nodes to synchronise. Now let's generate a block and get the transaction confirmed for Bob:

~~~~
e2-cli generate 1
~~~~

Checking the mempool again for each client will show that it is now empty:

~~~~
e1-cli getrawmempool
e2-cli getrawmempool
~~~~

Checking the "blocks" property of getinfo for each client should show 203, the 202 blocks that we generate by twice running "generate 101" further up and the block that we just mined that contains Bob's last transaction where he sent to the confidential address: 

~~~~
e1-cli getinfo
e2-cli getinfo
~~~~

Note that although Bob sent an amount of 1 to himself the net effect is that he now has slightly less than he did before - the is because some of the transaction amount was spent on fees. Although Bob mined the block and will collect the fees he will need to wait until the block has matured before he sees it in his wallet.

##### NOTE: sidechain fees are defaulted to "bitcoin". This can be changed using the -feeasset parameter.

The above shows that the client's blockchains and mempools are in sync. If they are not, wait a few seconds and try the calls above again as it may take a moment for the nodes to synchronise. They display the same results not because they share a common data file but because they are connected nodes on the same Elements network and broadcast transactions and blocks between each other, in the same way a Bitcoin node would between its peers. You can check this for yourself by looking at the separate data files in the following locations and noting that they are separate stores of the same blockchain data:

<div class="console-output">/$HOME/elementsdir1/elementsregtest
/$HOME/elementsdir2/elementsregtest
</div>

Now let's examine the transaction as it is seen by Bob's wallet and also how it is seen from the point of view of Alice's wallet. First the view from Bob's wallet:

~~~~
e2-cli gettransaction $TXID
~~~~

The output from that initially looks like just a huge random assortment of letters and numbers… but if you scroll up you will see some more readable content above that, which turns out to be the hex value of the transaction.

Looking in the "details" section near the top, you will see that there are two amount values each under a category:

<div class="console-output">"details": [
  {
    ...
    "category": "send",
    "amount": -1.00000000,
    ...
  },
  {
    ....
    "category": "receive",
    "amount": 1.00000000,
    ...
  }
]
</div>

And so we can confirm that Bob's wallet can view the actual amounts being sent and received in this transaction. This is because the blinded transaction was sent from Bob's own wallet and so it has access to the required data to unblind the amount values. You will also see two other properties and their values within the two details sections: "amountblinder" and "assetblinder". These indicate that both the asset amount and the type of asset were blinded. This ensures that wallets without knowledge of the blinding key are prevented from viewing them.

Looking at the transaction from Alice's wallet we would expect both amount and type to be unknown. Checking this using Alice's wallet we may initially get an error:

~~~~
e1-cli gettransaction $TXID
~~~~

The reason that we might get an error using this command is that Alice's wallet may not contain details of the transaction yet. We can get the raw transaction data from Alice's node's copy of the blockchain using the getrawtransaction command so it definitely does like this:

~~~~
e1-cli getrawtransaction $TXID 1
~~~~

That returns raw transaction details. If you look within the "vout" section you can see that there are three instances. The first two instances are the receiving and change amounts and the third is the transaction fee. Of these three the fee is the only one in which you can see a value amount, as the fee is unblinded. For the first two instances you will see (amongst others) properties with values similar to this:

<div class="console-output">"value-minimum": 0.00000001,
"value-maximum": 11258999.06842624,
"amountcommitment": "0881c61d8a15ad26e6ef621ca99a188ccebbdb348d5285012393459b7e5b1e6113",
"assetcommitment": "0b1b7a1a4a604f4a68b3277e3a8926d74e86adce7b92e8e6ba67f9c5a8ad2cbcf4",
</div>

What this shows are the "blinded ranges" of the value amounts and the commitment data that acts as proof of the actual amount and type of asset transacted.

Even if we were to import Bob's private key into Alice's wallet it would still not be able to see the amounts and type of asset because it still has no knowledge of the blinding key used. Let's show this by first importing the private key used by Bob's wallet into Alice's:

~~~~
e1-cli importprivkey $(e2-cli dumpprivkey $ADDR)
~~~~

Now that we have imported the private key the call to "gettransaction" will not error:

~~~~
e1-cli gettransaction $TXID
~~~~

But because Alice still does not know the blinding key, the amount (scroll up to see it) will show as:

<div class="console-output">"amount": {
    "bitcoin": 0.00000000
</div>

It will also not show up in the wallet's balance for owned outputs. Check this by running:

~~~~
e1-cli getwalletinfo
~~~~

Which still shows the same balance as before we imported the private key:

<div class="console-output">"balance": {
    "bitcoin": 10500000.00000000
</div>

As the amount for the transaction is unknown it will also not show up in her Wallet's list of unspent outputs. Note: the two parameters are "minimum confirmations" and "maximum confirmations", passing 1 in for each means we only search the space the utxo above is in (although we expect not to see it):

~~~~
e1-cli listunspent 1 1
~~~~

Which simply returns nothing.

All of that shows that without knowledge of the blinding key the amount, and also the type, of asset being transacted is indeed hidden.

In order for anyone to view the amount and type of assets being transacted they need to know the blinding key that was used to generate the blinded address. To show this we can pull the blinding key Bob's wallet used, import it into Alice's wallet and try again to view the transaction. Let's export the key for that particular address from Bob's wallet and import it into Alice's in one step:

~~~~
e1-cli importblindingkey $ADDR $(e2-cli dumpblindingkey $ADDR)
~~~~

Now that Alice's wallet has knowledge of the blinding key used on that address, let's run the checks we did above from Alice's wallet:

~~~~
e1-cli getwalletinfo
~~~~

Magic! Alice's wallet now shows:

<div class="console-output">"balance": {
    "bitcoin": 10500001.00000000
</div>

Checking the unspent outputs Alice's wallet is aware of should now also show the output in its list. Remember that we previously imported the private key from Bob's wallet so Alice's now treats it as her own:

~~~~
e1-cli listunspent 1 1
~~~~

This shows that both the amount and the type within the unspent output are now visible to Alice's wallet:

<div class="console-output">"amount": 1.00000000,
"asset": "b2e15d0d7a0c94e4e2ce0fe6e8691b9e451377f6e46e8045a86f7c4b5d4f0f23",
"assetlabel": "bitcoin",
</div>

The "asset" value and "assetlabel" properties are something we will look at in more detail in the next section covering asset issuance. Before that we can show that Alice's wallet's view of the transaction is now identical to Bob's:

~~~~
e1-cli gettransaction $TXID
~~~~

We"ve seen that the use of a blinding key hides the amount and type of assets in an address and that be importing a blinding key we can reveal those values. In practical use a blinding key may be given to an auditor should there be a need to verify the amounts and types of assets held by a party within the system you are designing. The Confidential Transactions feature of Elements also allows for "range proofs" to be performed without the need to expose actual amounts. This allows for statements like "address abc holds at least an amount of x of asset y" to be cryptographically proven.

In the next section we will look at how to issue, label and re-issue your own assets within an Elements blockchain.


[Next: Issuing your own assets]({{ site.url }}/elements-code-tutorial/issuing-assets)

