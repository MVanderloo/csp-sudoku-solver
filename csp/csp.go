package csp

import (
	"Sudoku-CSP/util"
	"fmt"
)

type CSP struct {
	variables   []Variable
	domains     map[Variable]Domain
	constraints []Constraint
}

func NewCSP() CSP {
	return CSP{
		variables:   []Variable{},
		domains:     map[Variable]Domain{},
		constraints: []Constraint{},
	}
}

/**
 * Inserts a variable with a domain into the CSP
 */
func (csp *CSP) Insert(variable int, domain Domain) {
	if !util.Contains(csp.variables, Variable(variable)) {
		csp.variables = append(csp.variables, Variable(variable))
	}
	csp.domains[Variable(variable)] = NewDomain(domain)
}

/**
 * Inserts a NOT_EQUALS constraint across vars into the CSP
 **/
func (csp *CSP) Constrain(vars ...int) {
	csp.constraints = append(csp.constraints, NewNotEqualsConstraint(vars...))
}

/**
 * Inserts a SUM constraint of sum across vars into the CSP
 **/
func (csp *CSP) ConstrainSum(sum int, vars ...int) {
	csp.constraints = append(csp.constraints, NewSumConstraint(sum, vars...))
}

func (csp CSP) removeFromDomain(variable Variable, value int) {
	csp.domains[variable] = csp.domains[variable].Remove(value)
}

func (csp CSP) addToDomain(variable Variable, value int) {
	csp.domains[variable] = csp.domains[variable].Add(value)
}

func (csp CSP) Print() {
	fmt.Printf("variables(%v): %v\nconstraints(%v): %v\n", len(csp.domains), csp.domains, len(csp.constraints), csp.constraints)
}

/**
 * Returns the set of all variables that share a constraint with the input variable within the CSP
 */
func (csp CSP) getNeighbors(variable Variable) []Variable {
	var neighbors = []Variable{}
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

func (csp CSP) GetVar(variable Variable) Domain {
	return csp.domains[variable]
}

func (csp CSP) GetVars() map[Variable]Domain {
	return csp.domains
}

func (csp CSP) GetConstraints() []Constraint {
	return csp.constraints
}

func (csp CSP) IsSatisfied(assignment Assignment) bool {
	if !csp.isComplete(assignment) {
		return false
	}

	for k, _ := range assignment {
		if !csp.isConsistent(k, assignment) {
			return false
		}
	}

	return true
}
