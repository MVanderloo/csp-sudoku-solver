package csp

import "Sudoku-CSP/util"

type Constraint struct {
	constrained []Variable
	sum         int8
}

type Constraint_t int8

const (
	NOT_EQUALS Constraint_t = iota
	SUM
)

func (c Constraint) getType() Constraint_t {
	if c.sum == 0 {
		return NOT_EQUALS
	} else {
		return SUM
	}
}

func NewNotEqualsConstraint(vars ...int8) Constraint {
	var constrained = []Variable{}
	for _, variable := range vars {
		if !util.Contains(constrained, Variable(variable)) {
			constrained = append(constrained, Variable(variable))
		}
	}
	return Constraint{
		constrained: constrained,
		sum:         0,
	}
}

func NewSumConstraint(sum int8, vars ...int8) Constraint {
	var constrained = []Variable{}
	for _, variable := range vars {
		if !util.Contains(constrained, Variable(variable)) {
			constrained = append(constrained, Variable(variable))
		}
	}
	return Constraint{
		constrained: constrained,
		sum:         sum,
	}
}

/**
 * Returns if the variable is constrained by the constraint
 **/
func (c Constraint) constrains(v Variable) bool {
	return util.Contains(c.constrained, v)
}

/**
 * Returns the set of variables constrained by the constraint
 **/
func (c Constraint) GetConstrained() []Variable {
	return c.constrained
}
