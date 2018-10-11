---
layout: page
title: Installing Bitcoin
permalink: /elements-code-tutorial/easy-run-code
---

# Elements code tutorial

## An easy way to run the main tutorial code

Rather than have to copy and paste or type in each line of code in the tutorial you can use the code below.

Save the code in a file named "runtutorial.sh" and place it in your home directory.

To run this code just open a terminal in your home directory and run:

~~~~
bash runtutorial.sh
~~~~

Then press the return key to execute each line in turn.

* * *

##### NOTE: If you want to run some of the steps automatically and then have execution stop and wait for you to press enter before continuing one line at a time just move the 'trap read debug' statement near the top of the code so that it is above the line you want to stop at. Execution will then switch to one line at a time and wait for you to press return before executing the next command.<br/><br/>You will see that occasionally we will use the 'sleep' command to pause execution to allow the daemons to do things like stop, start and sync mempools etc.<br/><br/>It is perhaps a good idea to have the tutorial open as you run through this code as it explains what each step does and why.

 
~~~~
#!/bin/bash
set -x
trap read debug

# This code is based upon: https://github.com/ElementsProject/elements/tree/elements-0.14.1/contrib/assets_tutorial

################################
#
# Save this code in a file named runtutorial.sh and place it in your home directory.
#
# To run this code just open a terminal in your home directory and run:
# bash runtutorial.sh
#
# Then press the return key to execute each line in turn.
#
# If you want to run some of the steps automatically and then have execution stop 
# and wait for you to press enter before continuing one line at a time just move the 
# 'trap read debug' statement above so that it is above the line you want to stop at. 
# Execution will then switch to one line at a time and wait for you to press return
# before executing the next command.
# 
# You will see that occasionally we will use the 'sleep' command to pause execution 
# to allow the daemons to do things like stop, start and sync mempools etc.
#
#################################

# It is perhaps a good idea to have the tutorial open as you run through this code 
# as it explains what each step does and why.
# The code is based upon: https://github.com/ElementsProject/elements/tree/elements-0.14.1/contrib/assets_tutorial 

cd
cd elements
cd src

shopt -s expand_aliases

alias b-dae="bitcoind -datadir=$HOME/bitcoindir"
alias b-cli="bitcoin-cli -datadir=$HOME/bitcoindir"

alias e1-dae="$HOME/elements/src/elementsd -datadir=$HOME/elementsdir1"
alias e1-cli="$HOME/elements/src/elements-cli -datadir=$HOME/elementsdir1"

alias e2-dae="$HOME/elements/src/elementsd -datadir=$HOME/elementsdir2"
alias e2-cli="$HOME/elements/src/elements-cli -datadir=$HOME/elementsdir2"

echo "The following 3 lines may error - that is fine."

b-cli stop
e1-cli stop
e2-cli stop
sleep 5

cd

echo "The following 3 'rm' commands may error - that is fine."

rm -r ~/bitcoindir ; rm -r ~/elementsdir1 ; rm -r ~/elementsdir2
mkdir ~/bitcoindir ; mkdir ~/elementsdir1 ; mkdir ~/elementsdir2

cd elements
cd src

cp ~/elements/contrib/assets_tutorial/bitcoin.conf ~/bitcoindir/bitcoin.conf
cp ~/elements/contrib/assets_tutorial/elements1.conf ~/elementsdir1/elements.conf
cp ~/elements/contrib/assets_tutorial/elements2.conf ~/elementsdir2/elements.conf

b-dae

sleep 5

b-cli -getinfo

e1-dae
e2-dae

sleep 5

e1-cli getwalletinfo
e2-cli getwalletinfo

######## WALLET ###########

e1-cli sendtoaddress $(e1-cli getnewaddress) 21000000 "" "" true
e1-cli generate 101
sleep 5
e1-cli sendtoaddress $(e2-cli getnewaddress) 10500000 "" "" false
e1-cli generate 101
sleep 5

e1-cli getwalletinfo
e2-cli getwalletinfo

ADDR=$(e2-cli getnewaddress)

echo $ADDR

e2-cli validateaddress $ADDR

TXID=$(e2-cli sendtoaddress $ADDR 1)

sleep 2

e1-cli getrawmempool
e2-cli getrawmempool
e1-cli getinfo
e2-cli getinfo

e2-cli generate 1
sleep 2

