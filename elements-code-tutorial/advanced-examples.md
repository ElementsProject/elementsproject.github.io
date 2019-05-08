---
layout: page
title: Elements advanced examples
permalink: /elements-code-tutorial/advanced-examples
---

# Elements code tutorial

## Advanced examples

The code below shows you how to carry out the following actions within Elements:

* Creating a transaction manually (createrawtransaction) using an Issued Asset.

* Creating an asset issuance manually (rawissueasset).

* Proving that you issued an asset using 'Contract Hash'.

Save the code below in a file named **advancedexamples.sh** and place it in your home directory.

To run this code just open a terminal in your $HOME directory and run:

~~~~
bash advancedexamples.sh
~~~~

You can run each of the examples individualy by passing in the following command line arguments:

'RTIA' for raw transaction using an issued asset

'RIA' for raw issuance of an asset

'POI' for proof of issuance / contract hash

For example, to run the raw issuance example only:
~~~~
bash advancedexamples.sh RIA
~~~~
 
If you do not pass an argument in, all examples will run.

The examples work with v0.17 of Elements.

* * *

##### Note: If you want to run some of the steps automatically and then have execution stop and wait for you to press enter before continuing one line at a time: move the **trap read debug** statement down so that it is above the line you want to stop at. Execution will run each line automatically and stop when that line is reached. It will then switch to executing one line at a time, waiting for you to press return before executing the next command. Remove it to run without stopping.<br/><br/>You will see that occasionally we will use the **sleep** command to pause execution. This allows the daemons time to do things like stop, start and sync mempools.<br/><br/>There is a chance that the " and ' characters may not paste into your **runtutorial.sh** file correctly, so type over them yourself if you come across any issues executing the code.
 
~~~~
#!/bin/bash
set -x
trap read debug

#
# Save this code in a file named advancedexamples.sh and place it in your home directory.
#
# To run this code just open a terminal in your home directory and run:
# bash advancedexamples.sh
#
# You can run each of the examples individualy by passing in the following command line arguments:
# 'RTIA' for raw transaction using an issued asset
# 'RIA' for raw issuance of an asset
# 'POI' for proof of issuance / contract hash
# For example, to run the raw issuance example only:
# bash advancedexamples.sh RIA
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

