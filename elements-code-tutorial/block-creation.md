---
layout: page
title: Block creation
permalink: /elements-code-tutorial/block-creation
---

# Elements code tutorial

## Block creation in a Strong Federation - Block Signing

So far, in order to create blocks we have been calling the "generate" command from either node. That's worked so far because by default block signing is OP_TRUE.

Creating a new block has been as simple as executing:

~~~~
e1-cli generatetoaddress 1 $ADDRGEN1
~~~~

Elements supports a federated signing model which allows you to specify the number of signatures required in order to produce a valid block. Let's set up something more interesting next and tell our nodes to require a valid 2-of-2 multi-signature block creation condition. This is done using the **signblockscript** parameter, which can be added to the config file or passed into the node on startup. 

First, we need to get keys from both clients so that we can then make our block sign script. That's what we'll then have to fulfill in order to produce a valid block.

Generate a new address from each of our nodes:

~~~~
ADDR1=$(e1-cli getnewaddress)
ADDR2=$(e2-cli getnewaddress)
~~~~

Validate the addresses and then extract the public key for each:

~~~~
VALID1=$(e1-cli getaddressinfo $ADDR1)
PUBKEY1=$(echo $VALID1 | jq '.pubkey' | tr -d '"')
VALID2=$(e2-cli getaddressinfo $ADDR2)
PUBKEY2=$(echo $VALID2 | jq '.pubkey' | tr -d '"')
~~~~

Now extract the private keys which we'll import later:

~~~~
KEY1=$(e1-cli dumpprivkey $ADDR1)
KEY2=$(e2-cli dumpprivkey $ADDR2)
~~~~

Now we need to generate a redeem script for a 2 of 2 multi-sig. We do this by using the "createmultisig" command and passing the first parameter as 2 and then providing two pubkeys. If we wanted to do a "2 of 3" multi-sig we'd pass 2 and then three pubkeys etc:

~~~~
MULTISIG=$(e1-cli createmultisig 2 '''["'''$PUBKEY1'''", "'''$PUBKEY2'''"]''')
REDEEMSCRIPT=$(echo $MULTISIG | jq '.redeemScript' | tr -d '"')
echo $REDEEMSCRIPT
~~~~

Stop the nodes so we can then configure them to use our new block signing method:

~~~~
e1-cli stop
e2-cli stop
~~~~

Define the requirements of block creation (must be valid against our redeemscript) and store in a variable:

~~~~
SIGNBLOCKARG="-signblockscript=$(echo $REDEEMSCRIPT) -con_max_block_sig_size=150 -con_dyna_deploy_start=0"
~~~~

We'll have to wipe out the chain we've been using so far and also the wallets and start again with a new genesis block. Note that once created you can't swap blocksigners in and out on a chain for security reasons. This may change in a later Elements release.

~~~~
rm -r ~/elementsdir1/elementsregtest/blocks
rm -r ~/elementsdir1/elementsregtest/chainstate
rm ~/elementsdir1/elementsregtest/wallets/wallet.dat
rm -r ~/elementsdir2/elementsregtest/blocks
rm -r ~/elementsdir2/elementsregtest/chainstate
rm ~/elementsdir2/elementsregtest/wallets/wallet.dat
~~~~

Start the daemons with the "signblockscript" we specified for the 2 of 2 block signing:

~~~~
e1-dae $SIGNBLOCKARG
e2-dae $SIGNBLOCKARG
~~~~

Now import the signing keys that we stored earlier before wiping the wallets: 

~~~~
e1-cli importprivkey $KEY1
e2-cli importprivkey $KEY2
~~~~

The "generate" command no longer works, even if the keys required for signing are imported into a wallet because we have started the daemons with the signblockscript argument. To make sure the error is caused by our introduction of signblockscript and not because we removed the wallets containing the generate receiving addresses, we'll create new ones. Because we have introduced the signblockscript, the generatetoaddress commands will error:

~~~~
ADDRGEN1=$(e1-cli getnewaddress)
ADDRGEN2=$(e2-cli getnewaddress)
e1-cli generatetoaddress 1 $ADDRGEN1
e2-cli generatetoaddress 1 $ADDRGEN2
~~~~

Both error with the message:

<div class="console-output">This method cannot be used with a block-signature-required chain
</div>

Because we started with the "signblockscript" argument we have to follow a new process for making, signing and released blocks.

* * *

The process will be:

* Someone (it doesn"t matter who) calls "getnewblockhex" command to propose a new block.

* The required number of block signers sign the proposed block in turn.

* The signed block is combined by someone using the "combineblocksigs" command.

* The result from "combineblocksigs" is submitted using the "submitblock" command. Again, it doesn't matter who does this as long as the block is signed and valid.

* * *

So let's do this. Start by proposing a new block:

~~~~
HEX=$(e1-cli getnewblockhex)
~~~~

So just to double check, the block count should still be zero as the proposed block has yet to be signed:

~~~~
e1-cli getblockcount
~~~~

That returns 0. And if we try and submit the block as it is:

~~~~
e1-cli submitblock $HEX
~~~~

We get an error:

<div class="console-output">block-proof-invalid
</div>

And the block count is still zero:

~~~~
e1-cli getblockcount
~~~~

So let's sort that out and sign the block using each daemon to satisfy the 2 of 2 requirement:

~~~~
SIGN1=$(e1-cli signblock $HEX)
SIGN2=$(e2-cli signblock $HEX)
~~~~

##### NOTE: Signblock tests validity before signing. This signing step can be outsourced to a Hardware Security Module (HSM) to enforce a greater level of security.

We can have each node sign the proposed block hex easily as we are running our code on a single machine with access to the output from each node's call to 'signblock'. In a production environment you will likely require remote nodes to exchange signed blocks between each other so that the requirements of the block signing script can be met. An example of how to do this can be found here: [https://github.com/nkostoulas/block-signing-federation](https://github.com/nkostoulas/block-signing-federation). The demo code uses Kafka as a communication protocol between 3 signing nodes running locally, implementing a 2-of-3 federation and can be extended for nodes running in remote locations by changing the demo's configuration.

As we are running our nodes on the same machine with easy access to the output from each node's execution of 'signblock', we can simply gather signature and pubkey data and combine them into a fully signed block like this:

~~~~
SIGN1DATA=$(echo $SIGN1 | jq '.[0]')
SIGN2DATA=$(echo $SIGN2 | jq '.[0]')
COMBINED=$(e1-cli combineblocksigs $HEX "[$SIGN1DATA,$SIGN2DATA]")
~~~~

We see a result of "True" for the "complete" property as we have signatures from enough keys to satisfy the 2 of 2 requirement. So "complete" in this context means "has enough signatures for the 'n of m' multi-sig to be valid".

Now we will extract the signed block hex from the results of the combineblocksig command and submit the block. It doesn’t matter who does this as long as they have a signed and valid block hex:

~~~~
SIGNEDBLOCK=$(echo $COMBINED | jq '.hex' | tr -d '"')
e2-cli submitblock $SIGNEDBLOCK
~~~~

Check that worked:

~~~~
e1-cli getblockcount
e2-cli getblockcount
~~~~

Yes it did, we now have moved forward one block!

We can now shut the daemons down in preparation for the next section which will explain how to "peg" your blockchain to another so it runs as a Sidechain:

~~~~
e1-cli stop
e2-cli stop
~~~~

If you do not require your blockchain to operate as a sidechain you can skip the next section and should now also execute the "b-cli stop" command to shut the bitcoin daemon down as well.


[Next: Elements as a Sidechain]({{ site.url }}/elements-code-tutorial/sidechain)

