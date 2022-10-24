package sudoku

/**
 * Functions that implement the CSP for sudoku
 **/

func (s Sudoku) getVariables() []Tile {
	return s.arr[:]
}

func (s Sudoku) getConstraints() []Constraint {
	return s.constraints[:]
}
