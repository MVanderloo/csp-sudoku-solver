package csp

import "Sudoku-CSP/util"

type Constraint struct {
	constrained []Variable
	sum         int
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

func NewNotEqualsConstraint(vars ...int) Constraint {
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

func NewSumConstraint(sum int, vars ...int) Constraint {
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
func (c Constraint) constrains(variable Variable) bool {
	return util.Contains(c.constrained, variable)
}

/**
 * Returns the set of variables constrained by the constraint
 **/
func (c Constraint) GetConstrained() []Variable {
	return c.constrained
}

// func (c Constraint) isSatisfied(assignment Assignment) bool {
// 	switch c.getType() {
// 	case NOT_EQUALS:
// 		for i, e1 := range c.constrained {
// 			for j, e2 := range c.constrained {
// 				if i == j {
// 					continue
// 				} else if assignment[e1] == assignment[e2] {
// 					return false
// 				}
// 			}
// 		}

// 	case SUM:
// 		var sum int = 0
// 		for _, e := range c.constrained {
// 			sum += assignment[e]
// 		}
// 		return sum == c.sum

// 	}

// 	return true
// }
