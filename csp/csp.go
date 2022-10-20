package main

import "fmt"

/**
 * Value for the CSP
 * limited to integers to allow for comparison, ordering, and arithmatic operations
 **/
type Value interface {
	int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64
}

/**
 * Variable for the CSP
 * Holds a value that is optional and a domain of values it can take on
 **/
type Variable[V Value] interface {
	assign(V)       // assigns a value to the variable
	getDomain() []V // gets the domain of values it can take on
	removeVal(V)    // removes a value from the domain
	print()         // print the variable
}

type Constraint_t interface {
	comparable // be able to check if a constraint is equal to another
	print()    // print the constraint
}

/**
 * Constraint for the CSP
 **/
type Constraint[C Constraint_t, V Value, X Variable[V]] struct {
	name         C           // the name of the constraint
	constrained  []*X        // the variables that are constrained
	isSatisfied  func() bool // returns true if all the all the variables are assigned and are consistent
	isAdmissable func() bool // returns true if there are not any inconsistencies in the variable assignments
}

/**
 * Prints the name of the constraint followed by the values of the variables it constrains
 * Could be improved by giving variables some identifier of where they are in the puzzle
 **/
func (c Constraint[Constraint_t, Value, Variable]) print() {
	c.name.print()
	fmt.Print(": ")
	for i, val := range c.constrained {
		fmt.Print(val)
		if i != len(c.constrained)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println()
}

/**
 * A constraint satisfaction problem has a collection of variables of type value
 * and a collection of constraints that must be satisfied
 **/
type CSP[V Value, X Variable[V], C Constraint_t] interface {
	getVariables() []X
	getConstraints() []Constraint[C, V, X]
}
