package main

import (
	"bufio"
	"flag"
	"log"
	"os"

	"gopkg.in/bfontaine/edn.v1"
)

func main() {
	flag.Parse()

	input := bufio.NewReader(os.Stdin)
	output := bufio.NewWriter(os.Stdout)
	opts := &edn.PPrintOpts{}

	if err := edn.PPrintStream(output, input, opts); err != nil {
		log.Fatal(err)
	}
	output.Flush()
}
