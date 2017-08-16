# Epp

**Epp** is a (fast) [EDN][] pretty-printer based on [`go-edn`][go-edn].

[EDN]: https://github.com/edn-format/edn#edn
[go-edn]: https://github.com/go-edn/edn

## Installation

1. Ensure that [Go](https://golang.org) is installed and `~/go/bin` is in your `$PATH`
2. Run `go get github.com/bfontaine/epp`

## Usage

    cat yourfile.edn | epp


# Implementation Details

This uses my fork of [`go-edn`][my-edn], a version that brings O(1) memory
usage and a +30% speedup compared to the normal one, by using buffered streamed
I/O instead of large bytes buffers.

There’s [a PR][pr] upstream to get these changes included.

[my-edn]: https://github.com/bfontaine/edn/tree/pprint-stream
[pr]: https://github.com/go-edn/edn/pull/7
