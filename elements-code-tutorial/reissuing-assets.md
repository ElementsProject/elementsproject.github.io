---
layout: page
title: Installing Bitcoin
permalink: /elements-code-tutorial/reissuing-assets
---

# Elements code tutorial

## Reissuing assets

Now we'll look at reissuing an amount of the asset. Reissuing just means "creating some more of" in this context. We have already created and labelled the new asset so the command to issue some more is quite simple. We'll store the result of the "reissueasset" command in a variable ("RTRANS") and strip out the transaction ID from that, storing it in another variable ("RTXID") for future use:

~~~~
RTRANS=$(e1-cli reissueasset $ASSET 99)
RTXID=$(echo $RTRANS | jq '.txid' | tr -d '"')
~~~~

We've just created 99 more units of the asset! As an aside: because we have labelled the asset we could have also passed "demoasset" in as the first parameter instead of the hex identifier and this would have worked exactly the same.

To check this issuance history of our asset (and ignore the "bitcoin" issuance) we can run the "listissuances" command and specify the asset we are interested in:

~~~~
e1-cli listissuances $ASSET
~~~~

For the above command it is worth noting that only the hex value can be passed in and not the label. Along with the original issuance you should see a new entry with the following property:

<div class="console-output">"isreissuance": true,
</div>

This property allows us to differentiate between initial issuances and reissuances. Note that the transaction ID where amounts of the asset were created is also included in the returned data.

Let's look at the details of the transaction where we reissued our asset: 

~~~~
e1-cli gettransaction $RTXID
~~~~

Scroll to the top of the returned transaction data as there are a few things worth noting here. The first is that within the "amount" section we can see that 0 "bitcoin" and 99 "demoasset" were transacted:

