---
layout: page
title: Installing Elemets
permalink: /elements-code-tutorial/installing-elements
---

# Elements code tutorial

## Installing Elements

As we will be building Elements from the source code we first need to pull the code from the GitHub repository where it is kept. We'll use "git" for this which we will now install:

~~~~
sudo apt install git
~~~~

Now pull the code from the repository to your machine (after moving to your home directory):

~~~~
cd
git clone https://github.com/ElementsProject/elements.git
~~~~

That's pulled all the code from the Elements repository into a newly created directory in Home called "elements". 

##### NOTE: To update the code with any changes made to it in the future you can move into the home/elements folder and run "git pull origin elements-0.14.1" then "make" and then "make install". If you get a permissions error when updating then run "sudo make install". You may need to change the branch from "-0.14.1" in the above command if this changes in the future. 

Before we can compile and install Elements we need to install software that the build is dependant on. Run the following terminal commands in turn. You will need to enter "y" when prompted for some of the commands: (note some lines wrap in the text below - each entry starting "sudo apt-get" should be executed in its entirety.

~~~~
sudo apt-get install build-essential libtool autotools-dev autoconf pkg-config libssl-dev
sudo apt-get install libboost-all-dev
sudo apt-get install libqt5gui5 libqt5core5a libqt5dbus5 qttools5-dev qttools5-dev-tools libprotobuf-dev protobuf-compiler
sudo apt-get install libqrencode-dev autoconf openssl libssl-dev libevent-dev
sudo apt-get install libminiupnpc-dev
sudo apt-get install libdb4.8-dev libdb4.8++-dev
sudo apt install jq
~~~~

Move into the Elements directory:

~~~~
cd elements
~~~~

Now let's configure, compile and install Elements. If you get a permission denied error when running the final command use the 'sudo make install' command instead.

##### NOTE: The "make" command may take a while to complete!

~~~~
./autogen.sh
./configure
make
make install
~~~~

Check that the install worked:

~~~~
which elementsd
~~~~

This should return:

<div class="console-output">/usr/local/bin/elementsd</div>

If you are using a Virtual Machine, now would be a good point to snapshot the machine as we have now set up Bitcoin and Elements and will begin the code walkthrough. If you would like to run Elements as a standalone blockchain please see the [Elements as a standalone Blockchain]({{ site.url }}/elements-code-tutorial/blockchain) section for instructions, after following through the [Setting up your working environment]({{ site.url }}/elements-code-tutorial/working-environment) section.


[Next: Setting up your working environment]({{ site.url }}/elements-code-tutorial/working-environment)

