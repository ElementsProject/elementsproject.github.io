---
layout: page
title: Installing Elemets
permalink: /elements-code-tutorial/installing-elements
---

# Elements code tutorial

## Installing Elements

As we will be building Elements from the source code, we first need to pull the code from the [GitHub repository](https://github.com/elementsproject/elements) where it is maintained. We'll use "git" for this, which we will now install. You can ignore installation steps within this tutorial for software that you already have.

~~~~
sudo apt install git
~~~~

Now pull the code from the repository to your machine (after moving to your home directory):

~~~~
cd
git clone https://github.com/ElementsProject/elements.git
~~~~

That's pulled all the code from the Elements repository into a newly created directory in Home called "elements". 

Before we can compile and install Elements, we need to install software that the build process and this tutorial is dependant upon. Run the following terminal commands in turn. You will need to enter "y" when prompted for some of the commands. The most up to date set of dependencies for Ubuntu can be found [here](https://github.com/ElementsProject/elements/blob/master/doc/build-unix.md) and others within the relevant 'build-*.md' file [here](https://github.com/ElementsProject/elements/tree/master/doc). 

##### Note: Some lines wrap in the text below. Each line starting with "sudo apt-get" should be executed in its entirety.

~~~~
sudo apt-get install build-essential libtool autotools-dev autoconf pkg-config libssl-dev
sudo apt-get install libboost-all-dev
sudo apt-get install libqt5gui5 libqt5core5a libqt5dbus5 qttools5-dev qttools5-dev-tools libprotobuf-dev protobuf-compiler imagemagick librsvg2-bin
sudo apt-get install libqrencode-dev autoconf openssl libssl-dev libevent-dev
sudo apt-get install libminiupnpc-dev
sudo apt-get install libdb4.8-dev libdb4.8++-dev
sudo apt install jq
~~~~

Move into the Elements directory:

~~~~
cd elements
~~~~

Now let's configure, compile and install Elements. If you get a permission denied error when running the final command use "sudo make install" instead.

##### Note: The "make" command may take a while to complete as it will also run the Elements test-suite as part of the build process.

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

Which should return:

<div class="console-output">/usr/local/bin/elementsd</div>

If you are using a Virtual Machine, now would be a good point to take a snapshot of the machine's state as we have set up Bitcoin and Elements ready for use. 

[Next: Setting up your working environment]({{ site.url }}/elements-code-tutorial/working-environment)

