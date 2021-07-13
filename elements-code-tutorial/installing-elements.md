---
layout: page
title: Installing Elements
permalink: /elements-code-tutorial/installing-elements
---

# Elements code tutorial

## Installing Elements

As we will be building Elements from the source code, we first need to pull the code from the [GitHub repository](https://github.com/elementsproject/elements) where it is maintained. We'll use "git" for this. You can ignore installation steps within this tutorial for software that you already have. If you do not have git installed (you can check by running the ``git --version`` command) you can install it using:

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

##### Note: Some lines wrap in the text below. Each line starting with "sudo apt-get" should be executed in its entirety. The first command will update your existing packages list, enabling you to install all the required dependancies.

~~~~
sudo apt-get update
sudo apt-get install build-essential libtool autotools-dev autoconf pkg-config libssl-dev
sudo apt-get install libboost-all-dev
sudo apt-get install libqt5gui5 libqt5core5a libqt5dbus5 qttools5-dev qttools5-dev-tools libprotobuf-dev protobuf-compiler imagemagick librsvg2-bin
sudo apt-get install libqrencode-dev autoconf openssl libssl-dev libevent-dev
sudo apt-get install libminiupnpc-dev
sudo apt install jq
~~~~

Now we need to build and install the Berkeley database.

##### Note: You **must** replace ``/home/yourusername`` below with the location of your home directory. Again, note that some lines wrap in the text below.

~~~~
mkdir bdb4
wget 'http://download.oracle.com/berkeley-db/db-4.8.30.NC.tar.gz'
tar -xzvf db-4.8.30.NC.tar.gz
sed -i 's/__atomic_compare_exchange/__atomic_compare_exchange_db/g' db-4.8.30.NC/dbinc/atomic.h
cd db-4.8.30.NC/build_unix/
../dist/configure --enable-cxx --disable-shared --with-pic --prefix=/home/yourusername/bdb4/
make install
~~~~

Now let's configure, compile and install Elements.

##### Note: You **must** replace ``/home/yourusername`` below (which occurs twice) with the location of your home directory.

~~~~
cd
cd elements
./autogen.sh
./configure LDFLAGS="-L/home/yourusername/bdb4/lib/" CPPFLAGS="-I/home/yourusername/bdb4/include/"
make
sudo make install
~~~~

The "make" command may take a while to complete as it will also run the Elements test-suite as part of the build process.

Check that the install worked:

~~~~
which elementsd
~~~~

Which should return:

<div class="console-output">/usr/local/bin/elementsd</div>

If you are using a Virtual Machine, now would be a good point to take a snapshot of the machine's state as we have set up Bitcoin and Elements ready for use. 

[Next: Setting up your working environment]({{ site.url }}/elements-code-tutorial/working-environment)

