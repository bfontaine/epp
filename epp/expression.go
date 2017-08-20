package epp

// WIP

type Expression struct {
	s string
}

func (e Expression) Identity() bool {
	return e.s == "" || e.s == "."
}

func ParseExpression(expr string) Expression {
	return Expression{}
}
