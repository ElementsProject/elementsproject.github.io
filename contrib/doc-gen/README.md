
# Generate docs for Elements RPC

This is a golang script, needed for generating the RPC bitcoin documentation

What is necessary to run this:
(1) install golang
(2) install bitcoin core, set it up to use regtest
(3) run bitcoind
(4) run `go run generate.go` while being in contrib/doc-gen, and with bitcoin-cli in PATH
(5) add the generated files to git
