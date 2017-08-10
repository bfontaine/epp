package main

import (
	"bufio"
	"flag"
	"log"
	"os"

	"github.com/bfontaine/edn"
)

func main() {
	flag.Parse()

	// buffered I/O makes my 55M-file benchmark go 3x times faster
	input := bufio.NewReader(os.Stdin)
	output := bufio.NewWriter(os.Stdout)

	// This takes ~6s on my 55M benchmark file
	err := edn.PPrintStream(output, input, &edn.PPrintOpts{})

	if err != nil {
		log.Fatal(err)
	}
}
