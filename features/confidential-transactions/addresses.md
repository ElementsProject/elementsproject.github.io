---
layout: page
title: Confidential Transactions - Addresses
permalink: features/confidential-transactions/addresses
redirect_from:
  - /elements/confidential-transactions/addresses/
  - /elements/confidential-transactions/addresses
---

# Confidential Transactions - Confidential Addresses

## The most visible property of Confidential Transactions is the introduction of a new default address type, the Confidential Address

An example of a Confidential Address is shown below:

~~~~
CTEoaMTWgfUKoST1fkitVWuHHueqiL4EcUjrwJXuZW9RWKbSzYNq8ttd8KzJr3KBMzxh6HC83CgroNLR
~~~~

The most obvious differences are that it starts with ``CTE`` and is longer than usual. This is due to the inclusion of a public blinding key prepended to the base address. In the Elements Wallet, the blinding key is derived by using the wallet's master blinding key and unblinded P2PKH address. Therefore the receiver alone can decrypt the sent amount, and can hand it to auditors if needed. On the sender's side, ``sendtoaddress`` will use this pubkey to transmit the necessary info to the receiver, encrypted, and inside the transaction itself. The sender's wallet will also record the amount and hold this in the ``wallet.dat`` metadata as well as the ``audit.log`` file.

You can use the validateaddress command to show details of a Confidential Address:

~~~~
elements-cli validateaddress CTEoaMTWgfUKoST1fkitVWuHHueqiL4EcUjrwJXuZW9RWKbSzYNq8ttd8KzJr3KBMzxh6HC83CgroNLR
~~~~

The output of which will look something like: 

<div class="console-output">
CTEoaMTWgfUKoST1fkitVWuHHueqiL4EcUjrwJXuZW9RWKbSzYNq8ttd8KzJr3KBMzxh6HC83CgroNLR
{
  "isvalid": true,
  "address": "CTEoaMTWgfUKoST1fkitVWuHHueqiL4EcUjrwJXuZW9RWKbSzYNq8ttd8KzJr3KBMzxh6HC83CgroNLR",
  "scriptPubKey": "76a914f5b3992c73d1bee5844d840bbb30fe00f989c1fc88ac",
  "confidential_key": "028fafda60c741b4735c9b14fc768dbda05762f0355a09ceb8fcdd4da114c415bd",
  "unconfidential": "2dwpu6ECUS2NHMwwAmqCGJ6uNESr6tUmAS9",
  "ismine": true,
  "iswatchonly": false,
  "isscript": false,
  "pubkey": "02e232e5ec38bb4081915eff2bd00d81b1ef2168d095a9111f245e3e586c2b08ba",
  "iscompressed": true,
  "account": "",
  "timestamp": 1544461265,
  "hdkeypath": "m/0'/0'/2'",
  "hdmasterkeyid": "fff01863811e38af701d29f2a719cae80df7be23"
}

</div>

As you can see, the unconfidential (unblinded) P2PKH address starts with a `2`.

Confidential P2SH addresses start with `Azp`, and unconfidential (unblinded) P2SH addresses start with `X`.

We can check this by first creating a 2-of-2 multisignature address:
~~~~
elements-cli createmultisig 2 '["0222c31615e457119c2cb33821c150585c8b6a571a511d3cd07d27e7571e02c76e", "039bac374a8cd040ed137d0ce837708864e70012ad5766030aee1eb2f067b43d7f"]'
~~~~

Which gives us the unconfidential address:

<div class="console-output">
{
  "address": "XCSVzf6jD4p3GUg1XLxTYkzZvH1CHrjvDA",
  "redeemScript": "52210222c31615e457119c2cb33821c150585c8b6a571a511d3cd07d27e7571e02c76e21039bac374a8cd040ed137d0ce837708864e70012ad5766030aee1eb2f067b43d7f52ae"
}

</div>

And then creating a confidential address by adding a blinding pubkey to it
(Note that the pubkey here is arbitrarily chosen for illustration purpose,
wou will need to use your own unique pubkey corresponding to your blinding private key in practice)

~~~~
elements-cli createblindedaddress XCSVzf6jD4p3GUg1XLxTYkzZvH1CHrjvDA 02b0d67f275cc93ca2ac507375a1112982e8b50a627c3becb66a2ff27bc4fad0ac
~~~~

We get confidential P2SH address as a result:

<div class="console-output">
Azpom1jJ3mGZzLwiB1bqkvAiRn1givCZ8WhWuY2BmUQMqqjq4uzbF9SNWy5icEq2yqsQCUd8u2epStKL

</div>

If we convert these addresses from base58 encoding to hexadecimal,

`XCSVzf6jD4p3GUg1XLxTYkzZvH1CHrjvDA` --> 4b**0bf6d977a489e1ebb4b7963c8a28a08bd70b85ed**
`Azpom1jJ3mGZzLwiB1bqkvAiRn1givCZ8WhWuY2BmUQMqqjq4uzbF9SNWy5icEq2yqsQCUd8u2epStKL` --> 04**4b**02b0d67f275cc93ca2ac507375a1112982e8b50a627c3becb66a2ff27bc4fad0ac**0bf6d977a489e1ebb4b7963c8a28a08bd70b85ed**

We can see that confidential address is comprized of confidential address prefix `04`,
followed by unblinded address prefix `4b`, then follows the blinding pubkey, and then the
rest of bytes of confidential address.

You **must** use the confidential address in ``sendtoaddress``, ``sendfrom``, ``sendmany`` and ``createrawtransaction`` if you want to create confidential transactions. Therefore, when you want to receive confidential transactions you must give the *confidential* address to the sender. For all other RPC's except ``dumpblindingkey`` it does not matter whether the confidential or unconfidential address is provided.

### Find out more about Confidential Transactions

Gregory Maxwell's [original investigation]({{ site.url }}/features/confidential-transactions/investigation) into Confidential Transactions.

Try it yourself in the [Elements Code Tutorial]({{ site.url }}/elements-code-tutorial/confidential-transactions).


