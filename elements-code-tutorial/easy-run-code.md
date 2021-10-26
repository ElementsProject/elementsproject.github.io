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
# Press Ctrl + c to stop script execution
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

echo "regtest=1
txindex=1
daemon=1
rpcuser=user3
rpcpassword=password3
fallbackfee=0.0002
[regtest]
rpcport=18888
port=18889

" > ~/bitcoindir/bitcoin.conf

echo "chain=elementsregtest
rpcuser=user1
rpcpassword=password1
daemon=1
server=1
listen=1
txindex=1
validatepegin=1
mainchainrpcport=18888
mainchainrpcuser=user3
mainchainrpcpassword=password3
initialfreecoins=2100000000000000
fallbackfee=0.0002
[elementsregtest]
rpcport=18884
port=18886
anyonecanspendaremine=1
connect=localhost:18887

" > ~/elementsdir1/elements.conf

echo "chain=elementsregtest
rpcuser=user2
rpcpassword=password2
daemon=1
server=1
listen=1
txindex=1
mainchainrpcport=18888
mainchainrpcuser=user3
mainchainrpcpassword=password3
initialfreecoins=2100000000000000
fallbackfee=0.0002
[elementsregtest]
rpcport=18885
port=18887
anyonecanspendaremine=1
connect=localhost:18886

" > ~/elementsdir2/elements.conf

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

# Create new wallets
e1-cli createwallet ""
e2-cli createwallet ""
e1-cli rescanblockchain
e2-cli rescanblockchain

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

ASSET=$(echo $ISSUE | jq -r '.asset')
TOKEN=$(echo $ISSUE | jq -r '.token')
ITXID=$(echo $ISSUE | jq -r '.txid')
IVIN=$(echo $ISSUE | jq -r '.vin')

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

IADDR=$(e1-cli gettransaction $ITXID | jq -r '.details[0].address')
e2-cli importaddress $IADDR

# Or if the address is not known to e2 but the TXID is (requires index=1 in config file to work):
#ISSUE_RAW_TX=$(e2-cli getrawtransaction $ITXID 1)
#ISSUE_VOUTS=$(echo $ISSUE_RAW_TX | jq -r '.vout')
#VOUT_ADDRESS_ISSUE=$(echo $ISSUE_VOUTS | jq -r '.[0].scriptPubKey.addresses[0]')
#e2-cli importaddress $VOUT_ADDRESS_ISSUE

e2-cli listissuances

ISSUEKEY=$(e1-cli dumpissuanceblindingkey $ITXID $IVIN)

e2-cli importissuanceblindingkey $ITXID $IVIN $ISSUEKEY

e2-cli listissuances

E2DEMOADD=$(e2-cli getnewaddress)
e1-cli sendtoaddress $E2DEMOADD 10 "" "" false false 1 UNSET false demoasset
sleep 10
e1-cli generatetoaddress 1 $ADDRGEN1
sleep 10

e2-cli getwalletinfo
e1-cli getwalletinfo

E1DEMOADD=$(e1-cli getnewaddress)
e2-cli sendtoaddress $E1DEMOADD 10 "" "" false false 1 UNSET false $ASSET
sleep 10
e2-cli generatetoaddress 1 $ADDRGEN2
sleep 10
e1-cli getwalletinfo
e2-cli getwalletinfo

### Reissuing Assets ###

RTRANS=$(e1-cli reissueasset $ASSET 99)
RTXID=$(echo $RTRANS | jq -r '.txid')

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
e1-cli sendtoaddress $RITRECADD 1 "" "" false false 1 UNSET false $TOKEN
e1-cli generatetoaddress 1 $ADDRGEN1
sleep 10
e1-cli getwalletinfo
e2-cli getwalletinfo

RISSUE=$(e2-cli reissueasset $ASSET 10)
e2-cli getwalletinfo

e2-cli generatetoaddress 1 $ADDRGEN2
sleep 10

e1-cli listissuances

RITXID=$(echo $RISSUE | jq -r '.txid')
RIADDR=$(e2-cli gettransaction $RITXID | jq -r '.details[0].address')
e1-cli importaddress $RIADDR

# Or if the address is not known to e1 but the TXID is (requires index=1 in config file to work):
#REISSUE_RAW_TX=$(e1-cli getrawtransaction $RITXID 1)
#REISSUE_VOUTS=$(echo $REISSUE_RAW_TX | jq -r '.vout')
#VOUT_ADDRESS_REISSUE=$(echo $REISSUE_VOUTS | jq -r '.[0].scriptPubKey.addresses[0]')
#e1-cli importaddress $VOUT_ADDRESS_REISSUE

e1-cli listissuances

UBRISSUE=$(e2-cli issueasset 55 1 false)

UBASSET=$(echo $UBRISSUE | jq -r '.asset')

e2-cli getwalletinfo

e2-cli generatetoaddress 1 $ADDRGEN2
sleep 10

e1-cli listissuances

UBRITXID=$(echo $UBRISSUE | jq -r '.txid')

UBRIADDR=$(e2-cli gettransaction $UBRITXID | jq -r '.details[0].address')
e1-cli importaddress $UBRIADDR

