package csp

import (
	"Sudoku-CSP/util"
)

type Variable struct {
	domain []int
}

/**
 * Variable constructor
 **/
func NewVariable(domain []int) Variable {
	if domain == nil {
		panic("Variable must have a non-nil domain")
	}
	return Variable{
		domain: domain,
	}
}

func NewSudokuVariable() Variable {
	return NewVariable([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
}

/**
 * Removes a value from the domain of the variable
 **/
func (v Variable) Remove(value int) Variable {
	v.domain = util.RemoveOrdered(v.domain, value)
	return v
}

/**
 * Checks if a variable can take on a value
 **/
func (v Variable) Contains(value int) bool {
	return util.Contains(v.domain, value)
}

/**
 * Checks if a variable's domain contains a value other than value
 **/
func (v Variable) DomainContainsOtherThan(value int) bool {
	for _, domain_value := range v.domain {
		if domain_value != value {
			return true
		}
	}
	return false
}

// /**
//  * Returns a copy of the domain of the variable
//  **/
// func (v Variable) getDomain() []int {
// 	// return x.domain[:] // copy of domain
// 	return v.domain
// }
