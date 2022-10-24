package csp

import "Sudoku-CSP/util"

/**
 * Assigns a value to the tile if it is in it's domain
 **/
func (v *Variable) assign(val Value) {
	if util.Contains(v.domain, val) {
		v.assignment = val
		v.domain = nil
	}
}

/**
 * Returns a copy of the domain of the variable
 **/
func (x Variable) getDomain() []Value {
	// return x.domain
	return x.domain[:] // copy of domain
}

/**
 * Removes a value from the domain of the variable
 **/
func (x *Variable) removeVal(val Value) {
	x.domain = util.Remove(x.domain, val)
}