<div class="console-output">"amount": {
  "bitcoin": 0.00000000,
  "33244cc19dd9df0fd901e27246e3413c8f6a560451e2f3721fb6f636791087c7": 0.00000000,
  "demoasset": 99.00000000,
</div>

This information suggests that an Elements transaction can transact more than one type of asset within the same transaction, which is indeed the case. 

##### NOTE: To send different types of asset in the same transaction, the 'sendmany" command is used. The syntax is the same as in bitcoin.

The "amount" section shows the net effect of the transaction as: 0 "bitcoin", 99 "demoasset" and also another asset that is 0. That unlabelled asset is our issuance token (the hex for which will differ from that above but the results are otherwise the same). What this shows is that once the sent and received amounts are totalled we have created 99 "demoasset". We'll see how the values in the "amount" section are derived now by scrolling down the returned data and looking within the "details" section. You will see that amounts of 99 "demoasset" and 1 reissuance token sere sent:

<div class="console-output">"category": 'send",
"amount": -99.00000000,

"category": 'send",
"amount": -1.00000000,
</div>

And that further on the same amounts were received:

<div class="console-output">"category": "receive",
"amount": 99.00000000,

"category": "receive",
"amount": 1.00000000,
</div>

We can see how the "amount" section above therefore lists the net transfer of 0 (-1 +1) tokens. The reason why the net of this is the creation of 99 new "demoasset" is that the "reissueasset" command essentially spends from a zero balance address and so the received amount has the effect of creating 99 new "demoasset". The 99 new "demoasset" are basically spent into existence. It is worth highlighting again that in order to reissue an asset you must hold a related reissuance token. They must therefore be allocated wisely.

To check that the blinding works the same for a reissuance transaction as it does for a normal transaction we can check Bob's view of Alice's reissuance transaction. Wait a few seconds after running the "generate" command to let Bob's node sync, then run the second and third lines:

~~~~
e1-cli generate 1
RAWRTRANS=$(e2-cli getrawtransaction $RTXID)
e2-cli decoderawtransaction $RAWRTRANS
~~~~

We can see that the amounts and asset types are indeed blinded with results like this:

<div class="console-output">"value-minimum": 0.00000001,
"value-maximum": 42.94967296,
</div>

You could unblind these using the techniques we used for the initial issuance should you want to.

So we have seen that reissuance is just a special kind of spending transaction whereby you can create more of the original asset so long as you hold a valid reissuance token in your wallet. Next we will look at how to transfer the reissuance tokens.

Let's send the reissuance token from Alice to Bob so that he can reissue our "demoasset". Note that if there was always going to be a need for them both to reissue the asset at the same time then we could have just created two reissuance tokens and sent one to Bob, leaving Alice still holding the other. Either way we would need to send from one wallet to the other so let's begin. First we'll double check that Alice's wallet currently holds the reissuance token and Bob's does not:

~~~~
e1-cli getwalletinfo
e2-cli getwalletinfo
~~~~

Alice's wallet has "bitcoin", "demoasset" and the demo asset's reissuance token whereas Bob's only has "bitcoin".

We'll just prove that the token is needed to reissue by trying to reissue from Bob's wallet without the token:

~~~~
e2-cli reissueasset $ASSET 10
~~~~

That fails as expected and gives the following error message:

<div class="console-output">No available reissuance tokens in wallet.
</div>

So let's send the reissuance token to Bob so that he can reissue some of our "demoasset". Have Bob's wallet generate a new address and save it in a variable:

~~~~
RITRECADD=$(e2-cli getnewaddress)
~~~~

Send the token from Alice's wallet to Bob's new address as if it were any other asset: (we'll use the hex of the token to say what type of asset we are sending and generate a block so the transaction confirms)

~~~~
e1-cli sendtoaddress $RITRECADD 1 "" "" false $TOKEN
e1-cli generate 1
~~~~

Check that Bob's wallet now has the token and that Alice's no longer does:

~~~~
e2-cli getwalletinfo
e1-cli getwalletinfo
~~~~

The token and right to reissue it provides is now Bob's! 

##### NOTE: Remember from an earlier note that we can divide a reissuance token like any other asset in Elements. Our send of "1" token in this instance actually transferred 100,000,000 of the smallest possible amount of the token. You can try sending something like "0.1" of the token back to Alice and check if she is again able to reissue (she will and so will Bob - who will still hold "0.9").

Bob still doesn't have any of the "demoasset" itself of course but now that his wallet holds the reissuance token we can reissue any amount of "demoasset" and it will show in his wallet:

~~~~
RISSUE=$(e2-cli reissueasset $ASSET 10)
e2-cli getwalletinfo
~~~~

Bob's wallet now has "bitcoin", the reissuance token for our new asset and an amount of the new asset itself:

<div class="console-output">"balance": {
    "bitcoin": 10499998.99841940,
    "78ee1e3b9f2edf730e7f624e9d0f92d3e1d364c0ee91525bbccf56377dcd5033": 1.00000000,
    "600010d2a60cf0d9395dced79af3ccdb7c908e80cddf125ed1af80dc87186aae": 10.00000000
},
</div>

Remember that the new asset we issued will still only display using its hex value in Bob's wallet as we didn't assign it a label like we did in Alice's wallet. In order for Alice to see this reissuance we need to make her wallet aware of it:

~~~~
e2-cli generate 1
~~~~

Show that Alice's wallet can't see it:

~~~~
e1-cli listissuances
~~~~

Import the address so that it can:

~~~~
RITXID=$(echo $RISSUE | jq '.txid' | tr -d '"')
RIADDR=$(e2-cli gettransaction $RITXID | jq '.details[0].address' | tr -d '"')
e1-cli importaddress $RIADDR
~~~~

Now Alice's wallet can see the reissuance:

~~~~
e1-cli listissuances
~~~~

As expected though, the amounts are blinded. You can unblind by importing the blinding key as we did earlier should you want to.

In Elements you can also carry out an unblinded asset issue:

~~~~
UBRISSUE=$(e2-cli issueasset 55 1 false)
UBASSET=$(echo $UBRISSUE | jq '.asset' | tr -d '"')
~~~~

Which shows up as normal in Bob's wallet after he issues it:

~~~~
e2-cli getwalletinfo
~~~~

And this time if we import the address into Alice's wallet she should be able to see the amount issued, proving it was issued unblinded. Following the same process as before to import the address into Alice's wallet:

~~~~
e2-cli generate 1
e1-cli listissuances
UBRITXID=$(echo $UBRISSUE | jq '.txid' | tr -d '"')
UBRIADDR=$(e2-cli gettransaction $UBRITXID | jq '.details[0].address' | tr -d '"')
e1-cli importaddress $UBRIADDR
~~~~

We can now see that Alice's wallet can see both the issuance and the amount issued (55) without the need to import the blinding key:

~~~~
e1-cli listissuances
~~~~

It may also be necessary to destroy asset amounts as well as create them in an Elements based blockchain. This is easily done using the "destroyamount" command:

~~~~
e2-cli destroyamount $UBASSET 5
~~~~

Check the amount has gone from the 55 issued down to 50:

~~~~
e2-cli getwalletinfo
~~~~

It will show the amount as 50, proving that an amount of 5 were indeed destroyed:

<div class="console-output">"balance": {
    "4021bf6faac59d7ec593859a741318752f72e637e7d5ecfa54725dba1508771b": 50.00000000,
</div>

Creating, reissuing and destroying assets is a key feature of Elements that can help you reflect the real world movement of assets represented on your blockchain.


[Next: Block creation in a Strong Federation]({{ site.url }}/elements-code-tutorial/block-creation)

