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

type Solveable interface {
	ToCSP() CSP
	Print()
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
func (csp *CSP) Insert(v int, domain Domain) {
	if !util.Contains(csp.variables, Variable(v)) {
		csp.variables = append(csp.variables, Variable(v))
	}
	csp.domains[Variable(v)] = NewDomain(domain)
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

/**
 * Prints the CSP
 **/
func (csp CSP) Print() {
	fmt.Printf("variables(%v): %v\nconstraints(%v): %v\n", len(csp.variables), csp.domains, len(csp.constraints), csp.constraints)
}

func (csp CSP) removeFromDomain(v Variable, value int) {
	csp.domains[v] = csp.domains[v].Remove(value)
}

func (csp CSP) addToDomain(v Variable, value int) {
	csp.domains[v] = csp.domains[v].Add(value)
}

/**
 * Returns the set of all variables that share a constraint with the input variable within the CSP
 */
func (csp CSP) getNeighbors(v Variable) []Variable {
	var neighbors = []Variable{}
	for _, constraint := range csp.constraints { // for each constraint
		if constraint.constrains(v) { // if the variable is included
			for _, neighbor := range constraint.constrained { // for each neighbor in the constraint
				if v != neighbor && !util.Contains(neighbors, neighbor) { // if neighbors does not contain the variable
					neighbors = append(neighbors, neighbor) // append variable to neighbors
				}
			}
		}
	}
	return neighbors
}

func (csp CSP) GetDomain(v Variable) Domain {
	return csp.domains[v]
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

	for v := range assignment {
		if !csp.isConsistent(v, assignment) {
			return false
		}
	}

	return true
}
