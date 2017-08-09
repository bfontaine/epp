package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"os"

	edn "gopkg.in/edn.v1"
)

func main() {
	flag.Parse()

	var dst bytes.Buffer

	src, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	// This takes ~9s on my 55M benchmark file; we might be able to speed up
	// things by concurrently reading on stdin and writing on stdout instead of
	// these read+pprint+write steps.
	err = edn.PPrint(&dst, src, &edn.PPrintOpts{})

	if err != nil {
		log.Fatal(err)
	}

	os.Stdout.Write(dst.Bytes())
}
