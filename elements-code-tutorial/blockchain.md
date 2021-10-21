---
layout: page
title: Elements Blockchain
permalink: /elements-code-tutorial/blockchain
---

# Elements code tutorial

## Elements as a standalone Blockchain

So far, we have worked with Elements by running it as a sidechain to Bitcoin's blockchain. We'll now take a look at how to run a standalone blockchain, with no links or references to the default "bitcoin" asset at all. 

In this section we will:

**1.**&nbsp;&nbsp;&nbsp;&nbsp;Initialize a new Elements blockchain with a default asset named "newasset".

**2.**&nbsp;&nbsp;&nbsp;&nbsp;Specify 1,000,000 "newasset" to be created on initialization. 

**3.**&nbsp;&nbsp;&nbsp;&nbsp;Specify 2 reissuance tokens for "newasset" on initialization.

**4.**&nbsp;&nbsp;&nbsp;&nbsp;Claim all the anyone-can-spend "newasset" coins. 

**5.**&nbsp;&nbsp;&nbsp;&nbsp;Claim all the anyone-can-spend reissuance tokens for "newasset".

**6.**&nbsp;&nbsp;&nbsp;&nbsp;Send the asset and its reissuance token to another node's wallet.

**7.**&nbsp;&nbsp;&nbsp;&nbsp;Reissue more "newasset" from both nodes.

We'll assume that you have already run the [Installing Elements]({{ site.url }}/elements-code-tutorial/installing-elements) and [Setting up your working environment]({{ site.url }}/elements-code-tutorial/working-environment) tutorial sections to install and prepare Elements for use.
   
In order to run Elements as a standalone blockchain, we will make use of a few parameters that can either be added to the elements.conf file or passed in on node startup. They are:

<div class="console-output">defaultpeggedassetname

feeasset

initialreissuancetokens

Initialfreecoins
</div>

First we'll make sure our daemons are not running: 

~~~~
e1-cli stop
e2-cli stop
~~~~

We will have no use for the Bitcoin daemon so if it is still running from a previous tutorial session you can use the command "b-cli stop" to shut it down.

Next we'll need to clear out any existing blockchain and wallet data as we will be creating a new blockchain from scratch. If these error that is fine and likely just because the data does not exist anyway:

~~~~
cd 
rm -rf ~/elementsdir1/elementsregtest
rm -rf ~/elementsdir2/elementsregtest
~~~~

Now we will perform steps 1 to 3 above. Start our two nodes with a few parameters which are used to configure the initialization of our blockchain:

~~~~
STANDALONEARGS=("-validatepegin=0" "-defaultpeggedassetname=newasset" "-initialfreecoins=100000000000000" "-initialreissuancetokens=200000000")
e1-dae ${STANDALONEARGS[@]}
e2-dae ${STANDALONEARGS[@]}
~~~~

Now we need to create default wallets and rescan the blockchain:

~~~
e1-cli createwallet ""
e2-cli createwallet ""
e1-cli rescanblockchain
e2-cli rescanblockchain
~~~

We'll look at what these parameters do in more detail next.

* * * 

### validatepegin
As we will not be running our blockchain as a sidechain, we need to disable the validation of the peg-in, as there will be no peg to validate.

### defaultpeggedassetname
This allows you to specify the name of the default asset created upon blockchain initialization. If you do not provide this, the default asset created by Elements will be labelled as "bitcoin".

