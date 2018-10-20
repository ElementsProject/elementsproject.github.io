---
layout: page
title: Additional opcodes
permalink: features/opcodes
---

#  Additional opcodes

## Principal Investigator: Patrick Strateman

Elements includes several new script opcodes, in addition to the ones already supported by Bitcoin:

* Elements reintroduces some safe but disabled opcodes, including string concatenation, substrings, integer shifts, and several bitwise operations.

* A new [DETERMINISTICRANDOM](https://github.com/ElementsProject/elements/search?q=deterministicrandom&unscoped_q=deterministicrandom) operation which produces a random number within a range from a seed.

* A new [CHECKSIGFROMSTACK](https://github.com/ElementsProject/elements/search?q=CHECKSIGFROMSTACK&unscoped_q=CHECKSIGFROMSTACK) operation which verifies a signature against a message on the stack, rather than the spending transaction itself.

These new opcodes have several use cases, including double-spent protection bonds, lotteries, merkle tree constructions to allow 1-of-N multisig with huge N (thousands), and probabilistic payments.

