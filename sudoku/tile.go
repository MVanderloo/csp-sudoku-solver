package sudoku

import (
	. "Sudoku-CSP/csp"
)

func NewSudokuTile() Variable {
	var domain = make([]Value, 9)
	for i := range domain {
		domain[i] = Value(i + 1)
	}

	return NewUnassignedVariable(domain)
}
