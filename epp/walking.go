package epp

import (
	"fmt"

	edn "gopkg.in/edn.v1"
)

func (p *PartialEdn) toMap() (map[string]*PartialEdn, error) {
	// unfortunately unmarshaling in a map[string]something doesn't work; the
	// edn package refuses to e.g. unmarshal a keyword in a string.
	m := make(map[interface{}]*PartialEdn)
	err := edn.Unmarshal(p.content, &m)
	if err != nil {
		return nil, err
	}

	sm := make(map[string]*PartialEdn, len(m))
	for k, v := range m {
		switch sk := k.(type) {
		case edn.Keyword:
			// special case not to prepend ":" to the string. "Keyword" is just
			// a type alias for string.
			sm[string(sk)] = v
		case fmt.Stringer:
			sm[sk.String()] = v
		default:
			// we should probably raise an error here
		}
	}

	return sm, nil
}

func (p *PartialEdn) GetField(name string) (*PartialEdn, error) {
	m, err := p.toMap()
	if err != nil {
		return nil, err
	}

	return m[name], nil
}

func (p *PartialEdn) GetIndex(n int) (*PartialEdn, error) {
	ls := make([]*PartialEdn, 0)

	if n < 0 {
		return nil, nil
	}

	err := edn.Unmarshal(p.content, &ls)
	if err != nil {
		return nil, err
	}

	if n > len(ls) {
		return nil, nil
	}

	return ls[n], nil
}

func (p *PartialEdn) Fields() (*PartialEdn, error) {
	m, err := p.toMap()
	if err != nil {
		return nil, err
	}

	fields := make([]string, 0, len(m))
	for field := range m {
		fields = append(fields, field)
	}

	content, err := edn.Marshal(fields)
	if err != nil {
		return nil, err
	}

	return &PartialEdn{content: content}, nil
}
