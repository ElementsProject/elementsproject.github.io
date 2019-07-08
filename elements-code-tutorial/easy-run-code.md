---
layout: page
title: Elements easy-run code
permalink: /elements-code-tutorial/easy-run-code
---

# Elements code tutorial

## An easy way to run the main tutorial's code

Rather than have to copy and paste or type in each line of code in the tutorial, you can use the code below.

Save the code in a file named **runtutorial.sh** and place it in your home directory.

To run this code just open a terminal in your $HOME directory and run:

~~~~
bash runtutorial.sh
~~~~

Then press the return key to execute each line in turn.

More advanced examples, like manual 'raw' issuance of an asset, can be found [here]({{ site.url }}/elements-code-tutorial/advanced-examples).

* * *

##### Note: If you want to run some of the steps automatically and then have execution stop and wait for you to press enter before continuing one line at a time: move the **trap read debug** statement down so that it is above the line you want to stop at. Execution will run each line automatically and stop when that line is reached. It will then switch to executing one line at a time, waiting for you to press return before executing the next command.<br/><br/>You will see that occasionally we will use the **sleep** command to pause execution. This allows the daemons time to do things like stop, start and sync mempools.<br/><br/>It is perhaps a good idea to have the relevant tutorial pages open as you run through this code for reference, as it is not itself annotated in any meaningful way.<br/><br/>There is a chance that the " and ' characters may not paste into your **runtutorial.sh** file correctly, so type over them yourself if you come across any issues executing the code.
 
~~~~
#!/bin/bash
set -x

# This code is based upon the Python example at https://github.com/ElementsProject/elements found within the contrib/assets_tutorial folder

################################
#
# Save this code in a file named runtutorial.sh and place it in your home directory.
#
# To run this code just open a terminal in your home directory and run:
# bash runtutorial.sh
#
# If you want to run some of the steps automatically and then have execution stop
# and wait for you to press enter before continuing one line at a time: uncomment 
# and move the 'trap read debug' statement down so that it is above the line you want
# to stop at. Execution will run each line automatically and stop when that line is 
# reached. It will then switch to executing one line at a time, waiting for you 
# to press return before executing the next command.
# 
# You will see that occasionally we will use the **sleep** command to pause execution.
# This allows the daemons time to do things like stop, start and sync mempools. You
# can probably decrease the sleeps without issue. The numbers used below are so a low
# powered machine like a Raspberry Pi can run without incident.
# 
# It is perhaps a good idea to have the relevant tutorial pages open as you run 
# through this code for reference, as it is not itself annotated in any meaningful way.
#
# There is a chance that the " and ' characters may not paste into your 
# **runtutorial.sh** file correctly, so type over them yourself if you come 
# across any issues executing the code.
#
#################################

# Remove to run without stopping:
# trap read debug

shopt -s expand_aliases

alias b-dae="bitcoind -datadir=$HOME/bitcoindir"
alias b-cli="bitcoin-cli -datadir=$HOME/bitcoindir"

alias e1-dae="$HOME/elements/src/elementsd -datadir=$HOME/elementsdir1"
alias e1-cli="$HOME/elements/src/elements-cli -datadir=$HOME/elementsdir1"

alias e2-dae="$HOME/elements/src/elementsd -datadir=$HOME/elementsdir2"
alias e2-cli="$HOME/elements/src/elements-cli -datadir=$HOME/elementsdir2"

echo "The following 3 lines may error - that is fine."

# Ignore error
set +o errexit

b-cli stop
e1-cli stop
e2-cli stop
sleep 15

echo "The following 3 'rm' commands may error - that is fine."

rm -r ~/bitcoindir ; rm -r ~/elementsdir1 ; rm -r ~/elementsdir2
mkdir ~/bitcoindir ; mkdir ~/elementsdir1 ; mkdir ~/elementsdir2

cp ~/elements/contrib/assets_tutorial/bitcoin.conf ~/bitcoindir/bitcoin.conf
cp ~/elements/contrib/assets_tutorial/elements1.conf ~/elementsdir1/elements.conf
cp ~/elements/contrib/assets_tutorial/elements2.conf ~/elementsdir2/elements.conf

b-dae

sleep 10

# Wait for bitcoin node to finish startup and respond to commands
until b-cli getwalletinfo
do
  echo "Waiting for bitcoin node to finish loading..."
  sleep 2
done

e1-dae
e2-dae

sleep 10

# Wait for e1 node to finish startup and respond to commands
until e1-cli getwalletinfo
do
  echo "Waiting for e1 to finish loading..."
  sleep 2
done

# Wait for e2 node to finish startup and respond to commands
until e2-cli getwalletinfo
do
  echo "Waiting for e2 to finish loading..."
  sleep 2
done

# Exit on error
set -o errexit

### Basic Operations ###

