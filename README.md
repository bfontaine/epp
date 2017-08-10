# Epp

**Epp** is a (fast) [EDN][] pretty-printer based on [`go-edn`][go-edn].

[EDN]: https://github.com/edn-format/edn#edn
[go-edn]: https://github.com/go-edn/edn

**Note:** This is an experimental branch that uses my fork of
[`go-edn`][my-edn], a version that brings a +30% speedup compared to the normal
one, by using buffered streamed I/O instead of large bytes buffers. This is
also more memory-efficient.

In order to get this version to build; you need to do the following steps:

    go get -d github.com/bfontaine/edn
    git -C "$GOPATH/src/github.com/bfontaine/edn" checkout pprint-stream
    go get github.com/bfontaine/edn

Then run  `go build .` in this directory.

[my-edn]: https://github.com/bfontaine/edn/tree/pprint-stream

## Installation

1. Ensure that [Go](https://golang.org) is installed and `~/go/bin` is in your `$PATH`
2. Run `go get github.com/bfontaine/epp`

## Usage

    cat yourfile.edn | epp


