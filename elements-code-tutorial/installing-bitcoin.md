---
layout: page
title: Installing Bitcoin
permalink: /elements-code-tutorial/installing-bitcoin
---

# Elements code tutorial

## Installing Bitcoin

First we will install Bitcoin on the machine. This will allow us to demonstrate how the Federated 2-Way Peg works in Elements later on in the tutorial. It is not required if you intend to use Elements as a standalone blockchain, but to fully understand the features available in Elements it is a good idea to follow along anyway. It doesn’t take long to install Bitcoin using the commands below and we will be running in "regtest" mode, so there is no blockchain to sync.

You can download the compiled libraries from the [Bitcoin Core Download page](https://bitcoincore.org/en/download/). Make sure you'll download the binaries, the `SHA256SUM` and `SHA256SUM.asc`-file.

Open a terminal and change your directory the one of the files mentioned above. If you are using `~Downloads` this becomes

~~~
cd ~/Downloads/
~~~

Verify the hash

~~~
sha256sum --ignore-missing --check SHA256SUMS
~~~

Verify that the files are signed by a bitcoin-core maintainer.

~~~
gpg --keyserver hkps://keys.openpgp.org --recv-keys E777299FC265DD04793070EB944D35F9AC3DB76A
gpg --verify SHA256SUMS.asc
~~~

You might get a warning that the key is not certified. This means that to fully verify your download you should also confirm the signing key's fingerprint.

The next step is to install the actual binaries. Assuming you are installing version 22.0 this becomes.

~~~
tar -xf <path-to-downloaded-file>.tar.gz
sudo install -m 0755 -o root -g root -t /usr/local/bin bitcoin-22.0/bin/*
~~~

Check that the install worked:

~~~
which bitcoind
~~~~

Which should return:

<div class="console-output">/usr/bin/bitcoind</div>

That means that the Bitcoin software was installed. We will run it later so let’s move on with getting Elements set up.


[Next: Installing Elements]({{ site.url }}/elements-code-tutorial/installing-elements)

