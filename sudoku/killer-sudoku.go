package sudoku

import (
	"Sudoku-CSP/csp"
)

type KillerSudoku struct {
	board Sudoku
	cages []Cage
}

type Cage struct {
	sum   int
	cells []Coord
}

func NewCage(sum int, pairs [][2]int) Cage {
	var cells = []Coord{}
	for _, pair := range pairs {
		cells = append(cells, Coord{row: pair[0], col: pair[1]})
	}
	return Cage{sum: sum, cells: cells}
}

// Returns a sudoku board with all tiles empty and
func NewKillerSudoku(arr [][]int, cages []Cage) KillerSudoku {
	return KillerSudoku{
		board: NewSudoku(arr),
		cages: cages,
	}
}

func NewKillerSudokuFromString(str string, cages []Cage) KillerSudoku {
	return KillerSudoku{
		board: NewSudokuFromString(str),
		cages: cages,
	}
}

func (ks KillerSudoku) ToCSP() csp.CSP {
	csp_, id_mapping := ks.board.ToCSPWithIds()
	var vars []int
	for _, cage := range ks.cages {
		vars = make([]int, len(cage.cells))
		for i, cell := range cage.cells {
			vars[i] = id_mapping[cell]
		}
		csp_.ConstrainSum(cage.sum, vars...)
	}

	return csp_
}

func (ks KillerSudoku) Print() {
	ks.board.Print()
}
