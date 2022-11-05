package csp

import "fmt"

type CSP struct {
	variables   []Variable
	constraints []Constraint
}

func NewCSP() CSP {
	return CSP{
		variables:   []Variable{},
		constraints: []Constraint{},
	}
}

func (csp *CSP) Insert(assignment int, domain []int) int {
	var id int = len(csp.variables)
	csp.variables = append(csp.variables, NewVariable(assignment, domain))
	return id
}

func (csp *CSP) Constrain(sum int, vars ...int) {
	switch sum {
	case int(NOT_EQUALS):
		for _, v1 := range vars {
			for _, v2 := range vars {
				if v1 == v2 || len(csp.variables) <= v1 || len(csp.variables) <= v2 {
					continue
				}
				csp.constraints = append(csp.constraints, NewNotEqualsConstraint(v1, v2))
			}
		}
	default:
		csp.constraints = append(csp.constraints, NewSumConstraint(sum, vars...))
	}
	return
}

func (csp CSP) Print() {
	// fmt.Println(csp.variables)
	fmt.Println(csp.constraints)
}
