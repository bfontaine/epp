package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

type Epp struct {
	r *bufio.Reader

	indent []string
}

func NewEpp(input *os.File) *Epp {
	return &Epp{
		r:      bufio.NewReader(input),
		indent: make([]string, 0, 8),
	}
}

func (e *Epp) Indent(s string) {
	e.indent = append(e.indent, s)
}

func (e *Epp) Unindent() {
	if len(e.indent) == 0 {
		return
	}

	// we may want to use a stack instead
	e.indent = e.indent[:len(e.indent)-1]
}

func (e *Epp) indentLevel() string {
	if len(e.indent) == 0 {
		return ""
	}
	return e.indent[len(e.indent)-1]
}

func (e *Epp) Pprint(output *os.File) error {
	newline := false

	for {
		b, err := e.r.ReadByte()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		if newline {
			fmt.Fprintf(output, "\n%s", e.indentLevel())
			newline = false
			if b == ' ' {
				continue
			}
		}
		if b == '\n' {
			continue
		}

		output.Write([]byte{b})

		switch b {
		case '{':
			e.Indent(" ")
		case '}':
			e.Unindent()
			newline = true
		case ',':
			newline = true
		}
	}

	return nil
}

func main() {
	flag.Parse()

	e := NewEpp(os.Stdin)
	e.Pprint(os.Stdout)
}
