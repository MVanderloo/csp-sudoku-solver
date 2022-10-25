package sudoku

import . "Sudoku-CSP/csp"

/**
 * Functions that implement the CSP for sudoku
 **/

func (s Sudoku) getVariables() []Variable {
	// var variables = make([]*Variable, len(s.arr))
	// for i := range variables {
	// 	variables[i] = &s.arr[i]
	// }

	return s.arr[:]
}

func (s Sudoku) getConstraints() []Constraint {
	return s.constraints[:]
}
