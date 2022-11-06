package csp

import (
	"Sudoku-CSP/util"
	"fmt"
)

type CSP struct {
	variables   map[int]Variable
	constraints []Constraint
}

func NewCSP() CSP {
	return CSP{
		variables:   map[int]Variable{},
		constraints: []Constraint{},
	}
}

func (csp CSP) Insert(id int, domain []int) {
	csp.variables[id] = NewVariable(domain)
}

func (csp *CSP) RemoveFromDomain(id int, value int) {
	csp.variables[id] = csp.variables[id].Remove(value)
}

func (csp *CSP) Constrain(vars ...int) {
	csp.constraints = append(csp.constraints, NewNotEqualsConstraint(vars...))
}

func (csp *CSP) ConstrainSum(sum int, vars ...int) {
	csp.constraints = append(csp.constraints, NewSumConstraint(sum, vars...))
}

func (csp CSP) Print() {
	fmt.Printf("variables: %v\nconstraints: %v\n", csp.variables, csp.constraints)
}

func (csp CSP) getNeighbors(variable int) []int {
	var neighbors = []int{}
	for _, constraint := range csp.constraints { // for each constraint
		if constraint.constrains(variable) { // if the variable is included
			for _, neighbor := range constraint.constrained { // for each neighbor in the constraint
				if variable != neighbor && !util.Contains(neighbors, neighbor) { // if neighbors does not contain the variable
					neighbors = append(neighbors, neighbor) // append variable to neighbors
				}
			}
		}
	}
	return neighbors
}
