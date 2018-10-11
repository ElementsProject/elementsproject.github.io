---
layout: page
title: Installing Bitcoin
permalink: /elements-code-tutorial/blockchain
---

# Elements code tutorial

## Elements as a standalone Blockchain

So far, we have worked with Elements by running it as a sidechain to Bitcoin's blockchain. We'll now take a look at how to run a standalone blockchain, with no links or references to the default "bitcoin" asset at all. 

In this section we will:

**1.**&nbsp;&nbsp;&nbsp;&nbsp;Initialise a new elements blockchain with a default asset named "newasset".

**2.**&nbsp;&nbsp;&nbsp;&nbsp;Specify 1,000,000 new asset to be created on initialisation. 

**3.**&nbsp;&nbsp;&nbsp;&nbsp;Specify 2 reissuance tokens for our default asset on initialisation.

**4.**&nbsp;&nbsp;&nbsp;&nbsp;Claim all the anyone-can-spend "newasset" coins. 

**5.**&nbsp;&nbsp;&nbsp;&nbsp;Claim all the anyone-can-spend reissuance tokens for "newasset".

**6.**&nbsp;&nbsp;&nbsp;&nbsp;Send the asset and its reissuance token to another node's wallet.

**7.**&nbsp;&nbsp;&nbsp;&nbsp;Reissue more "newasset" from both nodes.

We'll assume that you have already run the [Installing Elements]({{ site.url }}/elements-code-tutorial/installing-elements) and [Setting up your working environment]({{ site.url }}/elements-code-tutorial/working-environment) tutorial sections to install and prepare elements.
   
In order to run Elements as a stand-alone blockchain we will make use of a few parameters that can be added to the elements.conf file or passed in on node start up. They are:

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
rm -r ~/elementsdir1/elementsregtest/blocks
rm -r ~/elementsdir1/elementsregtest/chainstate
rm ~/elementsdir1/elementsregtest/wallet.dat
rm -r ~/elementsdir2/elementsregtest/blocks
rm -r ~/elementsdir2/elementsregtest/chainstate
rm ~/elementsdir2/elementsregtest/wallet.dat
~~~~

Now we will perform steps 1 - 3 above. This is done by starting our two nodes with a few parameters used to configure the initialisation of our blockchain.

~~~~
STANDALONEARGS="-validatepegin=0 -defaultpeggedassetname=newasset -initialfreecoins=100000000000000 -initialreissuancetokens=200000000"
e1-dae $STANDALONEARGS
e2-dae $STANDALONEARGS
~~~~

Let's look at what these parameters do in more detail:

* * * 

### validatepegin
As we will not be running this blockchain as a sidechain we need to disable the validation of the peg in, as there will be no peg to validate.

### defaultpeggedassetname
Allows you to specify the name of the default asset created upon blockchain initialisation. If you do not provide this the default asset created by Elements will be labelled as "bitcoin".

### initialfreecoins
The number (in the equivalent of Bitcoin's satoshi unit) of the default asset to create. 

### initialreissuancetokens
The number (in the equivalent of Bitcoin's satoshi unit) of the reissuance token for the default asset to create. 

* * * 

Checking that the blockchain initialisation worked as expected:

~~~~
e1-cli getwalletinfo
e2-cli getwalletinfo
~~~~

Having initialised our new blockchain with 1,000,000 assets named "newasset" and 2 reissueance tokens for "newasset", we can now progress to steps 4 - 7. Start by having the e1 node claim the anyone-can-spend balances and generate some blocks so they become spendable:

~~~~
e1-cli sendtoaddress $(e1-cli getnewaddress) 1000000 "" "" true
e1-cli generate 101
~~~~

Note that we did not need to specify the asset being sent as "newasset" will be used by default. In order to claim the reissuance token, we first need to find what hex it has been assigned upon creation:

~~~~
e1-cli getwalletinfo
~~~~

The results of which will look something like this:

<div class="console-output">"balance": {
    "mynewasset": 1000000.00000000,
    "a6be6b365498cd451be75ba0f68c258ee01e08f3cb30d5f8469f6628db58dc61": 2.00000000
</div>

We'll store the reissuance hex in a variable for later use:

##### NOTE: The exact hex of your reissuance token may differ from that above if you have amended anything used to initialise the chain, and so you may need to change the following line to represent the hex you have.

<div class="console-output">DEFAULTRIT=a6be6b365498cd451be75ba0f68c258ee01e08f3cb30d5f8469f6628db58dc61
</div>

Now claim the anyone-can-spend reissuance token:

~~~~
e1-cli sendtoaddress $(e1-cli getnewaddress) 2 "" "" false $DEFAULTRIT
e1-cli generate 101
~~~~

Send some of the "newasset" to e2, who currently holds no amount of our default asset or its reissuance token. Note that we do not have to specify the type of asset to be sent as the default asset is "newasset":

~~~~
e1-cli sendtoaddress $(e2-cli getnewaddress) 500 "" "" false 
e1-cli generate 101
~~~~

Send some of the reissuance tokens to e2:

~~~~
e1-cli sendtoaddress $(e2-cli getnewaddress) 1 "" "" false $DEFAULTRIT
e1-cli generate 101
~~~~

Check that the wallet has updated accordingly:

~~~~
e1-cli getwalletinfo
e2-cli getwalletinfo
~~~~

Reissue some of the default asset from e1:

~~~~
e1-cli reissueasset newasset 100
e1-cli generate 101
~~~~

Check that worked:

~~~~
e1-cli getwalletinfo
~~~~

Reissue some of the default asset from e2:

~~~~
e2-cli reissueasset newasset 100
e2-cli generate 101
~~~~

Check that worked:

~~~~
e1-cli getwalletinfo
e2-cli getwalletinfo
~~~~

That's it! We have set up our own standalone blockchain and checked that basic functionality works as we would expect.

All other operations are the same as in the main sections of the tutorial but will use "newasset" instead of "bitcoin" as the default asset.


[Next: Developing applications on Elements]({{ site.url }}/elements-code-tutorial/application-development)

