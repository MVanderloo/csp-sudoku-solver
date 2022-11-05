package csp

type Constraint_t int8

const (
	NOT_EQUALS Constraint_t = iota
	SUM
)

type Constraint struct {
	constraint_type Constraint_t
	constrained     []int
	sum             int
}

func NewNotEqualsConstraint(v1, v2 int) Constraint {
	return Constraint{
		constraint_type: NOT_EQUALS,
		constrained:     []int{v1, v2},
	}
}

func NewSumConstraint(sum int, vars ...int) Constraint {
	return Constraint{
		constraint_type: SUM,
		constrained:     vars,
		sum:             sum,
	}
}
