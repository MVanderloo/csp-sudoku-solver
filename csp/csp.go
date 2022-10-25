package csp

/**
 * Value for the CSP
 **/
type Value int8

const EMPTY Value = 0

type Constraint_t int8 // different types of constraints

type Constraint struct {
	name        Constraint_t
	constrained []*Variable
}

/**
 * A constraint satisfaction problem has a collection of variables of type value
 * and a collection of constraints that must be satisfied
 **/
type CSP interface {
	getVariables() []*Variable
	getConstraintsOf(*Variable) []Constraint
}
