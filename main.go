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

	err = edn.PPrint(&dst, src, &edn.PPrintOpts{})

	if err != nil {
		log.Fatal(err)
	}

	os.Stdout.Write(dst.Bytes())
}