### initialfreecoins
The number (in the equivalent of Bitcoin's Satoshi unit) of the default asset to create. 

### initialreissuancetokens
The number (in the equivalent of Bitcoin's Satoshi unit) of the reissuance tokens for the default asset to create. 

* * * 

Checking that the blockchain initialization worked as expected:

~~~~
e1-cli getwalletinfo
e2-cli getwalletinfo
~~~~

The results of which will look something like the following. Note that we can see the asset and its reissuance token as already being present in both wallet balances because they are currently in 'anyone-can-spend' addresses. 

<div class="console-output">"balance": {
    "newasset": 1000000.00000000,
    "a6be6b365498cd451be75ba0f68c258ee01e08f3cb30d5f8469f6628db58dc61": 2.00000000
</div>

In order to claim the reissuance token, we need to take the hex it was assigned upon creation and store it in a variable for later use. The default asset has been labelled as 'newasset', as our chain initialization parameters specified, and so can easily be referred to later.

##### NOTE: The exact hex of your reissuance token may differ from that above, so you may need to change the following line of code to represent the hex you have.

~~~~
DEFAULTRIT=a6be6b365498cd451be75ba0f68c258ee01e08f3cb30d5f8469f6628db58dc61
~~~~

When implementing code in an application, you'll need to execute something similar to the code below in order to extract the hex id from the wallet balance list at runtime. It should be noted that the example code is limited in scope to our ‘two assets, one of which has a known label’ setup. You can ignore the following example code block for now and just use the line above and replace the actual hex id with the one your reissuance token was assigned.

~~~~
DEFAULTRIT=$(e1-cli getwalletinfo | jq '[.balance] | .[0]' | jq -r 'keys | .[0]')

if [ $DEFAULTRIT = "newasset" ]; then
  DEFAULTRIT=$(e1-cli getwalletinfo | jq '[.balance] | .[0]' | jq -r 'keys | .[1]')
fi
~~~~

Having initialized our new blockchain with 1,000,000 assets named "newasset" and 2 reissueance tokens for "newasset", we can now progress through steps 4 to 7. Start by having the e1 node claim the anyone-can-spend balances:

~~~~
e1-cli sendtoaddress $(e1-cli getnewaddress) 1000000 "" "" true
~~~~

Note that we did not need to specify the asset being sent, as "newasset" will be used by default.

Now claim the anyone-can-spend reissuance token and generate some blocks to confirm the transactions. We also need to recreate the generate receiving addresses, because we deleted the corresponding wallets above.

~~~~
e1-cli sendtoaddress $(e1-cli getnewaddress) 2 "" "" false false 1 UNSET false $DEFAULTRIT
ADDRGEN1=$(e1-cli getnewaddress)
ADDRGEN2=$(e2-cli getnewaddress)
e1-cli generatetoaddress 101 $ADDRGEN1
~~~~

It is worth noting that addresses in Elements can receive different types of asset, so we could have sent both 'newasset' and its reissuance token to the same address.

Send some of the "newasset" to e2, who currently holds none of our default asset or its reissuance tokens. Note that we do not have to specify the type of asset to be sent as the default asset is "newasset":

~~~~
e1-cli sendtoaddress $(e2-cli getnewaddress) 500 "" "" false 
~~~~

Send some of the reissuance tokens to e2 and confirm the two transactions:

~~~~
e1-cli sendtoaddress $(e2-cli getnewaddress) 1 "" "" false false 1 UNSET false $DEFAULTRIT
e1-cli generatetoaddress 101 $ADDRGEN1
~~~~

Check that the wallet has updated accordingly:

~~~~
e1-cli getwalletinfo
e2-cli getwalletinfo
~~~~

Reissue some of the default asset from e1:

~~~~
e1-cli reissueasset newasset 100
e1-cli generatetoaddress 101 $ADDRGEN1
~~~~

Check that worked:

~~~~
e1-cli getwalletinfo
~~~~

Reissue some of the default asset from e2:

~~~~
e2-cli reissueasset newasset 100
e2-cli generatetoaddress 101 $ADDRGEN2
~~~~

Check that worked:

~~~~
e1-cli getwalletinfo
e2-cli getwalletinfo
~~~~

That's it! We have set up our own standalone blockchain and checked that basic functionality works as we would expect.

All other operations are the same as in the main sections of the tutorial but will use "newasset" instead of "bitcoin" as the default asset.


[Next: Developing applications on Elements]({{ site.url }}/elements-code-tutorial/application-development)

