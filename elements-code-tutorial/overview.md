---
layout: page
title: Elements Tutorial
permalink: /elements-code-tutorial/overview
---

# A simple guide to setting up an Elements Blockchain or Sidechain solution.

## This guide will take you through the process of installing and running Elements step-by-step. 

### The code examples will show you how to:

* Configure two Elements instances that can send transactions between each other.

* Use some basic wallet functionality to send assets between instances, query balances etc.

* Understand how Confidential Transactions work and how to view amounts and asset types sent between two participants using "unblinding".

* Issue your own native blockchain assets.

* Send the new asset between network participants.

* Reissue more of the asset.

* Destroy an amount of the asset.

* Sign blocks using the Strong Federation ’n of m’ model of block creation.

### ...and if you choose to run Elements as a sidechain:

* Send assets from a main chain (Bitcoin in this example) to our Elements blockchain using the Federated Peg feature.

* Send assets from our Elements blockchain back to the Bitcoin main chain.

* * * 

### The tutorial is divided up into logical parts:

[Installing Bitcoin]({{ site.url }}/elements-code-tutorial/installing-bitcoin)

[Installing Elements]({{ site.url }}/elements-code-tutorial/installing-elements)

[Setting up your working environment]({{ site.url }}/elements-code-tutorial/working-environment)

[Using Elements to perform basic operations]({{ site.url }}/elements-code-tutorial/basic-operations)

[Using Confidential Transactions]({{ site.url }}/elements-code-tutorial/confidential-transactions)

[Issuing your own assets]({{ site.url }}/elements-code-tutorial/issuing-assets)

[Reissuing assets]({{ site.url }}/elements-code-tutorial/reissuing-assets)

[Block creation in a Strong Federation]({{ site.url }}/elements-code-tutorial/block-creation)

[Elements as a Sidechain]({{ site.url }}/elements-code-tutorial/sidechain)

[Elements as a standalone Blockchain]({{ site.url }}/elements-code-tutorial/blockchain)

[Developing applications on Elements]({{ site.url }}/elements-code-tutorial/application-development)

[Desktop application example in Python]({{ site.url }}/elements-code-tutorial/desktop-application-python)

[Web application example]({{ site.url }}/elements-code-tutorial/web-application)

[Desktop application example in C# using .NET Core]({{ site.url }}/elements-code-tutorial/desktop-application-dotnetcore)

[An easy way to run the main tutorial code]({{ site.url }}/elements-code-tutorial/easy-run-code)

* * * 

#### If you want to just run the code and not follow the tutorial you can skip to the [An easy way to run the main tutorial code]({{ site.url }}/elements-code-tutorial/easy-run-code) section, although this code is not annotated and steps are not explained.

The instructions have been tested against newly installed Ubuntu 16.04, 17.10.1 and 18.04.1 machines using the "Minimal Installation" Ubuntu install option. 

Please note that the `terminal commands` used within the tutorial may wrap over more than one line and that each line should be run in its entirety.

By following the guide through to completion you should have enough knowledge to build and deploy your own Elements based blockchain.

[Get started: Installing Bitcoin]({{ site.url }}/elements-code-tutorial/installing-bitcoin)