ADDRGENB=$(b-cli getnewaddress)
ADDRGEN1=$(e1-cli getnewaddress)
ADDRGEN2=$(e2-cli getnewaddress)

e1-cli sendtoaddress $(e1-cli getnewaddress) 21000000 "" "" true
e1-cli generatetoaddress 101 $ADDRGEN1
sleep 10
e1-cli sendtoaddress $(e2-cli getnewaddress) 10500000 "" "" false
e1-cli generatetoaddress 101 $ADDRGEN1
sleep 10

e1-cli getwalletinfo
e2-cli getwalletinfo

ADDR=$(e2-cli getnewaddress)

echo $ADDR

e2-cli getaddressinfo $ADDR

TXID=$(e2-cli sendtoaddress $ADDR 1)

sleep 10

e1-cli getrawmempool
e2-cli getrawmempool
e1-cli getblockcount
e2-cli getblockcount

e2-cli generatetoaddress 1 $ADDRGEN2
sleep 10

e1-cli getrawmempool
e2-cli getrawmempool
e1-cli getblockcount
e2-cli getblockcount

e2-cli gettransaction $TXID

# Ignore error
set +o errexit

echo "This may error - that is ok, not aware of tx"
e1-cli gettransaction $TXID

# Exit on error
set -o errexit

e1-cli getrawtransaction $TXID 1

e1-cli importaddress $ADDR

e1-cli gettransaction $TXID true

e1-cli importblindingkey $ADDR $(e2-cli dumpblindingkey $ADDR)

e1-cli gettransaction $TXID true

### Issued Assets ###

e1-cli getwalletinfo

e1-cli dumpassetlabels

ISSUE=$(e1-cli issueasset 100 1)

ASSET=$(echo $ISSUE | jq '.asset' | tr -d '"')
TOKEN=$(echo $ISSUE | jq '.token' | tr -d '"')
ITXID=$(echo $ISSUE | jq '.txid' | tr -d '"')
IVIN=$(echo $ISSUE | jq '.vin' | tr -d '"')

echo $ASSET

e1-cli listissuances

e1-cli stop
sleep 15
e1-dae -assetdir=$ASSET:demoasset
sleep 10

# Ignore error
set +o errexit

# Wait for e1 node to finish startup and respond to commands
until e1-cli getwalletinfo
do
  echo "Waiting for e1 to finish loading..."
  sleep 2
done

# Exit on error
set -o errexit

e1-cli listissuances

e1-cli generatetoaddress 1 $ADDRGEN1

sleep 10
e2-cli listissuances

IADDR=$(e1-cli gettransaction $ITXID | jq '.details[0].address' | tr -d '"')

e2-cli importaddress $IADDR

e2-cli listissuances

ISSUEKEY=$(e1-cli dumpissuanceblindingkey $ITXID $IVIN)

e2-cli importissuanceblindingkey $ITXID $IVIN $ISSUEKEY

e2-cli listissuances

E2DEMOADD=$(e2-cli getnewaddress)
e1-cli sendtoaddress $E2DEMOADD 10 "" "" false false 1 UNSET demoasset
sleep 10
e1-cli generatetoaddress 1 $ADDRGEN1
sleep 10

e2-cli getwalletinfo
e1-cli getwalletinfo

E1DEMOADD=$(e1-cli getnewaddress)
e2-cli sendtoaddress $E1DEMOADD 10 "" "" false false 1 UNSET $ASSET
sleep 10
e2-cli generatetoaddress 1 $ADDRGEN2
sleep 10
e1-cli getwalletinfo
e2-cli getwalletinfo

### Reissuing Assets ###

RTRANS=$(e1-cli reissueasset $ASSET 99)
RTXID=$(echo $RTRANS | jq '.txid' | tr -d '"')

e1-cli listissuances $ASSET

e1-cli gettransaction $RTXID

e1-cli generatetoaddress 1 $ADDRGEN1
sleep 10
RAWRTRANS=$(e2-cli getrawtransaction $RTXID)
e2-cli decoderawtransaction $RAWRTRANS

e1-cli getwalletinfo
e2-cli getwalletinfo

# Ignore error
set +o errexit

echo "This will error and that is expected:"
e2-cli reissueasset $ASSET 10

# Exit on error
set -o errexit

RITRECADD=$(e2-cli getnewaddress)
e1-cli sendtoaddress $RITRECADD 1 "" "" false false 1 UNSET $TOKEN
e1-cli generatetoaddress 1 $ADDRGEN1
sleep 10
e1-cli getwalletinfo
e2-cli getwalletinfo

RISSUE=$(e2-cli reissueasset $ASSET 10)
e2-cli getwalletinfo

e2-cli generatetoaddress 1 $ADDRGEN2
sleep 10

e1-cli listissuances

