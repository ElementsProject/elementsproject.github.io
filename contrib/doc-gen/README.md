
# Generate docs for Elements RPC

This is a golang script, needed for generating the RPC elements documentation.

What is necessary to run this:

1. Install `golang`.
1. Install elements core, set it up to use `elementsregtest`.
1. Run `elementsd`.
1. Run `go run generate.go` while being in contrib/doc-gen, and with `elements-cli` in PATH.
1. Add the generated files to git.
