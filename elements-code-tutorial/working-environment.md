---
layout: page
title: Working environment
permalink: /elements-code-tutorial/working-environment
---

# Elements code tutorial

## Setting up your working environment

#### Note: If you have already followed the code in this section through and understand what each step below does: you can run all the code without having to type/copy and paste it in line by line by following the instructions in the [An easy way to run the main tutorial code]({{ site.url }}/elements-code-tutorial/easy-run) section. That contains the same code as this step-by-step guide, but in a format that can be executed one line at a time just by pressing enter.

First we need to set up our working directories. Start by making sure we are running commands from the home directory:

~~~~
cd
~~~~

We'll create some test directories to hold the data used by our Bitcoin node and also two instances of an Elements node. First we'll remove any that are left from running through this guide previously. Don't worry if the first line throws any errors - it just means the directories are not there and can't be removed and is not a problem.

~~~~
rm -r ~/bitcoindir
rm -r ~/elementsdir1
rm -r ~/elementsdir2
mkdir ~/bitcoindir
mkdir ~/elementsdir1
mkdir ~/elementsdir2
~~~~

We need to set up our config files now. We'll do that by copying the configuration files from the elements source code example. 

~~~~
cp ~/elements/contrib/assets_tutorial/bitcoin.conf ~/bitcoindir/bitcoin.conf
cp ~/elements/contrib/assets_tutorial/elements1.conf ~/elementsdir1/elements.conf
cp ~/elements/contrib/assets_tutorial/elements2.conf ~/elementsdir2/elements.conf
~~~~

If you take a quick look in each of the 3 config files you will see that they contain a flag to tell the nodes to operate in "regtest" mode. They also contain RPC information such as port, username and password. The Elements config files also contain the details of the Bitcoin node's RPC configuration as they will need to make calls to it later.

By using a different port for each Element daemon we can run more than one node at a time on a single machine. This is useful for the purposes of our tutorial.

Before we start running our nodes we'll take the time to configure some aliases. These will let us pass commands to each node's client in a much simpler way. 

Before we do that it is worth noting how the set up of a daemon and client works for both Bitcoin and Elements.

The Bitcoin and Elements node software runs as a daemon - executing as a background service. The Bitcoin and Elements client software simplifies sending commands to the daemon and reading the output. 

Once our Bitcoin and Elements daemons have been started we will only interact with them through the use of a client, which will use the same config file as its associated daemon and therefore have permission to access it via RPC calls.

Now let's create the aliases for our Bitcoin and Elements daemons and clients: 

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
If you want to skip ahead to running a non-sidechain Elements based blockchain you can move to the [Elements as a standalone Blockchain]({{ site.url }}/elements-code-tutorial/blockchain) section now. You can return to the [Using Elements to perform basic operations]({{ site.url }}/elements-code-tutorial/basic-operations) section after that, but may have to amend the commands issued occasionally. 

##### NOTE: If you want to run a few examples of your own after following this tutorial you can add the aliases above to your ‘bashrc' file which means that they are available to you in every new terminal window. To do this run  nano ~/.bashrc  and paste the 6 alias lines above into the file, save and exit. The next time you open a new terminal window the aliases will have already been set. We don't need to do this now but if at any point you need to exit the tutorial and start again you can run the following commands to shut down the Bitcoin and Elements daemons first before restarting:

~~~~
b-cli stop
e1-cli stop
e2-cli stop
~~~~

If we try and execute a simple command using the Bitcoin RPC client it will error as we haven't started the daemon yet:

~~~~
b-cli -getinfo
~~~~

That gives us an error - but we were expecting that, so let's start the Bitcoin daemon:

~~~~
b-dae
~~~~

And now try again:

~~~~
b-cli -getinfo
~~~~

That should work! You should get a response telling you general information about the node.

So now we have the Bitcoin daemon running we need to start our Elements daemons. Remember that we are able to run two instances on the same machine because we are using different configuration files with different ports and RPC permissions set. Start the daemons:

~~~~
e1-dae
e2-dae
~~~~

Check this by running:

~~~~
ps -A | grep elementsd
~~~~

The two instances should show in the list.

[Next: Using Elements to perform basic operations]({{ site.url }}/elements-code-tutorial/basic-operations)