RITXID=$(echo $RISSUE | jq '.txid' | tr -d '"')
RIADDR=$(e2-cli gettransaction $RITXID | jq '.details[0].address' | tr -d '"')

e1-cli importaddress $RIADDR
e1-cli listissuances

UBRISSUE=$(e2-cli issueasset 55 1 false)

UBASSET=$(echo $UBRISSUE | jq '.asset' | tr -d '"')

e2-cli getwalletinfo

e2-cli generatetoaddress 1 $ADDRGEN2
sleep 10

e1-cli listissuances

UBRITXID=$(echo $UBRISSUE | jq '.txid' | tr -d '"')

UBRIADDR=$(e2-cli gettransaction $UBRITXID | jq '.details[0].address' | tr -d '"')

e1-cli importaddress $UBRIADDR

e1-cli listissuances

e2-cli destroyamount $UBASSET 5
e2-cli getwalletinfo

### Block Signing ###

e1-cli generatetoaddress 1 $ADDRGEN1
sleep 10

ADDR1=$(e1-cli getnewaddress)
ADDR2=$(e2-cli getnewaddress)

VALID1=$(e1-cli getaddressinfo $ADDR1)
PUBKEY1=$(echo $VALID1 | jq '.pubkey' | tr -d '"')

VALID2=$(e2-cli getaddressinfo $ADDR2)
PUBKEY2=$(echo $VALID2 | jq '.pubkey' | tr -d '"')

KEY1=$(e1-cli dumpprivkey $ADDR1)
KEY2=$(e2-cli dumpprivkey $ADDR2)

MULTISIG=$(e1-cli createmultisig 2 '''["'''$PUBKEY1'''", "'''$PUBKEY2'''"]''')
REDEEMSCRIPT=$(echo $MULTISIG | jq '.redeemScript' | tr -d '"')
echo $REDEEMSCRIPT

e1-cli stop
e2-cli stop
sleep 15

SIGNBLOCKARG="-signblockscript=$(echo $REDEEMSCRIPT) -con_max_block_sig_size=150"

rm -r ~/elementsdir1/elementsregtest/blocks
rm -r ~/elementsdir1/elementsregtest/chainstate
rm ~/elementsdir1/elementsregtest/wallets/wallet.dat
rm -r ~/elementsdir2/elementsregtest/blocks
rm -r ~/elementsdir2/elementsregtest/chainstate
rm ~/elementsdir2/elementsregtest/wallets/wallet.dat

e1-dae $SIGNBLOCKARG
e2-dae $SIGNBLOCKARG

sleep 10

# Ignore error
set +o errexit

# Wait for e1 node to finish startup and respond to commands
until e1-cli getwalletinfo
do
  echo "Waiting for e1 to finish loading..."
  sleep 2
done

# Wait for e2 node to finish startup and respond to commands
until e2-cli getwalletinfo
do
  echo "Waiting for e2 to finish loading..."
  sleep 2
done

# Exit on error
set -o errexit

e1-cli importprivkey $KEY1
e2-cli importprivkey $KEY2

ADDRGEN1=$(e1-cli getnewaddress)
ADDRGEN2=$(e2-cli getnewaddress)

# Ignore error
set +o errexit

echo "This will error - that is ok:"
e1-cli generatetoaddress 1 $ADDRGEN1
e2-cli generatetoaddress 1 $ADDRGEN2

# Exit on error
set -o errexit

HEX=$(e1-cli getnewblockhex)

e1-cli getblockcount
 
e1-cli submitblock $HEX

e1-cli getblockcount

SIGN1=$(e1-cli signblock $HEX)
SIGN2=$(e2-cli signblock $HEX)

SIGN1DATA=$(echo $SIGN1 | jq '.[0]')
SIGN2DATA=$(echo $SIGN2 | jq '.[0]')

COMBINED=$(e1-cli combineblocksigs $HEX "[$SIGN1DATA,$SIGN2DATA]")

SIGNEDBLOCK=$(echo $COMBINED | jq '.hex' | tr -d '"')

e2-cli submitblock $SIGNEDBLOCK

e1-cli getblockcount
e2-cli getblockcount

e1-cli stop
e2-cli stop
sleep 15

### Sidechain - Peg-In ###

rm -r ~/elementsdir1/elementsregtest/blocks
rm -r ~/elementsdir1/elementsregtest/chainstate
rm ~/elementsdir1/elementsregtest/wallets/wallet.dat
rm -r ~/elementsdir2/elementsregtest/blocks
rm -r ~/elementsdir2/elementsregtest/chainstate
rm ~/elementsdir2/elementsregtest/wallets/wallet.dat

FEDPEGARG="-fedpegscript=5221$(echo $PUBKEY1)21$(echo $PUBKEY2)52ae"

e1-dae $FEDPEGARG
e2-dae $FEDPEGARG
sleep 10

