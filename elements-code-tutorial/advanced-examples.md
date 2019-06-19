---
layout: page
title: Elements advanced examples
permalink: /elements-code-tutorial/advanced-examples
---

# Elements code tutorial

## Advanced examples

[Example 1](#raw): Manually creating a transaction, manually issuing an asset and using an asset contract.

[Example 2](#multi): Issuing an asset to a multi-sig and spending from the multi-sig.

* * *

<a id="raw"></a>
### Example 1: Manually creating a transaction, manually issuing an asset and using an asset contract

The code below shows you how to carry out the following actions within Elements:

* Creating a transaction manually (createrawtransaction) using an Issued Asset.

* Creating an asset issuance manually (rawissueasset).

* Proving that you issued an asset using 'Contract Hash'.

Save the code below in a file named **advancedexamplesraw.sh** and place it in your home directory.

To run the code, open a terminal in your $HOME directory and run:

~~~~
bash advancedexamplesraw.sh
~~~~

You can run each of the examples individualy by passing in the following command line arguments:

'RTIA' for raw transaction using an issued asset

'RIA' for raw issuance of an asset

'POI' for proof of issuance / contract hash

For example, to run the raw issuance example only:
~~~~
bash advancedexamplesraw.sh RIA
~~~~
 
If you do not pass an argument in, all examples will run.

The examples work with v0.17 of Elements.

##### Note: If you want to run some of the steps automatically and then have execution stop and wait for you to press enter before continuing one line at a time: move the **trap read debug** statement down so that it is above the line you want to stop at. Execution will run each line automatically and stop when that line is reached. It will then switch to executing one line at a time, waiting for you to press return before executing the next command. Remove it to run without stopping.<br/><br/>You will see that occasionally we will use the **sleep** command to pause execution. This allows the daemons time to do things like stop, start and sync mempools.<br/><br/>There is a chance that the " and ' characters may not paste into your **runtutorial.sh** file correctly, so type over them yourself if you come across any issues executing the code.
 
~~~~
#!/bin/bash
set -x
trap read debug

#
# Save this code in a file named advancedexamplesraw.sh and place it in your home directory.
#
# To run this code just open a terminal in your home directory and run:
# bash advancedexamplesraw.sh
#
# You can run each of the examples individualy by passing in the following command line arguments:
# 'RTIA' for raw transaction using an issued asset
# 'RIA' for raw issuance of an asset
# 'POI' for proof of issuance / contract hash
# For example, to run the raw issuance example only:
# bash advancedexamplesraw.sh RIA
# 
# If you do not pass an argument in, all examples will run.
#
# Press the return key to execute each line in turn.
#
# If you want to run some of the steps automatically and then have execution stop
# and wait for you to press enter before continuing one line at a time: move 
# the 'trap read debug' statement down so that it is above the line you want to 
# stop at. Execution will run each line automatically and stop when that line is 
# reached. It will then switch to executing one line at a time, waiting for you 
# to press return before executing the next command, or just remove it completely
# to run the code without interuption.

if [ "$1" != "" ]; then
    EXAMPLETYPE=$1
    echo "Running '$EXAMPLETYPE' example"
else
    EXAMPLETYPE="ALL"
    echo "Running all examples"
fi

##############################################
#                                            #
#  RESET CHAIN STATE AND SET UP ENVIRONMENT  #
#                                            #
##############################################

# RESET CHAIN STATE AND SET UP ENVIRONMENT >>>

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

# The following 3 lines may error without issue if the daemons are not already running
b-cli stop
e1-cli stop
e2-cli stop
sleep 5

cd

# The following 3 lines may error without issue
rm -r ~/bitcoindir ; rm -r ~/elementsdir1 ; rm -r ~/elementsdir2
mkdir ~/bitcoindir ; mkdir ~/elementsdir1 ; mkdir ~/elementsdir2

cd elements
cd src

cp ~/elements/contrib/assets_tutorial/bitcoin.conf ~/bitcoindir/bitcoin.conf
cp ~/elements/contrib/assets_tutorial/elements1.conf ~/elementsdir1/elements.conf
cp ~/elements/contrib/assets_tutorial/elements2.conf ~/elementsdir2/elements.conf

b-dae

sleep 5

e1-dae
e2-dae

sleep 5

e1-cli getwalletinfo
e2-cli getwalletinfo

e1-cli sendtoaddress $(e1-cli getnewaddress) 21000000 "" "" true
e1-cli generate 101
sleep 5

# Turn on 'stop on error':
set -e

# <<< RESET CHAIN STATE AND SET UP ENVIRONMENT


###########################################
#                                         #
#  RAW TRANSACTION USING AN ISSUED ASSET  #
#                                         #
###########################################

# RAW TRANSACTION USING AN ISSUED ASSET >>>

if [ "RTIA" = $EXAMPLETYPE ] || [ "ALL" = $EXAMPLETYPE ] ; then

    # Issue an asset and get the asset hex so we can send some
    ISSUE=$(e1-cli issueasset 100 1)

    ASSET=$(echo $ISSUE | jq '.asset' | tr -d '"')

    # Check the asset shows up in our wallet
    e1-cli getwalletinfo

    # Get a list of unspent we can use as inputs - referenced via transaction id, vout and amount
    UTXO=$(e1-cli listunspent 0 0 [] true '''{"''asset''":"'''$ASSET'''"}''')

    TXID=$(echo $UTXO | jq '.[0].txid' | tr -d '"' )
    VOUT=$(echo $UTXO | jq '.[0].vout' | tr -d '"' )
    AMOUNT=$(echo $UTXO | jq '.[0].amount' | tr -d '"' )

    # Get an address to send the asset to - we'll use unconfidential
    ADDR=$(e1-cli getnewaddress)

    VALIDATEADDR=$(e1-cli validateaddress $ADDR)

    UNCONADDR=$(echo $VALIDATEADDR | jq '.unconfidential' | tr -d '"')

    # Build the raw transaction (send 3 of the asset)
    SENDAMOUNT="3.00"

    RAWTX=$(e1-cli createrawtransaction '''[{"''txid''":"'''$TXID'''", "''vout''":'$VOUT', "''asset''":"'''$ASSET'''"}]''' '''{"'''$UNCONADDR'''":'$SENDAMOUNT'}''' 0 false '''{"'''$UNCONADDR'''":"'''$ASSET'''"}''')

    # Fund the tx
    FRT=$(e1-cli fundrawtransaction $RAWTX)

    # blind and sign the tx
    HEXFRT=$(echo $FRT | jq '.hex' | tr -d '"')

    BRT=$(e1-cli blindrawtransaction $HEXFRT)

    SRT=$(e1-cli signrawtransactionwithwallet $BRT)

    HEXSRT=$(echo $SRT | jq '.hex' | tr -d '"')

    # Send the raw tx and confirm
    TX=$(e1-cli sendrawtransaction $HEXSRT)

    e1-cli generate 101

    # Decode the raw transaction so we can see the amount of the asset sent and the address it went to
    GRT=$(e1-cli getrawtransaction $TX)

    DRT=$(e1-cli decoderawtransaction $GRT)

    echo "An amount of" $SENDAMOUNT "of the '"$ASSET"' asset was sent to address '"$UNCONADDR"'"

    echo "END OF 'RTIA' EXAMPLE"
    
fi
# <<< RAW TRANSACTION USING AN ISSUED ASSET


###########################################
#                                         #
#             RAW ISSUE ASSET             #
#                                         #
###########################################

# RAW ISSUE ASSET >>>

if [ "RIA" = $EXAMPLETYPE ] || [ "ALL" = $EXAMPLETYPE ] ; then

    # Fund the transaction...
    ADDR=$(e1-cli getnewaddress)

    e1-cli sendtoaddress $ADDR 2
    e1-cli generate 1

    RAWTX=$(e1-cli createrawtransaction [] '''{"'''$ADDR'''":1.00}''')

    FRT=$(e1-cli fundrawtransaction $RAWTX)

    HEXFRT=$(echo $FRT | jq '.hex' | tr -d '"')

    # Get the unconfidential address we will create the asset at...
    VALIDATEADDR=$(e1-cli validateaddress $ADDR)

    UNCONADDR=$(echo $VALIDATEADDR | jq '.unconfidential' | tr -d '"')

    # Create the raw issuance
    RIA=$(e1-cli rawissueasset $HEXFRT '''[{"''asset_amount''":33, "''asset_address''":"'''$UNCONADDR'''", "''blind''":"''false''"}]''')

    # The results of which include... 
    HEXRIA=$(echo $RIA | jq '.[0].hex' | tr -d '"')
    ASSET==$(echo $RIA | jq '.[0].asset' | tr -d '"')
    ENTROPY==$(echo $RIA | jq '.[0].entropy' | tr -d '"')
    TOKEN==$(echo $RIA | jq '.[0].token' | tr -d '"')

    # Blind, sign and send the transaction that creates the asset issuance...
    BRT=$(e1-cli blindrawtransaction $HEXRIA)

    SRT=$(e1-cli signrawtransactionwithwallet $BRT)
    
    HEXSRT=$(echo $SRT | jq '.hex' | tr -d '"')

    ISSUETX=$(e1-cli sendrawtransaction $HEXSRT)

    e1-cli generate 101

    # Check that worked...
    e1-cli getwalletinfo

    e1-cli listissuances      
  
    echo "END OF 'RIA' EXAMPLE"
  
fi

# <<< RAW ISSUE ASSET


###########################################
#                                         #
#    PROOF OF ISSUANCE / CONTRACT HASH    #
#                                         #
###########################################

# PROOF OF ISSUANCE / CONTRACT HASH >>>

if [ "POI" = $EXAMPLETYPE ] || [ "ALL" = $EXAMPLETYPE ] ; then

    b-cli generate 101
    e1-cli generate 101

    # We need to get a 'legacy' type (prefix 'CTE') address for this:
    NEWADDR=$(e1-cli getnewaddress "" legacy)
    
    VALIDATEADDR=$(e1-cli getaddressinfo $NEWADDR)

    PUBKEY=$(echo $VALIDATEADDR | jq '.pubkey' | tr -d '"')

    ADDR=$NEWADDR
    
    # Write to an asset registry to prove we (ABC Company) control the address in question...
    MSG="THE ADDRESS THAT I, ABC, HAVE CONTROL OVER IS:[$ADDR] PUBKEY:[$PUBKEY]"

    SIGNEDMSG=$(e1-cli signmessage $ADDR "$MSG")

    VERIFIED=$(e1-cli verifymessage $ADDR $SIGNEDMSG "$MSG")

    # Write that proof to the registry...
    echo "REGISTRY ENTRY = MESSAGE:[$MSG] SIGNED MESSAGE:[$SIGNEDMSG]"

    FUNDINGADDR=$(e1-cli getnewaddress)

    e1-cli sendtoaddress $FUNDINGADDR 2

    e1-cli generate 1

    RAWTX=$(e1-cli createrawtransaction [] '''{"'''$FUNDINGADDR'''":1.00}''')

    FRT=$(e1-cli fundrawtransaction $RAWTX)

    HEXFRT=$(echo $FRT | jq '.hex' | tr -d '"')

    # Write whatever you want as a contract text. Include reference to signed proof of ownership above...
    CONTRACTTEXT="THIS IS THE CONTRACT TEXT FOR THE XYZ ASSET. CREATED BY ABC. ADDRESS:[$ADDR] PUBKEY:[$PUBKEY]"

    # Sign it using the same address it references...
    SIGNEDCONTRACTTEXT=$(e1-cli signmessage $ADDR "$CONTRACTTEXT")

    # Hash that signed message, which we will use as the contract hash...
    # (hash as 32 bytes however you want to do this - we will use openssl commandline)
    CONTRACTTEXTHASH=$(echo -n $SIGNEDCONTRACTTEXT | openssl dgst -sha256)
    CONTRACTTEXTHASH=$(echo ${CONTRACTTEXTHASH#"(stdin)= "})

    echo $CONTRACTTEXTHASH

    # Issue the asset to the address that we signed for earlier and which we included in the signed contract hash...
    RIA=$(e1-cli rawissueasset $HEXFRT '''[{"''asset_amount''":33, "''asset_address''":"'''$ADDR'''", "''blind''":"''false''", "''contract_hash''":"'''$CONTRACTTEXTHASH'''"}]''')

    # Details of the issuance...
    HEXRIA=$(echo $RIA | jq '.[0].hex' | tr -d '"')
    ASSET=$(echo $RIA | jq '.[0].asset' | tr -d '"')
    ENTROPY=$(echo $RIA | jq '.[0].entropy' | tr -d '"')
    TOKEN=$(echo $RIA | jq '.[0].token' | tr -d '"')

    # Blind, sign and send the issuance transaction...
    BRT=$(e1-cli blindrawtransaction $HEXRIA)

    SRT=$(e1-cli signrawtransactionwithwallet $BRT)
    
    HEXSRT=$(echo $SRT | jq '.hex' | tr -d '"')

    ISSUETX=$(e1-cli sendrawtransaction $HEXSRT)

    e1-cli generate 101

    # In the output from decoderawtransaction you will see in the vout section the asset being issued to the address we signed from earlier...
    RT=$(e1-cli getrawtransaction $ISSUETX)

    DRT=$(e1-cli decoderawtransaction $RT)

    # Build an asset registry entry saying that we issued the asset...
    ASSETREGISTERMESSAGE="I, ABC, CREATED ASSET:[$ASSET] WITH ASSET ENTROPY:[$ENTROPY] AT ADDRESS:[$ADDR] IN TX:[$ISSUETX]"

    SIGNEDMSG=$(e1-cli signmessage $ADDR "$ASSETREGISTERMESSAGE")

    e1-cli verifymessage $ADDR $SIGNEDMSG "$ASSETREGISTERMESSAGE"

    # Then make the entry in the aset registry...
    echo "REGISTRY ENTRY = ASSET CREATION MESSAGE:[$ASSETREGISTERMESSAGE] SIGNED VERSION:[$SIGNEDMSG]"

    e1-cli listissuances

    e1-cli getwalletinfo

    # Proving the issuance was indeed made against the contract hash...
    # We need to provide the following to anyone wishing to validate that the contract has was used to produce the asset:
    #  - Hex of funded raw transaction used to fund the issuance
    #  - contract_hash
    # Not needed (other values can be used): asset_amount, asset_address, blind

    # If someone else tries to claim they created the asset and we didn't - they will need to prove they can sign for the address it was sent to and explain how come we can sign messages (as found in the asset registry) for that address.

    VERIFYISSUANCE=$(e1-cli rawissueasset $HEXFRT '''[{"''asset_amount''":33, "''asset_address''":"'''$ADDR'''", "''blind''":"''false''", "''contract_hash''":"'''$CONTRACTTEXTHASH'''"}]''')

    ASSETVERIFY=$(echo $VERIFYISSUANCE | jq '.[0].asset' | tr -d '"')
    ENTROPYVERIFY=$(echo $VERIFYISSUANCE | jq '.[0].entropy' | tr -d '"')
    TOKENVERIFY=$(echo $VERIFYISSUANCE | jq '.[0].token' | tr -d '"')

    [[ $ASSET = $ASSETVERIFY ]] && echo ASSET HEX: VERIFIED || echo ASSET HEX: NOT VERIFIED
    [[ $ENTROPY = $ENTROPYVERIFY ]] && echo ENTROPY: VERIFIED || echo ENTROPY: NOT VERIFIED
    [[ $TOKEN = $TOKENVERIFY ]] && echo TOKEN: VERIFIED || echo TOKEN: NOT VERIFIED

    echo "END OF 'POI' EXAMPLE"
        
fi

# <<< PROOF OF ISSUANCE / CONTRACT HASH

# CLOSE RUNNING NODES BEFORE EXIT:

e1-cli stop
e2-cli stop
b-cli stop
~~~~

* * *

<a id="multi"></a>
### Example 2: Issuing an asset to a multi-sig and spending from the multi-sig

The code below shows you how to carry out the following actions within Elements:

* Create a multi-sig address and issue an asset to the multi-sig.

* Spend from the multi-sig.

Save the code below in a file named **advancedexamplesmulti.sh** and place it in your home directory.

To run the code, open a terminal in your $HOME directory and run:

~~~~
bash advancedexamplesmulti.sh
~~~~

##### Note: The script contains instructions telling you how to prepare the script before running.

~~~~
#!/bin/bash
set -x

# This script will create a multi-signature address shared between Wallet 1 and Wallet 2 and issue a new asset to the address.
# The asset can then only be spent and reissued by Wallet 1 and Wallet 2 signing the spending/reissuance transaction.

####################    BEFORE RUNNING THIS SCRIPT    ####################

# -----
#   1
# -----

# Before running this script, change the following variables to point to: your local Elements binaries directory, the directory you'll use to store node data.
# Note that the script will delete the $DATA_DIR/elementsregtest directory when run, so back up anything you might already have in there before running.
BINARY_DIR="$HOME/elements/src"
DATA_DIR="$HOME/elementsdir1"

# -----
#   2
# -----

# Create the directory referenced in DATA_DIR, if it is not already there.
# Within that directory, create a file named elements.conf, if it is not already there.
# Add/set the following as the contents of the elements.conf file (remove the # characters).

#daemon=1
#chain=elementsregtest
#elementsregtest.wallet=wallet.dat
#elementsregtest.wallet=wallet_1.dat
#elementsregtest.wallet=wallet_2.dat
#elementsregtest.wallet=wallet_3.dat
#validatepegin=0
#initialfreecoins=2100000000000000

# The settings will run the daemons in regtest mode, allocate some initial funds and create the required wallets.
# We will use the 4 wallets to:
#
# wallet (w-cli) 
# Generate blocks and allocate initial funds
#
# wallet_1 (w1-cli) and 
# wallet_2 (w2-cli) 
# Create and sign multi-signature asset spends (and associated tokens)
#
# wallet_3 (w3-cli) 
# Receive assets (single-signature receive)

# -----
#  Tip
# -----

# Uncomment and move the following line to any point in the script to stop execution and then continue execution line-by-line by pressing enter.
#trap read debug


####################    SCRIPT PREPARATION    ####################

# Create some aliases to make calling the node/wallet easier
shopt -s expand_aliases
# Node
alias n-dae="$BINARY_DIR/elementsd -datadir=$DATA_DIR"
# Client wallets
alias w-cli="$BINARY_DIR/elements-cli -datadir=$DATA_DIR -rpcwallet=wallet.dat"
alias w1-cli="$BINARY_DIR/elements-cli -datadir=$DATA_DIR -rpcwallet=wallet_1.dat"
alias w2-cli="$BINARY_DIR/elements-cli -datadir=$DATA_DIR -rpcwallet=wallet_2.dat"
alias w3-cli="$BINARY_DIR/elements-cli -datadir=$DATA_DIR -rpcwallet=wallet_3.dat"


# The following 'create_multisig' function will be called within the script to create the multisig addresses for the:
# - Asset Issuance
# - Asset Reissuance Token
# - Issuance change
# - Asset Reissuance
# - Tokens in Reissuance
# - Token Reissuance change.

create_multisig () {
    # Accepts one argument - a label for the multi-signature address to be created
    local LABEL=$1

    # Get an address from wallet 1 and wallet 2 which we will later use to create a multi-sig
    # Wallet 1's Address for the multi-sig:
    ADDRESS_1=$(w1-cli getnewaddress $LABEL)
    ADDRESS_1_INFO=$(w1-cli getaddressinfo $ADDRESS_1)
    ADDRESS_1_PUBKEY=$(echo $ADDRESS_1_INFO | jq '.pubkey' | tr -d '"')
    # We wil use the confidential key from Wallet 1's address to create a blinded address later:
    ADDRESS_1_CONF_KEY=$(echo $ADDRESS_1_INFO | jq '.confidential_key' | tr -d '"')
    # We will use the blinding key for Wallet 1's address and import it later so Wallet 1 and Wallet 2 can subsequently see unblinded details
    BLINDING_KEY=$(w1-cli dumpblindingkey $ADDRESS_1)

    # Wallet 2's Address for the multi-sig:
    ADDRESS_2=$(w2-cli getnewaddress $LABEL)
    ADDRESS_2_INFO=$(w2-cli getaddressinfo $ADDRESS_2)
    ADDRESS_2_PUBKEY=$(echo $ADDRESS_2_INFO | jq '.pubkey' | tr -d '"')

    # Create a multi-sig '2 of 2' address (n of m, where m is the number of public keys provided)
    N=2
    MULTISIG=$(w1-cli createmultisig $N '''["'''$ADDRESS_1_PUBKEY'''", "'''$ADDRESS_2_PUBKEY'''"]''')
    MULTISIG_ADDRESS=$(echo $MULTISIG | jq '.address' | tr -d '"')
    REDEEM_SCRIPT=$(echo $MULTISIG | jq '.redeemScript' | tr -d '"')
    
    # Blind the multi-sig address using the confidential key from Wallet 1 that we stored earlier
    BLINDED_ADDRESS=$(w1-cli createblindedaddress $MULTISIG_ADDRESS $ADDRESS_1_CONF_KEY)

    # Import the redeem script, blinded address, and the blinding key for the blinded address into both wallets
    w1-cli importaddress $REDEEM_SCRIPT '' false true
    w1-cli importaddress $BLINDED_ADDRESS "$LABEL multisig" false
    w1-cli importblindingkey $BLINDED_ADDRESS $BLINDING_KEY
    w2-cli importaddress $REDEEM_SCRIPT '' false true
    w2-cli importaddress $BLINDED_ADDRESS "$LABEL multisig" false
    w2-cli importblindingkey $BLINDED_ADDRESS $BLINDING_KEY
    
    # Return the blinded address and the public key to the function caller (will be extracted by splitting on '|' by the caller)
    echo "$BLINDED_ADDRESS|$ADDRESS_1_PUBKEY"
}


####################    ENVIRONMENT PREPERATION    ####################

# First clear the existing chain data and wallets 
# The following lines may error without issue if the node is not already running or the directory does not already exist
w-cli stop
echo "Wait for the node to stop if it was running..."
sleep 10
echo "Delete the data directory if it exists..."
rm -r $DATA_DIR/elementsregtest

# Set to stop the rest of the script if there is an error
set -e

# Start the daemon and wait for it to initialise
n-dae
sleep 10

# Move some funds to Wallet 1
W1_ADDR=$(w1-cli getnewaddress)
w-cli sendtoaddress $W1_ADDR 1000
w-cli generate 1
w1-cli getbalance "*" 0 true


####################    MULTI-SIG ISSUANCE    ####################

ASSET_AMOUNT="0.00000020"
REISSUANCE_TOKEN_AMOUNT="0.00000005"
FEERATE="0.00040000"

# Create multi-sig address for the asset
CREATE=$(create_multisig "asset")  
MULTISIG_ASSET_ADDRESS="$(echo $CREATE | cut -d'|' -f1)"
PUBKEY="$(echo $CREATE | cut -d'|' -f2)"

# Create multi-sig address for the reissuance token

CREATE=$(create_multisig "reissuance")  
MULTISIG_REISSUANCE_ADDRESS="$(echo $CREATE | cut -d'|' -f1)"

# Create the base transaction
BASE=$(w1-cli createrawtransaction '''[]''' '''{"''data''":"''00''"}''')

# Fund the transaction
FUNDED=$(w1-cli fundrawtransaction $BASE '''{"''feeRate''":'$FEERATE'}''')
FUNDED_HEX=$(echo $FUNDED | jq '.hex' | tr -d '"')

# Store vin txid and vout
DECODED=$(w1-cli decoderawtransaction $FUNDED_HEX)
PREV_TX=$(echo $DECODED | jq '.vin[0].txid' | tr -d '"')
PREV_VOUT=$(echo $DECODED | jq '.vin[0].vout' | tr -d '"')

# Create the contract for the asset. The hash of the contract will be used to generate the asset id.
CONTRACT_TEXT="Your contract text. Can be used to bind the asset to a real world asset etc."

# We will hash using openssl, other options are available
CONTRACT_TEXT_HASH=$(echo -n $CONTRACT_TEXT | openssl dgst -sha256)
CONTRACT_TEXT_HASH=$(echo ${CONTRACT_TEXT_HASH#"(stdin)= "})

# Create the raw issuance (will not yet be complete or broadcast)
RAW_ISSUE=$(w1-cli rawissueasset $FUNDED_HEX '''[{"''asset_amount''":'$ASSET_AMOUNT', "''asset_address''":"'''$MULTISIG_ASSET_ADDRESS'''", "''token_amount''":'$REISSUANCE_TOKEN_AMOUNT', "''token_address''":"'''$MULTISIG_REISSUANCE_ADDRESS'''", "''blind''":"''false''", "''contract_hash''":"'''$CONTRACT_TEXT_HASH'''"}]''')

# Store details of the issuance for later use
HEX_RIA=$(echo $RAW_ISSUE | jq '.[0].hex' | tr -d '"')
ASSET=$(echo $RAW_ISSUE | jq '.[0].asset' | tr -d '"')
TOKEN=$(echo $RAW_ISSUE | jq '.[0].token' | tr -d '"')
ENTROPY=$(echo $RAWISSUE | jq '.[0].entropy' | tr -d '"')

# Blind the issuance transaction
BLIND=$(w1-cli blindrawtransaction $HEX_RIA)

# Sign the issuance transaction (we only need sign with Wallet 1 - it is subsequent spends that will require multiple signatures)
SIGNED=$(w1-cli signrawtransactionwithwallet $BLIND)
HEX_SRT=$(echo $SIGNED | jq '.hex' | tr -d '"')
DECODED=$(w1-cli decoderawtransaction $HEX_SRT)

# Test the transaction's acceptance into the mempool
TEST=$(w1-cli testmempoolaccept '''["'$HEX_SRT'"]''')
ALLOWED=$(echo $TEST | jq '.[0].allowed' | tr -d '"')

# If the transaction is valid
if [ "true" = $ALLOWED ] ; then
    # Broadcast the transaction
    TXID=$(w1-cli sendrawtransaction $HEX_SRT)
    # Confirm the transaction
    w-cli generate 101
    # Check the issuance can be seen by Wallet 1
    w1-cli listissuances
    # Wallet 2 won't be able to see the actual amounts just yet
    w2-cli listissuances
    # Import the issuance blinding key into Wallet 2 so it can
    ISSUANCE_BLINDING_KEY=$(w1-cli dumpissuanceblindingkey $TXID 0)    
    w2-cli importissuanceblindingkey $TXID 0 $ISSUANCE_BLINDING_KEY
    # Now Wallet 2 will be able to see the amounts
    w2-cli listissuances
fi

# Check that both Wallet 1 and Wallet 2 see the asset and reissuance token in their balances
w1-cli getbalance "*" 0 true
w2-cli getbalance "*" 0 true


####################    MULTI-SIG SPEND    ####################

# Have Wallet 1 and Wallet 2 spend some of the asset held in the multi-sig by sending it to Wallet 3's single-signature address

AMOUNT="0.00000001"

# Get a receiving address from Wallet 3
RECEIVING_ADDRESS=$(w3-cli getnewaddress "receiving")

# Create a multi-sig address for Wallet 1 and Wallet 2 to receive asset change to
CREATE=$(create_multisig "change")  
MULTISIG_ASSET_CHANGE_ADDRESS="$(echo $CREATE | cut -d'|' -f1)"

# Get a change address for the bitcoin return from transaction fee spend
BITCOIN_CHANGE=$(w1-cli getrawchangeaddress)

# Create the multi-sig spending transaction
RAW_TX=$(w1-cli createrawtransaction '''[]''' '''{"'''$RECEIVING_ADDRESS'''":'$AMOUNT'}''' 0 false '''{"'''$RECEIVING_ADDRESS'''":"'''$ASSET'''"}''')

# Fund the transaction
FUNDED_RAW_TX=$(w1-cli fundrawtransaction $RAW_TX '''{"'''includeWatching'''":true, "'''changeAddress'''":{"'''bitcoin'''":"'''$BITCOIN_CHANGE'''", "'''$ASSET'''":"'''$MULTISIG_ASSET_CHANGE_ADDRESS'''"}}''')
FUNDED_HEX=$(echo $FUNDED_RAW_TX | jq '.hex' | tr -d '"')

# Blind the transaction
BLINDED_RAW_TX=$(w1-cli blindrawtransaction $FUNDED_HEX)

# Have Wallet 1 sign the transaction
SIGNED_RAW_TX=$(w1-cli signrawtransactionwithwallet $BLINDED_RAW_TX)
SIGNED_RAW_TX_HEX=$(echo $SIGNED_RAW_TX | jq '.hex' | tr -d '"')

# Have Wallet 2 sign the transaction
SIGNED_RAW_TX_2=$(w2-cli signrawtransactionwithwallet $SIGNED_RAW_TX_HEX)
SIGNED_RAW_TX_2_HEX=$(echo $SIGNED_RAW_TX_2 | jq '.hex' | tr -d '"')

# Test the transaction wil be accepted into the mempool
TEST=$(w1-cli testmempoolaccept '''["'$SIGNED_RAW_TX_2_HEX'"]''')
ALLOWED=$(echo $TEST | jq '.[0].allowed' | tr -d '"')

# If the transaction is valid
if [ "true" = $ALLOWED ] ; then
    # Broadcast the valid transaction
    TX=$(w1-cli sendrawtransaction $SIGNED_RAW_TX_2_HEX)
    # Confirm the transaction
    w-cli generate 1
fi

# Check that Wallet 3 received the asset
w3-cli getbalance

# And the balance of Wallet 1 and Wallet 2 have changed
w1-cli getbalance "*" 0 true
w2-cli getbalance "*" 0 true

# Stop the daemon
w-cli stop
sleep 10

~~~~
