---
layout: page
title: How it works
permalink: /how-it-works
---

# How Elements works and the roles of network participants.

## Elements provides a technical solution to problems blockchain users face daily; transaction latency, lack of privacy, and risk to fungibility.

Elements overcomes these problems through its use of Federated Block Signing and Confidential Transactions.

Unlike the Bitcoin network, the process of block signing within Elements is not reliant on Dynamic Membership Multiparty Signatures (DMMS) and Proof of Work (PoW). Instead, Elements uses a **Strong Federation** of signatories, called **Block Signers**, who sign and create blocks every minute or so. This removes the transaction latency of the PoW mining process, which is subject to large block time variance due to its random poisson distrubution. The Federated Block Signing process achieves reliable block creation without introducing the need for third party trust.
 
The Strong Federation also contains members who enable the secure and controlled transfer of assets between a main chain and an Elements sidechain. Members who perform this role are called **Watchmen**. Next, we will look at the different roles played by members of the Strong Federation.

* * * 
 
### Strong Federations
 
Elements uses a consensus model proposed by Blockstream, called [Strong Federations](https://blockstream.com/strong-federations.pdf). A Strong Federation does not need Proof of Work and instead relies on the collective actions of a group of mutually-distrusting participants, called Functionaries.
 
The role of a Functionary is to propose, sign and verify the validity of actions on the network. Once a threshold of signatories have signed their acceptance of an action, consensus is said to have been reached and the action is given finality within the network.
 
The two roles a Functionary can fulfill within a Strong Federation are...
 
* **Watchmen** - participate in moving assets in and out of a Sidechain by signing multi-signature transactions.
 
* **Block Signers** - participate in creating blocks by adding their signature to count towards a threshold needed to validate proposed blocks, thereby defining the consensus history of transactions.

These actions are split between two distinct roles in order to enhance security and limit the damage an attacker can cause.

When combined, the roles of these participants allows Elements to deliver both rapid block creation (faster and final transaction confirmation) and assured, transferable assets (pegged assets directly linkable to another blockchain).
 
We will look at how Block Signers create blocks later and will begin by finding out how Watchmen enforce something called a Federated Peg, which allows the 1-to-1 transfer of assets between an Elements sidechain and another blockchain, typically Bitcoin.

* * * 
<a id="federatedpeg"></a>
### The role of Watchmen in a Strong Federation

In order for a sidechain to operate in a trustworthy manner it must allow participants to verify that the supply of assets is controlled and verifiable. An Elements sidechain uses a **Federated Peg** to enable the two way transfer of assets in and out of an Elements blockchain. This satisfies the requirements of provable issuance and inter-chain transfers.
 
At a high level, transfers into the sidechain occur when someone sends main chain assets to an address controlled by a multi-signature Watchmen wallet. This effectively freezes the assets on the main chain. Watchmen then validate the transaction and releases the same amount of the associated asset within the sidechain. The released assets are sent to a sidechain wallet that can prove claim to the original main chain assets. This process effectively moves assets from the parent chain to the sidechain. 


**Diagram showing how Watchmen enable the transfer of assets into an Elements sidechain:**

![watchmen]({{ site.url }}/images/peg-in.png)

<br/>

In order to transfer assets back to the main chain, a user makes a special peg-out transaction on the sidechain. This transaction is checked by Watchmen who then sign a transaction spending from the multi-signature wallet they control on the main chain. A threshold number of participants in the federation must sign before the main chain transaction becomes valid. When the Watchmen send an asset back to the main chain they also destroy the corresponding amount on the sidechain, effectively transferring the assets between blockchains.

**Diagram showing how Watchmen enable the transfer of assets out of an Elements sidechain:**

![watchmen]({{ site.url }}/images/peg-out.png)

<br/>

The Watchmen observe both the main blockchain and the Elements sidechain in order to validate asset transfers between them. A set of geographically and jurisdictionally distributed servers are preferred, creating a compromise-resistant network of Functionaries.
 
This network retains a number of the beneficial properties of a fully decentralized security model without introducing the need for a trusted 3rd party or single point of failure.

* * * 
<a id="signedblocks"></a>
### The role of Block Signers in a Strong Federation

We have already mentioned how a federation of Watchmen control the transfer of assets between blockchains and we will now look at how Block Signers perform their role within the Strong Federation.
 
A blockchain like Bitcoin’s is extended when anyone forming part of a dynamic group of block signers extends the chain by demonstrating proof of work expended. The dynamic nature of the set introduces the latency issues inherent to such systems.
 
By using a fixed signer set a Federated model replaces the dynamic set with a known set, multi-signature scheme. Reducing the number of participants needed to extend the blockchain increases the speed and scalability of the system, while validation by all parties ensures integrity of the transaction history.
 
Federated block signing consists of several phases:
 
**Step 1** - Block Signers propose candidate blocks in a round-robin fashion to all other participating Block Signers.
 
**Step 2** - Each Block Signer signals their intent by pre-committing to sign the given candidate block.
 
**Step 3** - If the given threshold for pre-commitment is met, each Block Signer signs the block.
 
**Step 4** - If the signature threshold (which may be different from that of step 3) is met, the block is accepted and sent to the network. The Strong Federation has reached consensus on the latest block of transactions.
 
**Step 5** - The next block is then proposed by the next Block Signer in the round-robin and the process repeats.

Because a Strong Federation’s block generation is not probabilistic and is based on a fixed set of signers, it will never be subject to multi-block reorganizations. This allows for a significant reduction in the wait time associated with confirming transactions. It also removes the incentive to mine for profit (i.e. the block rewards) and replaces it with an incentive to productively participate in a network where all participants have the same shared goal; ensuring the network continues to function in a manner that is beneficial to all. It does this without introducing a single point of failure or higher trust requirements.
 
* * * 
 
### Asset Issuance, Asset Reissuance and destroying Assets

Initial asset issuance, the re-issuance of additional amounts of existing assets and the destroying of assets are controlled by reissuance tokens. These tokens act as a verifiable right to adjust the circulating amounts of an asset and are exchangeable and verifiable amongst participants in the network. 
 
Elements allows for use-cases such as token issuance, digitizable collectables, reward points, and attested assets (for example gold coins) to be realized on a blockchain. With Elements, you can issue and transact as many different asset types as you like. 
 
Every asset type optionally benefits from features such as Confidential Transactions, which provides privacy over the amount and type of asset being transferred. This allows different assets to be given different privacy properties depending on the requirements of the asset use-case.
 
The **Federated 2-way Peg** feature allows such assets to be interoperable with other blockchains and representative of another blockchain’s native asset. By pegging your blockchain to another, you can extend the capabilities of the main chain and overcome some of its inherent limitations.

#### Whether your Elements project is set up to operate as a standalone blockchain or as a sidechain, Strong Federations technology provides compelling features, while retaining the properties of a trust-minimized system.

[Next: Running Elements as a Blockchain or Sidechain]({{ site.url }}/blockchain-or-sidechain)