e1-cli getrawmempool
e2-cli getrawmempool
e1-cli getinfo
e2-cli getinfo

e2-cli gettransaction $TXID

echo "This may error - that is ok"
e1-cli gettransaction $TXID

e1-cli getrawtransaction $TXID 1

e1-cli importprivkey $(e2-cli dumpprivkey $ADDR)

e1-cli gettransaction $TXID

e1-cli getwalletinfo

e1-cli listunspent 1 1

e1-cli importblindingkey $ADDR $(e2-cli dumpblindingkey $ADDR)

e1-cli getwalletinfo
e1-cli listunspent 1 1
e1-cli gettransaction $TXID

###### ASSETS #######

e1-cli getwalletinfo

e1-cli getwalletinfo bitcoin

e1-cli dumpassetlabels

e1-cli getwalletinfo b2e15d0d7a0c94e4e2ce0fe6e8691b9e451377f6e46e8045a86f7c4b5d4f0f23

ISSUE=$(e1-cli issueasset 100 1)

ASSET=$(echo $ISSUE | jq '.asset' | tr -d '"')
TOKEN=$(echo $ISSUE | jq '.token' | tr -d '"')
ITXID=$(echo $ISSUE | jq '.txid' | tr -d '"')
IVIN=$(echo $ISSUE | jq '.vin' | tr -d '"')

echo $ASSET

e1-cli listissuances

e1-cli stop
sleep 2
e1-dae -assetdir=$ASSET:demoasset
sleep 2
e1-cli listissuances

e1-cli generate 1

sleep 2
e2-cli listissuances

IADDR=$(e1-cli gettransaction $ITXID | jq '.details[0].address' | tr -d '"')

e2-cli importaddress $IADDR

e2-cli listissuances

ISSUEKEY=$(e1-cli dumpissuanceblindingkey $ITXID $IVIN)

e2-cli importissuanceblindingkey $ITXID $IVIN $ISSUEKEY

e2-cli listissuances

E2DEMOADD=$(e2-cli getnewaddress)
e1-cli sendtoaddress $E2DEMOADD 10 "" "" false demoasset
sleep 2
e1-cli generate 1
sleep 2

e2-cli getwalletinfo
e1-cli getwalletinfo

E1DEMOADD=$(e1-cli getnewaddress)
e2-cli sendtoaddress $E1DEMOADD 10 "" "" false $ASSET
sleep 2
e2-cli generate 1
sleep 2
e1-cli getwalletinfo
e2-cli getwalletinfo

RTRANS=$(e1-cli reissueasset $ASSET 99)
RTXID=$(echo $RTRANS | jq '.txid' | tr -d '"')

e1-cli listissuances $ASSET

e1-cli gettransaction $RTXID

e1-cli generate 1
sleep 2
RAWRTRANS=$(e2-cli getrawtransaction $RTXID)
e2-cli decoderawtransaction $RAWRTRANS

e1-cli getwalletinfo
e2-cli getwalletinfo

echo "This will error and that is expected:"
e2-cli reissueasset $ASSET 10

RITRECADD=$(e2-cli getnewaddress)
e1-cli sendtoaddress $RITRECADD 1 "" "" false $TOKEN
e1-cli generate 1
sleep 2
e1-cli getwalletinfo
e2-cli getwalletinfo

RISSUE=$(e2-cli reissueasset $ASSET 10)
e2-cli getwalletinfo

e2-cli generate 1
sleep 2

e1-cli listissuances

RITXID=$(echo $RISSUE | jq '.txid' | tr -d '"')
RIADDR=$(e2-cli gettransaction $RITXID | jq '.details[0].address' | tr -d '"')

e1-cli importaddress $RIADDR
e1-cli listissuances

UBRISSUE=$(e2-cli issueasset 55 1 false)

UBASSET=$(echo $UBRISSUE | jq '.asset' | tr -d '"')

e2-cli getwalletinfo

e2-cli generate 1
sleep 2
UBRITXID=$(echo $UBRISSUE | jq '.txid' | tr -d '"')

UBRIADDR=$(e2-cli gettransaction $UBRITXID | jq '.details[0].address' | tr -d '"')

e1-cli importaddress $UBRIADDR

e1-cli listissuances

e2-cli destroyamount $UBASSET 5
e2-cli getwalletinfo

###### BLOCKSIGNING #######

