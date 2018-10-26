---
layout: page
title: Issued Assets
permalink: features/issued-assets
redirect_from:
  - /elements/asset-issuance/
  - /elements/asset-issuance
---

#  Issued Assets

## You can issue your own Confidential Assets on Elements

Assets issued on Elements can represent fungible ownership of any asset including vouchers, coupons, currencies, deposits, bonds, shares, etc.

By default, these assets are covered by the Confidential Transactions feature, which hides both the amount and type of asset being transacted. The transacting parties may choose to reveal the details for a transaction to a third party by sharing the ["blinding key"]({{ site.url }}/elements-code-tutorial/confidential-transactions#blindingkey), which grants visibility into the transaction, otherwise it remains private between them.

Issued Assets on Elements opens the door for building trustless exchanges, options, and other advanced smart contracts involving self-issued assets.

You can read the [original investigation]({{ site.url }}/features/issued-assets/investigation) by Andrew Poelstra for the cryptographic breakdown of how Confidential Assets work.

### Issuing, reissuing and destroying assets

In Elements you can not only issue your own asset but also reissue and destroy amounts of it, using something called a reissuance token.

You can also change the name of the default asset that is created upon chain initialization and the default asset used to pay fees on the network.

Details of how to [issue your own assets]({{ site.url }}/elements-code-tutorial/issuing-assets), [reissue and destroy assets]({{ site.url }}/elements-code-tutorial/reissuing-assets) and change the [default and fee asset]({{ site.url }}/elements-code-tutorial/blockchain) are provided in the code tutorial.

### How it works

All outputs are tagged with an asset commitment. Like Confidential Transactions, the consensus rules are such that instead of checking that amounts are balanced, the value commitments are checked for balance. A new transaction type is added for creating issued assets (asset definition transactions). Asset definition transactions have explicit `assetIssuance` fields embedded within transaction inputs, up to one issuance per input, which denote the issuance of both the asset itself and the reissuance tokens if desired. Embedding issuances within an input allows us to re-use the entropy as a NUMS for the asset type itself. 

Explicit and consensus-enforced asset issuance has many advantages over other approaches, such as colored coins:

* Supports SPV wallets much more efficiently.
* Allows more complex consensus-enforced contracts.
* Benefits from other consensus-enforced extensions (i.e. Confidential Transactions would not be compatible with colored coins).
* Opens the door for further consensus-enforced extensions that rely on the chain being able to operate with multiple assets.

### Find out more about Issued Assets

Andrew Poelstra's [original investigation]({{ site.url }}/features/issued-assets/investigation).

Try it yourself in the [Elements Code Tutorial]({{ site.url }}/elements-code-tutorial/issuing-assets).
