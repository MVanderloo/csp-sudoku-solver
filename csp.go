package main

/**
 * Value for the CSP
 * For sudoku this is limited to integers,
 * which allow for comparison, ordering, and arithmatic operations
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
}

type Constraint_t interface {
	comparable
}

const (
	EQUAL int8 = iota
	NOT_EQUAL
	SUM_TO
)

/**
 * Constraint for the CSP
 **/
type Constraint[C Constraint_t, V Value, X Variable[V]] struct {
	constraint  C
	constrained []*X
}

type CSP[V Value, X Variable[V], C Constraint_t] interface {
	getVariables() []X
	getConstraints() []Constraint[C, V, X]
}
