---
layout: page
title: Installing Bitcoin
permalink: /elements-code-tutorial/installing-bitcoin
---

# Elements code tutorial

## Installing Bitcoin

First we will install Bitcoin on the machine. This will allow us to demonstrate how the Federated 2-Way Peg works in Elements later on in the tutorial. It is not required if you intend to use Elements as a standalone blockchain, but to fully understand the features available in Elements it is a good idea to follow along anyway. It doesn’t take long to install Bitcoin using the commands below and we will be running in "regtest" mode, so there is no blockchain to sync.

We will be using the [Bitcoin PPA for ubuntu](https://launchpad.net/~bitcoin/+archive/ubuntu/bitcoin) but you can also compile it from source by following instructions on the [Bitcoin Core repository](https://github.com/bitcoin/bitcoin). Open a new terminal window and run the following `terminal commands` one after the other:

~~~~
sudo apt-add-repository ppa:bitcoin/bitcoin
sudo apt-get update
sudo apt-get install bitcoind
~~~~

Check that the install worked:

~~~~
which bitcoind
~~~~

Which should return:

<div class="console-output">/usr/bin/bitcoind</div>

That means that the Bitcoin software was installed. We will run it later so let’s move on with getting Elements set up.


[Next: Installing Elements]({{ site.url }}/elements-code-tutorial/installing-elements)

