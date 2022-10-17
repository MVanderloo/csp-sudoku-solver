package main

/**
 * Value for the CSP
 * For sudoku this is limited to integers, which allow for comparison, ordering, and arithmatic
 **/
type Value interface {
  Integer
}

/**
 * Variable for the CSP
 **/
type Variable interface {
  assign(Value)
  getDomain() []Value
  removeVal(Value)
}

/**
 * Different types of constraints for the CSP
 **/
type Constraint_t[T comparable] interface {
  type T
}

/**
 * Constraint for the CSP
 **/
type Constraint[T any] struct {
  type Constraint_t
  constrained []*Variables
}

type CSP interface {
  getVariables() []Variable
  getConstraints() []Constraints
}
