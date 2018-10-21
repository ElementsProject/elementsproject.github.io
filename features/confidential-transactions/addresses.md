---
layout: page
title: Confidential Transactions - Addresses
permalink: features/confidential-transactions/addresses
redirect_from: /elements/confidential-transactions/addresses
---

# Confidential Transactions - Confidential Addresses

## The most visible property of Confidential Transactions is the introduction of a new default address type, the Confidential Address

An example of a Confidential Address is shown below:

~~~~
CTEwQjyErENrxo8dSQ6pq5atss7Ym9S7P6GGK4PiGAgQRgoh1iPUkLQ168Kqptfnwmpxr2Bf7ipQsagi
~~~~

The most obvious differences are that it starts with ``CT`` and is longer than usual. This is due to the inclusion of a public blinding key prepended to the base address. In the Elements Wallet, the blinding key is derived by using the wallet's master blinding key and unblinded P2PKH address. Therefore the receiver alone can decrypt the sent amount, and can hand it to auditors if needed. On the sender's side, ``sendtoaddress`` will use this pubkey to transmit the necessary info to the receiver, encrypted, and inside the transaction itself. The sender's wallet will also record the amount and hold this in the ``wallet.dat`` metadata as well as the ``audit.log`` file.

You can use the validateaddress command to show details of a Confidential Address:

~~~~
elements-cli validateaddress 
~~~~

The output of which will look something like: 

<div class="console-output">
CTEwQjyErENrxo8dSQ6pq5atss7Ym9S7P6GGK4PiGAgQRgoh1iPUkLQ168Kqptfnwmpxr2Bf7ipQsagi
{
  "isvalid": true,
  "address": "CTEwQjyErENrxo8dSQ6pq5atss7Ym9S7P6GGK4PiGAgQRgoh1iPUkLQ168Kqptfnwmpxr2Bf7ipQsagi",
  "scriptPubKey": "76a91448a67bbdaf57b6f55b50f02fcaacfa079900853588ac",
  "confidential_key": "02483237addc73befdb9b851f948c1488cbb7cf1a59ba8af36be1c479e0f6e8bc7",
  "unconfidential": "QTE8CaT6FcJEqkCR4ZGdoUJfas57eDqY6q",
  "ismine": true,
  "iswatchonly": false,
  "isscript": false,
  "pubkey": "0347b013d415f7dc964cfadd0bb0627c48ae6ae27a58cdd37d71990eaf8f38c60c",
  "iscompressed": true,
  "account": ""
}
</div>

As you can see the unconfidential P2PKH address starts with a `Q`. P2SH start with `H`. Most RPC calls outside of ``getnewaddress`` (e.g. ``listunspent``) will return the unblinded version of addresses. If you provide an unconfidential address to ``validateaddress`` it will show the confidential address.

You **must** use the confidential address in ``sendtoaddress``, ``sendfrom``, ``sendmany`` and ``createrawtransaction`` if you want to create confidential transactions. Therefore, when you want to receive confidential transactions you must give the *confidential* address to the sender. For all other RPC's except ``dumpblindingkey`` it does not matter whether the confidential or unconfidential address is provided.

### Find out more about Confidential Transactions

Gregory Maxwell's [original investigation]({{ site.url }}/features/confidential-transactions/investigation) into Confidential Transactions.

Try it yourself in the [Elements Code Tutorial]({{ site.url }}/elements-code-tutorial/confidential-transactions).