# Or if the address is not known to e1 but the TXID is (requires index=1 in config file to work):
#UBREISSUE_RAW_TX=$(e1-cli getrawtransaction $UBRITXID 1)
#UBREISSUE_VOUTS=$(echo $UBREISSUE_RAW_TX | jq -r '.vout')
#UBVOUT_ADDRESS_REISSUE=$(echo $UBREISSUE_VOUTS | jq -r '.[0].scriptPubKey.addresses[0]')
#e1-cli importaddress $UBVOUT_ADDRESS_REISSUE

e1-cli listissuances

e2-cli destroyamount $UBASSET 5
e2-cli getwalletinfo

### Block Signing ###

e1-cli generatetoaddress 1 $ADDRGEN1
sleep 10

ADDR1=$(e1-cli getnewaddress)
ADDR2=$(e2-cli getnewaddress)

VALID1=$(e1-cli getaddressinfo $ADDR1)
PUBKEY1=$(echo $VALID1 | jq -r '.pubkey')

VALID2=$(e2-cli getaddressinfo $ADDR2)
PUBKEY2=$(echo $VALID2 | jq -r '.pubkey')

KEY1=$(e1-cli dumpprivkey $ADDR1)
KEY2=$(e2-cli dumpprivkey $ADDR2)

MULTISIG=$(e1-cli createmultisig 2 '''["'''$PUBKEY1'''", "'''$PUBKEY2'''"]''')
REDEEMSCRIPT=$(echo $MULTISIG | jq -r '.redeemScript')
echo $REDEEMSCRIPT

e1-cli stop
e2-cli stop
sleep 15

SIGNBLOCKARGS=("-signblockscript=$(echo $REDEEMSCRIPT)" "-con_max_block_sig_size=214" "-con_dyna_deploy_start=0")

rm -r ~/elementsdir1/elementsregtest
rm -r ~/elementsdir2/elementsregtest

e1-dae ${SIGNBLOCKARGS[@]}
e2-dae ${SIGNBLOCKARGS[@]}

sleep 10

# Create new wallets
e1-cli createwallet ""
e2-cli createwallet ""
e1-cli rescanblockchain
e2-cli rescanblockchain

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

SIGN1=$(e1-cli signblock $HEX "$REDEEMSCRIPT")
SIGN2=$(e2-cli signblock $HEX "$REDEEMSCRIPT")

SIGN1DATA=$(echo $SIGN1 | jq '.[0]')
SIGN2DATA=$(echo $SIGN2 | jq '.[0]')

COMBINED=$(e1-cli combineblocksigs $HEX "[$SIGN1DATA,$SIGN2DATA]" "$REDEEMSCRIPT")

SIGNEDBLOCK=$(echo $COMBINED | jq -r '.hex')

e2-cli submitblock $SIGNEDBLOCK

e1-cli getblockcount
e2-cli getblockcount

e1-cli stop
e2-cli stop
sleep 15

### Sidechain - Peg-In ###

rm -r ~/elementsdir1/elementsregtest
rm -r ~/elementsdir2/elementsregtest

# When testing, you can also use the OP_TRUE script -fedpegscript=51 
# so that you do not have to provide any pubkey values as we do below.
FEDPEGARG="-fedpegscript=5221$(echo $PUBKEY1)21$(echo $PUBKEY2)52ae"

e1-dae $FEDPEGARG
e2-dae $FEDPEGARG
sleep 10

# Create new wallets
e1-cli createwallet ""
e2-cli createwallet ""
e1-cli rescanblockchain
e2-cli rescanblockchain

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

MAINCHAIN=$(echo $ADDRS | jq -r '.mainchain_address')
CLAIMSCRIPT=$(echo $ADDRS | jq -r '.claim_script')

b-cli getwalletinfo

TXID=$(b-cli sendtoaddress $MAINCHAIN 1)

b-cli getwalletinfo

b-cli generatetoaddress 101 $ADDRGENB

b-cli getwalletinfo

PROOF=$(b-cli gettxoutproof '''["'''$TXID'''"]''')
RAW=$(b-cli getrawtransaction $TXID)

CLAIMTXID=$(e1-cli claimpegin $RAW $PROOF $CLAIMSCRIPT)

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

rm -r ~/elementsdir1/elementsregtest
rm -r ~/elementsdir2/elementsregtest

STANDALONEARGS=("-validatepegin=0" "-defaultpeggedassetname=newasset" "-initialfreecoins=100000000000000" "-initialreissuancetokens=200000000")

e1-dae ${STANDALONEARGS[@]}
e2-dae ${STANDALONEARGS[@]}
sleep 10

# Create new wallets
e1-cli createwallet ""
e2-cli createwallet ""
e1-cli rescanblockchain
e2-cli rescanblockchain

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

DEFAULTRIT=$(e1-cli getwalletinfo | jq '[.balance] | .[0]' | jq -r 'keys | .[0]')

if [ $DEFAULTRIT = "newasset" ]; then
  DEFAULTRIT=$(e1-cli getwalletinfo | jq '[.balance] | .[0]' | jq -r 'keys | .[1]')
fi

echo $DEFAULTRIT

e1-cli sendtoaddress $(e1-cli getnewaddress) 1000000 "" "" true

e1-cli sendtoaddress $(e1-cli getnewaddress) 2 "" "" false false 1 UNSET false $DEFAULTRIT
e1-cli generatetoaddress 101 $ADDRGEN1
sleep 10

e1-cli sendtoaddress $(e2-cli getnewaddress) 500 "" "" false 

e1-cli sendtoaddress $(e2-cli getnewaddress) 1 "" "" false false 1 UNSET false $DEFAULTRIT
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


