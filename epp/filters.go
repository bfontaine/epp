package epp

import (
	"regexp"
	"strconv"
	"strings"
)

type Filter interface {
	Apply(*PartialEdn) (*PartialEdn, error)
}

// Filters declarations

type identityFilter struct{}
type keysFilter struct{}

type fieldFilter struct {
	name string
}

type indexFilter struct {
	n int
}

type chainFilter struct {
	filters []Filter
}

var _ Filter = identityFilter{}
var _ Filter = fieldFilter{}
var _ Filter = indexFilter{}
var _ Filter = keysFilter{}
var _ Filter = chainFilter{}

// Filters definitions

func (f identityFilter) Apply(p *PartialEdn) (*PartialEdn, error) {
	return p, nil
}

func (f fieldFilter) Apply(p *PartialEdn) (*PartialEdn, error) {
	return p.GetField(f.name)
}

func (f indexFilter) Apply(p *PartialEdn) (*PartialEdn, error) {
	return p.GetIndex(f.n)
}

func (f keysFilter) Apply(p *PartialEdn) (*PartialEdn, error) {
	return p.Fields()
}

func (f chainFilter) Apply(p *PartialEdn) (*PartialEdn, error) {
	var e error

	for _, filter := range f.filters {
		p, e = filter.Apply(p)
		if e != nil {
			return nil, e
		}

		if p == nil {
			p = Nil
		}
	}

	return p, e
}

// end Filters definitions

func IsIdentityFilter(f Filter) bool {
	_, ok := f.(*identityFilter)
	return ok
}

var filtersSepRe = regexp.MustCompile(`\s*\.`)

func ParseFilter(text string) (Filter, error) {
	text = strings.TrimSpace(text)

	startsWithDot := strings.HasPrefix(text, ".")

	if startsWithDot {
		text = text[1:]
	}

	parts := filtersSepRe.Split(text, -1)

	chain := make([]Filter, 0, len(parts))

	for _, part := range parts {
		// identity
		if part == "" {
			continue
		}

		if !startsWithDot {
			// special functions
			switch part {
			case "keys":
				chain = append(chain, keysFilter{})
				continue
			}
		}

		// index
		if n, err := strconv.Atoi(part); err == nil {
			chain = append(chain, indexFilter{n: n})
			continue
		}

		// field
		chain = append(chain, fieldFilter{name: part})
	}

	if len(chain) == 0 {
		return identityFilter{}, nil
	}

	return chainFilter{filters: chain}, nil
}