e1-cli generate 1
sleep 2

ADDR1=$(e1-cli getnewaddress)
ADDR2=$(e2-cli getnewaddress)

VALID1=$(e1-cli validateaddress $ADDR1)
PUBKEY1=$(echo $VALID1 | jq '.pubkey' | tr -d '"')

VALID2=$(e2-cli validateaddress $ADDR2)
PUBKEY2=$(echo $VALID2 | jq '.pubkey' | tr -d '"')

KEY1=$(e1-cli dumpprivkey $ADDR1)
KEY2=$(e2-cli dumpprivkey $ADDR2)

MULTISIG=$(e1-cli createmultisig 2 '''["'''$PUBKEY1'''", "'''$PUBKEY2'''"]''')
REDEEMSCRIPT=$(echo $MULTISIG | jq '.redeemScript' | tr -d '"')
echo $REDEEMSCRIPT

e1-cli stop
e2-cli stop
sleep 5

SIGNBLOCKARG="-signblockscript=$(echo $REDEEMSCRIPT)"

rm -r ~/elementsdir1/elementsregtest/blocks
rm -r ~/elementsdir1/elementsregtest/chainstate
rm ~/elementsdir1/elementsregtest/wallet.dat
rm -r ~/elementsdir2/elementsregtest/blocks
rm -r ~/elementsdir2/elementsregtest/chainstate
rm ~/elementsdir2/elementsregtest/wallet.dat

e1-dae $SIGNBLOCKARG
e2-dae $SIGNBLOCKARG

sleep 5

e1-cli importprivkey $KEY1
e2-cli importprivkey $KEY2

echo "This will error - that is ok:"
e1-cli generate 1
e2-cli generate 1

HEX=$(e1-cli getnewblockhex)

e1-cli getblockcount
 
e1-cli submitblock $HEX

e1-cli getblockcount

SIGN1=$(e1-cli signblock $HEX)
SIGN2=$(e2-cli signblock $HEX)

BLOCKRESULT=$(e1-cli combineblocksigs $HEX '''["'''$SIGN1'''", "'''$SIGN2'''"]''')

COMPLETE=$(echo $BLOCKRESULT | jq '.complete' | tr -d '"')

SIGNBLOCK=$(echo $BLOCKRESULT | jq '.hex' | tr -d '"')

echo $COMPLETE

e2-cli submitblock $SIGNBLOCK

e1-cli getblockcount
e2-cli getblockcount

e1-cli stop
e2-cli stop
sleep 5

######## Pegging #######

rm -r ~/elementsdir1/elementsregtest/blocks
rm -r ~/elementsdir1/elementsregtest/chainstate
rm ~/elementsdir1/elementsregtest/wallet.dat
rm -r ~/elementsdir2/elementsregtest/blocks
rm -r ~/elementsdir2/elementsregtest/chainstate
rm ~/elementsdir2/elementsregtest/wallet.dat

FEDPEGARG="-fedpegscript=5221$(echo $PUBKEY1)21$(echo $PUBKEY2)52ae"

e1-dae $FEDPEGARG
e2-dae $FEDPEGARG
sleep 5

e1-cli generate 101
b-cli generate 101

e1-cli getpeginaddress
e1-cli getpeginaddress

ADDRS=$(e1-cli getpeginaddress)

MAINCHAIN=$(echo $ADDRS |  jq '.mainchain_address' | tr -d '"')
SIDECHAIN=$(echo $ADDRS | jq '.claim_script' | tr -d '"')

b-cli getwalletinfo

TXID=$(b-cli sendtoaddress $MAINCHAIN 1)

b-cli getwalletinfo

b-cli generate 101

b-cli getwalletinfo

PROOF=$(b-cli gettxoutproof '''["'''$TXID'''"]''')
RAW=$(b-cli getrawtransaction $TXID)

CLAIMTXID=$(e1-cli claimpegin $RAW $PROOF)

e2-cli generate 1
sleep 2

e1-cli getrawtransaction $CLAIMTXID 1

e1-cli getwalletinfo


#### Pegging Out ####

e1-cli sendtomainchain $(b-cli getnewaddress) 10

e1-cli generate 1

e1-cli getwalletinfo

e1-cli stop
e2-cli stop
b-cli stop
sleep 5
echo “Completed with no errors”
~~~~

* * *

