# Epp

**Epp** is a (fast) [EDN][] pretty-printer based on [`go-edn`][go-edn].

[EDN]: https://github.com/edn-format/edn#edn
[go-edn]: https://github.com/go-edn/edn

## Installation

1. Ensure that [Go](https://golang.org) is installed and `~/go/bin` is in your `$PATH`
2. Run `go get github.com/bfontaine/epp`

## Usage

    epp [options] [<filter> [<input-file>]]

Where valid options include:

* `-output <output-file>`: Write in a file instead of `stdout`.
* `-append`: When writing in a file, do it in append mode.

`<filter>` should be a valid filter according to the grammar below. If
`<input-file>` is given, the EDN is read from it instead of `stdin`.

Options and arguments expecting a filename accept `-` as an alias for `stdin`
for reading or `stdout` for writing.

### Filters

Filters follow a grammar similar to [jq](https://stedolan.github.io/jq/)’s.

* `.`: Identity.
* `.field`
* `.0`, `.1`, `.42`
* `keys`: Available keys in a map, in unspecified order.

Filters are separated by spaces.
