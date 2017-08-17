package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"gopkg.in/edn.v1"
)

const (
	Version = "0.1.1"
)

func main() {
	var printVersion bool

	flag.BoolVar(&printVersion, "version", false, "print the version and exit")
	flag.Parse()

	if printVersion {
		fmt.Printf("Epp version %s\n", Version)
		return
	}

	input := bufio.NewReader(os.Stdin)
	output := bufio.NewWriter(os.Stdout)
	opts := &edn.PPrintOpts{}

	if err := edn.PPrintStream(output, input, opts); err != nil {
		log.Fatal(err)
	}
	output.Flush()
}
