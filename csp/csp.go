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
 * Holds an optional value and a domain of
 **/
type Variable[V Value] interface {
	assign(V)
	getDomain() []V
	removeVal(V)
	print()
}

type Constraint_t interface {
	comparable
	print()
}

/**
 * Constraint for the CSP
 **/
type Constraint[C Constraint_t, V Value, X Variable[V]] struct {
	name        C
	constrained []*X
}

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

type CSP[V Value, X Variable[V], C Constraint_t] interface {
	getVariables() []X
	getConstraints() []Constraint[C, V, X]
}
