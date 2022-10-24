package csp

/**
 * Value for the CSP
 **/
type Value int8

const EMPTY Value = 0

/**
 * Variable for the CSP
 * Holds an optional value and a domain of values it can take on
 **/
type Variable struct {
	assignment Value
	domain     []Value
}

type Constraint_t int8 // different types of constraints

type Constraint struct {
	name          Constraint_t
	constrained   *Variable
	constrainedBy []*Variable
}

/**
 * A constraint satisfaction problem has a collection of variables of type value
 * and a collection of constraints that must be satisfied
 **/
type CSP interface {
	getVariables() []Variable
	getConstraintsOf(Variable, Variable) []Constraint
}
