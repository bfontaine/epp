package epp

// WIP

type Filter struct {
	s string
}

func (f Filter) Identity() bool {
	return f.s == "" || f.s == "."
}

func ParseFilter(text string) (Filter, error) {
	return Filter{}, nil
}
