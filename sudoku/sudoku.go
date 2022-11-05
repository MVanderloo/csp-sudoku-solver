package sudoku

import (
	"Sudoku-CSP/csp"
	"fmt"
)

// defines the size of the sudoku board and the values each tile can take on
// Values range from 1-SIZE inclusive
// the type of Value must be large enough to hold it's size
const BOX_SIZE int = 3
const SIZE int = BOX_SIZE * BOX_SIZE

// Sudoku board is a flat array with length SIZE*SIZE projected as a square 2d array length SIZE
type Sudoku struct {
	arr [SIZE][SIZE]int
}

// Returns a sudoku board with all tiles empty and
func NewSudoku() Sudoku {
	return Sudoku{}
}

// Returns a sudoku board with the specified assignments
func NewSudokuPartial(assignments [][]int) Sudoku {
	var s = NewSudoku()

	for i, e1 := range assignments {
		if e1 == nil {
			continue
		}
		for j, e2 := range e1 {
			if i >= SIZE || j >= SIZE {
				continue
			}

			if e2 < 0 || e2 > SIZE {
				s.arr[i][j] = 0
			} else {
				s.arr[i][j] = assignments[i][j]
			}
		}
	}

	return s
}

func (s Sudoku) ToCSP() csp.CSP {
	var csp_ csp.CSP = csp.NewCSP()
	type Cell struct{ row, col int }
	var id_mapping = make(map[Cell]int)

	for i, row := range s.arr {
		for j, e := range row {
			if e == 0 || (e < 0 || e > 9) {
				id_mapping[Cell{i, j}] = csp_.Insert(0, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
			} else {
				id_mapping[Cell{i, j}] = csp_.Insert(e, nil)
			}
		}
	}

	fmt.Println(id_mapping)

	var constrained [9]int
	var temp []int
	for i := 0; i < SIZE; i++ {

		// row constraints
		for j := range constrained {
			constrained[j] = id_mapping[Cell{i, j}]
		}

		copy(temp, constrained[:])
		csp_.Constrain(0, temp...)

		// column constraints
		for j := range constrained {
			constrained[j] = id_mapping[Cell{j, i}]
		}

		copy(temp, constrained[:])
		csp_.Constrain(0, temp...)
	}

	// box constraints
	for i := 0; i < BOX_SIZE*BOX_SIZE; i += BOX_SIZE {
		for j := 0; j < BOX_SIZE*BOX_SIZE; j += BOX_SIZE {
			constrained[0] = id_mapping[Cell{i, j}]
			constrained[1] = id_mapping[Cell{i + 1, j}]
			constrained[2] = id_mapping[Cell{i + 2, j}]
			constrained[3] = id_mapping[Cell{i, j + 1}]
			constrained[4] = id_mapping[Cell{i + 1, j + 1}]
			constrained[5] = id_mapping[Cell{i + 2, j + 1}]
			constrained[6] = id_mapping[Cell{i, j + 2}]
			constrained[7] = id_mapping[Cell{i + 1, j + 2}]
			constrained[8] = id_mapping[Cell{i + 2, j + 2}]

			copy(temp, constrained[:])
			csp_.Constrain(0, temp...)
		}
	}

	return csp_
}
