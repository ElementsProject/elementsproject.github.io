---
layout: page
title: Elements Features
permalink: /features
---

# The technical features of Elements

## Elements has many useful features that are enabled by default

<a id="confidentialtransactions"></a>
### Confidential Transactions

By default, all addresses in Elements are blinded using Confidential Transactions. Blinding is the process by which the amount and type of asset being transferred is cryptographically hidden from everyone, except the participants and those they choose to reveal the blinding key to.

Find our more about [Confidential Transactions]({{ site.url }}/features/confidential-transactions).

Try it yourself in the [Elements Code Tutorial]({{ site.url }}/elements-code-tutorial/confidential-transactions).

<a id="issuedassets"></a>
### Issued Assets
Issued Assets on Elements allows multiple types of asset to be issued and transferred between network participants. An Issued Asset also benefits from Confidential Transactions and can be reissued or destroyed by anyone holding the relevant reissuance token.

Find out more about [Issued Assets]({{ site.url }}/features/issued-assets).

Try it yourself in the [Elements Code Tutorial]({{ site.url }}/elements-code-tutorial/issuing-assets).

<a id="federatedpeg"></a>
### Federated Two-Way Peg

Elements is a general purpose blockchain platform that can also be “pegged” to an existing blockchain (such as Bitcoin) in order to enable the two way transfer of assets from one chain to the other. Implementing Elements as a sidechain allows you to work around some of the inherent properties of the main chain, while retaining a good degree of the security provided by assets secured on the main chain.

Find out more about [Federated Two-Way Pegs]({{ site.url }}/how-it-works#federatedpeg).

Try it yourself in the [Elements Code Tutorial]({{ site.url }}/elements-code-tutorial/sidechain).

<a id="signedblocks"></a>
### Signed Blocks
Elements uses a Strong Federation of signatories, called Block Signers, who sign and create blocks in a reliable and timely manner. This removes the transaction latency of the PoW mining process, which is subject to large block time variance due to its random poisson distrubution. The Federated Block Signing process achieves reliable block creation without introducing the need for third party trust.

Find out more about [Signed Blocks]({{ site.url }}/how-it-works#signedblocks).

Try it yourself in the [Elements Code Tutorial]({{ site.url }}/elements-code-tutorial/block-creation).

<a id="opcodes"></a>
### Additional opcodes
Elements introduces several new script opcodes, in addition to the ones already supported by Bitcoin.

Find out more about the [Additional opcodes]({{ site.url }}/features/opcodes).

