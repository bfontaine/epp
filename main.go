package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	edn "gopkg.in/edn.v1"

	"github.com/bfontaine/epp/epp"
)

const (
	Version = "0.1.2"
)

func printUsage() {
	fmt.Fprintf(os.Stderr,
		`Epp usage:
%s [options] [<filter> [<input>]]

The default filter is "." and the default input is stdin.
No filter other than the default one is supported for now.

Valid options:
`, os.Args[0])
	flag.PrintDefaults()
}

func main() {
	var printVersion, appendOutput bool
	var outputFilename string

	flag.Usage = printUsage

	flag.StringVar(&outputFilename, "output", "-", "output file")
	flag.BoolVar(&appendOutput, "append", false,
		"open the output file in append mode")

	flag.BoolVar(&printVersion, "version", false, "print the version and exit")
	flag.Parse()

	if printVersion {
		fmt.Printf("Epp version %s\n", Version)
		return
	}

	var input *bufio.Reader
	var output *bufio.Writer

	filterText := "."
	inputFilename := "-"

	narg := flag.NArg()
	// $ epp [options]
	if narg > 0 {
		filterText = flag.Arg(0)
		// $ epp [options] <filter>
		if narg > 1 {
			// $ epp [options] <filter> <filename>
			inputFilename = flag.Arg(1)
			if narg > 2 {
				flag.Usage()
				os.Exit(1)
			}
		}
	}

	if inputFilename == "-" {
		input = bufio.NewReader(os.Stdin)
	} else {
		f, err := os.Open(inputFilename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening '%s': %v", inputFilename, err)
			os.Exit(1)
		}
		defer f.Close()
		input = bufio.NewReader(f)
	}

	if outputFilename == "-" {
		output = bufio.NewWriter(os.Stdout)
	} else {
		mode := os.O_CREATE
		if appendOutput {
			mode = mode | os.O_APPEND
		} else {
			mode = mode | os.O_WRONLY | os.O_TRUNC
		}

		f, err := os.OpenFile(outputFilename, mode, 0666)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening '%s': %v\n", outputFilename, err)
			os.Exit(1)
		}
		defer f.Close()
		output = bufio.NewWriter(f)
	}

	filter, err := epp.ParseFilter(filterText)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing filter '%s': %v\n", filterText, err)
		os.Exit(2)
	}

	// shortcut
	if epp.IsIdentityFilter(filter) {
		err := edn.PPrintStream(output, input, &edn.PPrintOpts{})
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error pretty-printing EDN: %v\n", err)
			os.Exit(5)
		}
		return
	}

	p, err := epp.Parse(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing EDN: %v\n", err)
		os.Exit(3)
	}

	p, err = filter.Apply(p)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error filtering EDN: %v\n", err)
		os.Exit(4)
	}

	err = p.PPrint(output)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error pretty-printing EDN: %v\n", err)
		os.Exit(5)
	}

	output.Flush()
}
