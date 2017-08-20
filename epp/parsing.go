package epp

import (
	"bytes"
	"io"
	"io/ioutil"

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

	// unfortunately Unmarshal takes a bytes array instead of a reader
	content, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	err = edn.Unmarshal(content, &p)
	return &p, err
}

func (p *PartialEdn) PPrint(w io.Writer) error {
	content, err := edn.Marshal(p)
	if err != nil {
		return err
	}

	err = edn.PPrintStream(w, bytes.NewBuffer(p.content), &edn.PPrintOpts{})
	if err != nil {
		return err
	}
	_, err = w.Write(content)
	return err
}
