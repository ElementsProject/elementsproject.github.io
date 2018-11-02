---
layout: page
title: Working environment
permalink: /elements-code-tutorial/working-environment
---

# Elements code tutorial

## Setting up your working environment

First we need to set up our working directories. Start by moving to the home directory:

~~~~
cd
~~~~

We'll create some test directories to hold the data used by our Bitcoin node and also two instances of an Elements node. First we'll try and remove any files and directories that might be left over from running through this tutorial previously. Don't worry if the first 3 lines throw errors, it just means the directories are not there and can't be removed.

~~~~
rm -r ~/bitcoindir
rm -r ~/elementsdir1
rm -r ~/elementsdir2
mkdir ~/bitcoindir
mkdir ~/elementsdir1
mkdir ~/elementsdir2
~~~~

We need to set up our config files now. We'll do that by copying the example configuration files contained within the elements source code into the working directories we just created. 

~~~~
cp ~/elements/contrib/assets_tutorial/bitcoin.conf ~/bitcoindir/bitcoin.conf
cp ~/elements/contrib/assets_tutorial/elements1.conf ~/elementsdir1/elements.conf
cp ~/elements/contrib/assets_tutorial/elements2.conf ~/elementsdir2/elements.conf
~~~~

If you take a quick look in each of the 3 config files you will see that they contain a flag to tell the nodes to operate in "regtest" mode. They also contain RPC information such as port, username and password. The Elements config files also contain details of the Bitcoin node's RPC authentication data. They need access to this information in order to authenticate and make calls to the Bitcoin node later.

By using a different port for each Element daemon we can run more than one node at a time on a single machine. This is useful for the purposes of our tutorial.

Before we start running our nodes we'll take the time to configure some terminal aliases. These will let us pass commands to each node's client in a much simpler way. 

Before we do that it is worth noting how the set up of a daemon and client works for both Bitcoin and Elements.

**The Bitcoin and Elements node software (bitcoind, elementsd) both run as daemons, executing code as background services that can be remotely accessed using the Bitcoin and Elements client software (bitcoin-cli, elements-cli). The clients simplify the process of sending commands to their associated daemon and receiving the returned output data.**

#### Note: Once our Bitcoin and Elements daemons have been started we will only interact with them through the use of the clients, which will use the same config file as their associated daemon, enabling them to pass authentication checks when making RPC calls.

Now let's create the terminal aliases for our Bitcoin and Elements daemons and clients: 

~~~~
cd elements
cd src
shopt -s expand_aliases
alias b-dae="bitcoind -datadir=$HOME/bitcoindir"
alias b-cli="bitcoin-cli -datadir=$HOME/bitcoindir"
alias e1-dae="$HOME/elements/src/elementsd -datadir=$HOME/elementsdir1"
alias e1-cli="$HOME/elements/src/elements-cli -datadir=$HOME/elementsdir1"
alias e2-dae="$HOME/elements/src/elementsd -datadir=$HOME/elementsdir2"
alias e2-cli="$HOME/elements/src/elements-cli -datadir=$HOME/elementsdir2"
~~~~

We now have an easy way to call each daemon and client. Instead of having to type "$HOME/elements/src/elements-cli -datadir=$HOME/elementsdir1" every time we can just use the "e1-cli" alias for example.

#### Note: If you want to skip ahead to running a non-sidechain Elements based blockchain you can move to the [Elements as a standalone Blockchain]({{ site.url }}/elements-code-tutorial/blockchain) section now. You can return to the [Using Elements to perform basic operations]({{ site.url }}/elements-code-tutorial/basic-operations) section after that, but you may have to amend the commands used occasionally. 

If you want to run a few examples of your own after following this tutorial you can add the aliases above to your ‘bashrc' file which means that they will be available to you in every new terminal window you open. To do this run "nano ~/.bashrc" and paste the 6 alias lines above into the file, save and exit. The next time you open a new terminal window the aliases will have already been set. 

**We don't need to do this now** but if at any point you need to exit the tutorial and start again you should run the following commands to shut down the Bitcoin and Elements daemons first before restarting:

~~~~
b-cli stop
e1-cli stop
e2-cli stop
~~~~

If we now try and execute a simple command using the Bitcoin RPC client it will error as we haven't started the Bitcoin daemon yet:

~~~~
b-cli -getinfo
~~~~

That gives us an error, but we were expecting that. Let's start the Bitcoin daemon:

~~~~
b-dae
~~~~

Give it a few seconds to start up, and then try again:

~~~~
b-cli -getinfo
~~~~

That should work! You should get a response telling you general information about the node.

So now we have the Bitcoin daemon running we need to start our Elements daemons. Remember that we are able to run two instances on the same machine because we are using different configuration files with different ports and RPC permissions set. Start the daemons:

~~~~
e1-dae
e2-dae
~~~~

Give them a few seconds to start up and then check they are running:

~~~~
e1-cli getwalletinfo
e2-cli getwalletinfo
~~~~

A quick query of the running processes should also list them as active:

~~~~
ps -A | grep elementsd
~~~~

The two instances should show in the list.

[Next: Using Elements to perform basic operations]({{ site.url }}/elements-code-tutorial/basic-operations)

