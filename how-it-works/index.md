---
layout: page
title: How it works
permalink: /how-it-works
redirect_from: 
  - /elements/deterministic-pegs/
  - /elements/signed-blocks/
  - /elements/deterministic-pegs
  - /elements/signed-blocks
  - /blockchain-or-sidechain
---

# How Elements works and the roles of network participants

## Elements provides a technical solution to problems blockchain users face daily; transaction latency, lack of privacy, and risk to fungibility

Elements overcomes these problems through its use of Federated Block Signing and Confidential Transactions.

Unlike the Bitcoin network, the process of block signing within Elements is not reliant on Dynamic Membership Multiparty Signatures (DMMS) and Proof of Work (PoW). Instead, Elements uses a **Strong Federation** of signatories, called **Block Signers**, who sign and create blocks in a reliable and timely manner. This removes the transaction latency of the PoW mining process, which is subject to large block time variance due to its random poisson distribution. The Federated Block Signing process achieves reliable block creation without introducing the need for third party trust.
 
When Elements is being run as a sidechain, the Strong Federation will also contain members who enable the secure and controlled transfer of assets between a main chain and the Elements sidechain. Members who perform this role are called **Watchmen**. Next, we will look at the different roles played by members of the Strong Federation.

* * * 
 
### Strong Federations
 
Elements uses a consensus model proposed by Blockstream, called [Strong Federations](https://blockstream.com/strong-federations.pdf). A Strong Federation does not need Proof of Work and instead relies on the collective actions of a group of mutually-distrusting participants, called Functionaries.
  
The roles a Functionary can fulfill within a Strong Federation are...
 
* **Block Signers** - required in either setup, Block Signers participate in creating blocks by adding their signature to count towards a threshold needed to validate proposed blocks, thereby defining the consensus history of transactions.

* **Watchmen** - in a sidechain setup, Watchmen participate in moving assets in and out of the sidechain by signing multi-signature transactions.
 
These actions are split between two distinct roles in order to enhance security and limit the damage an attacker can cause. The Watchmen role is only required if Elements is running as a sidechain, the Block Signer role is required in either sidechain or standalone blockchain setups.

When combined, the roles of these participants allows Elements to deliver both rapid block creation (faster and final transaction confirmation) and assured, transferable assets (pegged assets directly linkable to another blockchain).
 
We'll begin by seeing how Watchmen enforce something called a Federated Peg, which allows the 1-to-1 transfer of assets between an Elements sidechain and another blockchain, typically Bitcoin.

* * * 
<a id="federatedpeg"></a>
### The role of Watchmen in a Strong Federation

In order for a sidechain to operate in a trustworthy manner it must allow participants to verify that the supply of assets is controlled and verifiable. An Elements sidechain uses a **2-Way Federated Peg** to enable the two way transfer of assets in and out of an Elements blockchain. This satisfies the requirements of provable issuance and inter-chain transfers. The Federated 2-way Peg feature allows an asset to be interoperable with other blockchains and representative of another blockchain’s native asset. By pegging your blockchain to another, you can extend the capabilities of the main chain and overcome some of its inherent limitations.
 
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

When run in either sidechain or standalone blockchain mode, Elements allows for the issuance of new asset types. The issuance of new asset types is open to all network nodes. Anyone can destroy an asset if they hold at least the amount they are destroying in their wallet.

The re-issuance of additional amounts of existing assets are controlled by reissuance tokens. These tokens act as a verifiable right to increase the circulating amount of an asset and are exchangeable and verifiable amongst participants in the network. Reissuance tokens can only be created with the initial issuance of an asset.
 
Elements allows for use-cases such as token issuance, digitizable collectables, reward points, and attested assets (for example gold coins) to be realized on a blockchain. With Elements, you can issue and transact as many different asset types as you like. 
 
Every asset type optionally benefits from features such as Confidential Transactions, which provides privacy over the amount and type of asset being transferred. This allows different assets to be given different privacy properties depending on the requirements of the asset use-case.

#### Whether your Elements project is set up to operate as a standalone blockchain or as a sidechain, Strong Federations technology provides compelling features, while retaining the properties of a trust-minimized system.

* * * 

<a id="blockchain-or-sidechain"></a>
# The correct approach to deploying a solution built on Elements

## Elements can operate as a general purpose, standalone Blockchain or as a pegged Sidechain

A general purpose blockchain built on Elements acts as a standalone blockchain which has no dependency on another blockchain for its asset issuance.

In the context of the Elements platform, a Sidechain is an extension to an existing blockchain. Assets are transferable between chains allowing the main chain to benefit from the enhanced features of the sidechain, such as rapid transfer finality and confidential transactions. While a sidechain is aware of the main chain and its transaction history, the main chain has no awareness of the sidechain, and none is required for its operation. This enables sidechains to innovate without restriction or the delays associated with main chain protocol improvement proposals. Indeed, rather than trying to alter it directly, extending the main protocol with a sidechain allows the main chain itself to remain secure and specialized, underpinning the smooth operation of the sidechain.
 
An example of an Elements based sidechain in production use is Blockstream's Liquid; a sidechain with different features, capabilities, and benefits than the main Bitcoin blockchain. By extending the functionality of Bitcoin and leveraging its underlying security, the Liquid network is able to provide new functionality that was previously not available to Bitcoin users. Liquid was designed to address the needs of exchanges, brokers and OTC trading desks. It enables the rapid, confidential and secure transfer of funds between participants and provides a solution to the inherent problem of delayed transaction finality on the Bitcoin network. Every Liquid bitcoin held within the sidechain is pegged to bitcoin on the main chain using a Federated 2-way Peg. This allows bitcoin to be deposited in the sidechain whilst retaining the ability to withdraw assets back to the main Bitcoin blockchain.
 
Elements provides the tools needed to prototype and deploy your own blockchains and sidechains in a cost effective and timely manner.

### When to use a Blockchain

A blockchain is an inherently inefficient means of data storage when compared to centralized database solutions. This inefficiency is a price worth paying only if you genuinly require the properties that a blockchain offers. It is important to correctly evaluate the requirements of your proposed system before deciding on a blockchain based solution.

##### Note: If your requirements do not necessitate it, implementing a blockchain based system will be counterproductive.
 
When should you use a blockchain as a solution?

* No central or single party should have sole authority to write data entries.

* A group of participants need to be able to access and audit data and be sure that it is cryptographically protected against tampering.

* Speed of transaction is not a key requirement. Note that Elements' Federated Block Signing process is still much faster than Dynamic-Membership Multi-party Signature (DMMS) Proof of Work (PoW) as found in Bitcoin though.

* Immutability (the inability to adjust confirmed transactions) is required.

* Security and redundancy with no single point of failure is required.

* Participants have the same incentives to participate.

##### Deciding if a blockchain is an appropriate way to deliver your project is a very important process as Blockchains only provide solutions to very specific types of problem.

### In the Code Tutorial

Learn how to run Elements as a [standalone blockchain]({{ site.url }}/elements-code-tutorial/blockchain)

Learn how to run Elements as a [sidechain]({{ site.url }}/elements-code-tutorial/sidechain)
