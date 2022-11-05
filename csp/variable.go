package csp

import "Sudoku-CSP/util"

type Variable struct {
	assignment int
	domain     []int
}

/**
 * Variable constructor
 **/
func NewVariable(assignment int, domain []int) Variable {
	return Variable{
		assignment: assignment,
		domain:     domain,
	}
}

func NewEmptyVariable(domain []int) Variable {
	return NewVariable(-1, domain)
}

func NewAssignedVariable(assignment int) Variable {
	return NewVariable(assignment, nil)
}

/**
 * Variable constructor.
 **/
func NewSudokuVariable() Variable {
	return Variable{
		domain: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}
}

/**
 * Removes a value from the domain of the variable
 **/
func (v Variable) Remove(value int) {
	v.domain = util.Remove(v.domain, value)
}

/**
 * Checks if a variable can take on a value
 **/
func (v Variable) DomainContains(value int) bool {
	return util.Contains(v.domain, value)
}

// /**
//  * Assigns a value to the tile if it is in it's domain
//  **/
// func (v Variable) Assign(val Value) {
// 	if Contains(v.Domain, val) {
// 		v.Assignment = val
// 		v.Domain = nil
// 	}
// }

// /**
//  * Returns a copy of the domain of the variable
//  **/
// func (x Variable) GetDomain() []Value {
// 	// return x.domain
// 	return x.Domain[:] // copy of domain
// }

// /**
//  * Checks if a variable has an assignment
//  **/
// func (x Variable) IsAssigned() bool {
// 	return len(x.Domain) == 1
// }
