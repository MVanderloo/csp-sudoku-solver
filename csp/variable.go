package csp

import "Sudoku-CSP/util"

/**
 * Variable for the CSP
 * Holds an optional value and a domain of values it can take on
 **/
type Variable struct {
	Assignment Value
	Domain     []Value
}

/**
 * Variable constructor
 **/
func NewVariable(assignment Value, domain []Value) Variable {
	return Variable{
		Assignment: assignment,
		Domain:     domain,
	}
}

/**
 * Variable constructor. Sets assignment to EMPTY
 **/
func NewUnassignedVariable(domain []Value) Variable {
	return NewVariable(EMPTY, domain)
}

/**
 * Assigns a value to the tile if it is in it's domain
 **/
func (v *Variable) Assign(val Value) {
	if util.Contains(v.Domain, val) {
		v.Assignment = val
		v.Domain = nil
	}
}

/**
 * Returns a copy of the domain of the variable
 **/
func (x Variable) GetDomain() []Value {
	// return x.domain
	return x.Domain[:] // copy of domain
}

/**
 * Removes a value from the domain of the variable
 **/
func (x *Variable) RemoveVal(val Value) {
	x.Domain = util.Remove(x.Domain, val)
}

/**
 * Checks if a variable has an assignment
 **/
func (x Variable) IsAssigned() bool {
	return x.Assignment != EMPTY
}

/**
 * Checks if a variable can take on a value
 **/
func (x Variable) DomainContains(val Value) bool {
	return util.Contains(x.Domain, val)
}
