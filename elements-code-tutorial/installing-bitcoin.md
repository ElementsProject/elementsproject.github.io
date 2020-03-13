---
layout: page
title: Installing Bitcoin
permalink: /elements-code-tutorial/installing-bitcoin
---

# Elements code tutorial

## Installing Bitcoin

First we will install Bitcoin on the machine. This will allow us to demonstrate how the Federated 2-Way Peg works in Elements later on in the tutorial. It is not required if you intend to use Elements as a standalone blockchain, but to fully understand the features available in Elements it is a good idea to follow along anyway. It doesn’t take long to install Bitcoin using the commands below and we will be running in "regtest" mode, so there is no blockchain to sync.

For ease of use, we will be using the [Bitcoin PPA for ubuntu](https://launchpad.net/~bitcoin/+archive/ubuntu/bitcoin). It should be noted that the PPA is now marked as no longer being supported, so you may prefer to follow the instructions on the [Bitcoin Core repository](https://github.com/bitcoin/bitcoin) and compile Bitcoin Core from source. If you experience issues relating to the Berkeley database during Bitcoin build configuration, follow the Berkeley install steps from the [Installing Elements]({{ site.url }}/elements-code-tutorial/installing-elements) section.

Assuming you choose to use the PPA, open a new terminal window and run the following `terminal commands` one after the other:

~~~~
sudo apt-add-repository ppa:bitcoin/bitcoin
sudo add-apt-repository universe
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

