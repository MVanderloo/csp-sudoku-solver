package csp

import "Sudoku-CSP/util"

type Constraint struct {
	constrained []int
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
	var constrained = []int{}
	for _, variable := range vars {
		if !util.Contains(constrained, variable) {
			constrained = append(constrained, variable)
		}
	}
	return Constraint{
		constrained: constrained,
		sum:         0,
	}
}

func NewSumConstraint(sum int, vars ...int) Constraint {
	var constrained = []int{}
	for _, variable := range vars {
		if !util.Contains(constrained, variable) {
			constrained = append(constrained, variable)
		}
	}
	return Constraint{
		constrained: constrained,
		sum:         sum,
	}
}

/**
 * Binary constraint holds ids of 2 variables where v1 is constrained by v2
 * All arcs represent a NOT_EQUALS constraint
 **/
type Arc struct {
	x1 int
	x2 int
}

// used for AC-3
func (c Constraint) toArcs() []Arc {
	var arcs = []Arc{}
	switch c.getType() {
	case SUM:
		return arcs // not generating arcs for sum constraints

	case NOT_EQUALS:
		for i, v1 := range c.constrained {
			for j, v2 := range c.constrained {
				if i == j {
					continue
				}

				arcs = append(arcs, Arc{v1, v2})
			}
		}
		return arcs

	default:
		return arcs
	}
}

func (c Constraint) constrains(variable int) bool {
	return util.Contains(c.constrained, variable)
}

func (c Constraint) isSatisfied(assignment Assignment) bool {
	switch c.getType() {
	case NOT_EQUALS:
		for i, e1 := range c.constrained {
			for j, e2 := range c.constrained {
				if i == j {
					continue
				} else if assignment[e1] == assignment[e2] {
					return false
				}
			}
		}

	case SUM:
		var sum int = 0
		for _, e := range c.constrained {
			sum += assignment[e]
		}
		return sum == c.sum

	}

	return true
}