# Ignore error
set +o errexit

# Wait for e1 node to finish startup and respond to commands
until e1-cli getwalletinfo
do
  echo "Waiting for e1 to finish loading..."
  sleep 2
done

# Wait for e2 node to finish startup and respond to commands
until e2-cli getwalletinfo
do
  echo "Waiting for e2 to finish loading..."
  sleep 2
done

# Exit on error
set -o errexit

ADDRGEN1=$(e1-cli getnewaddress)
ADDRGEN2=$(e2-cli getnewaddress)

e1-cli generatetoaddress 101 $ADDRGEN1
b-cli generatetoaddress 101 $ADDRGENB

e1-cli getpeginaddress
e1-cli getpeginaddress

ADDRS=$(e1-cli getpeginaddress)

MAINCHAIN=$(echo $ADDRS |  jq '.mainchain_address' | tr -d '"')
SIDECHAIN=$(echo $ADDRS | jq '.claim_script' | tr -d '"')

b-cli getwalletinfo

TXID=$(b-cli sendtoaddress $MAINCHAIN 1)

b-cli getwalletinfo

b-cli generatetoaddress 101 $ADDRGENB

b-cli getwalletinfo

PROOF=$(b-cli gettxoutproof '''["'''$TXID'''"]''')
RAW=$(b-cli getrawtransaction $TXID)

CLAIMTXID=$(e1-cli claimpegin $RAW $PROOF)

e2-cli generatetoaddress 1 $ADDRGEN2
sleep 10

e1-cli getrawtransaction $CLAIMTXID 1

e1-cli getwalletinfo

### Sidechain - Peg-Out ###

e1-cli sendtomainchain $(b-cli getnewaddress) 10

e1-cli generatetoaddress 1 $ADDRGEN1

e1-cli getwalletinfo

e1-cli stop
e2-cli stop
b-cli stop
sleep 15

### Standalone Blockchain ###

rm -r ~/elementsdir1/elementsregtest/blocks
rm -r ~/elementsdir1/elementsregtest/chainstate
rm ~/elementsdir1/elementsregtest/wallets/wallet.dat
rm -r ~/elementsdir2/elementsregtest/blocks
rm -r ~/elementsdir2/elementsregtest/chainstate
rm ~/elementsdir2/elementsregtest/wallets/wallet.dat

STANDALONEARGS="-validatepegin=0 -defaultpeggedassetname=newasset -initialfreecoins=100000000000000 -initialreissuancetokens=200000000"

e1-dae $STANDALONEARGS
e2-dae $STANDALONEARGS
sleep 10

# Ignore error
set +o errexit

# Wait for e1 node to finish startup and respond to commands
until e1-cli getwalletinfo
do
  echo "Waiting for e1 to finish loading..."
  sleep 2
done

# Wait for e2 node to finish startup and respond to commands
until e2-cli getwalletinfo
do
  echo "Waiting for e2 to finish loading..."
  sleep 2
done

# Exit on error
set -o errexit

ADDRGEN1=$(e1-cli getnewaddress)
ADDRGEN2=$(e2-cli getnewaddress)

e1-cli getwalletinfo
e2-cli getwalletinfo

DEFAULTRIT=$(e1-cli getwalletinfo | jq '[.balance] | .[0]' | jq 'keys | .[0]' | tr -d '"')

if [ $DEFAULTRIT = "newasset" ]; then
  DEFAULTRIT=$(e1-cli getwalletinfo | jq '[.balance] | .[0]' | jq 'keys | .[1]' | tr -d '"')
fi

echo $DEFAULTRIT

e1-cli sendtoaddress $(e1-cli getnewaddress) 1000000 "" "" true

e1-cli sendtoaddress $(e1-cli getnewaddress) 2 "" "" false false 1 UNSET $DEFAULTRIT
e1-cli generatetoaddress 101 $ADDRGEN1
sleep 10

e1-cli sendtoaddress $(e2-cli getnewaddress) 500 "" "" false 

e1-cli sendtoaddress $(e2-cli getnewaddress) 1 "" "" false false 1 UNSET $DEFAULTRIT
e1-cli generatetoaddress 101 $ADDRGEN1
sleep 10

e1-cli getwalletinfo
e2-cli getwalletinfo

e1-cli reissueasset newasset 100
e1-cli generatetoaddress 101 $ADDRGEN1
sleep 10

e1-cli getwalletinfo

e2-cli reissueasset newasset 100
e2-cli generatetoaddress 101 $ADDRGEN2
sleep 10

e1-cli getwalletinfo
e2-cli getwalletinfo

e1-cli stop
e2-cli stop
sleep 10

echo "Completed without error"
~~~~

* * *

More advanced examples, like manual 'raw' issuance of an asset, can be found [here]({{ site.url }}/elements-code-tutorial/advanced-examples).


