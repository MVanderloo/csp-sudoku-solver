package sudoku

type KillerSudoku struct {
	board Sudoku
	cages []Cage
}

type Cage struct {
	sum   int
	cages []struct{ x, y int }
}

// Returns a sudoku board with all tiles empty and
func NewKillerSudoku(board Sudoku, cages []Cage) KillerSudoku {
	return KillerSudoku{
		board: board,
		cages: cages,
	}
}
