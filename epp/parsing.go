package epp

import (
	"io"

	edn "gopkg.in/edn.v1"
)

type PartialEdn struct {
	content []byte
}

func (p *PartialEdn) UnmarshalEDN(b []byte) error {
	p.content = b
	return nil
}

func (p *PartialEdn) MarshalEDN() ([]byte, error) {
	return p.content, nil
}

func Parse(r io.Reader) (*PartialEdn, error) {
	var p PartialEdn
	err := edn.NewDecoder(r).Decode(&p)
	return &p, err
}

func (p *PartialEdn) PPrint(w io.Writer) error {
	content, err := edn.MarshalPPrint(p, &edn.PPrintOpts{})
	if err == nil {
		_, err = w.Write(content)
	}
	return err
}
