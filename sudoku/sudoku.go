package sudoku

import (
	"Sudoku-CSP/csp"
)

const BOX_SIZE int = 3
const SIZE int = BOX_SIZE * BOX_SIZE

type Sudoku struct {
	arr [SIZE][SIZE]int
}

// Returns a sudoku board with the specified assignments
func NewSudoku(assignments [][]int) Sudoku {
	var s = Sudoku{} // all values are 0

	if assignments == nil {
		return s
	}

	// input validation
	for i, e1 := range assignments {
		if e1 == nil { // if a row is nil it will be 0's
			continue
		}

		for j, e2 := range e1 {
			if i >= SIZE || j >= SIZE { // ignore values in array past 9 elements
				continue
			}

			if e2 < 0 || e2 > SIZE { // ignore values not 0-9
				s.arr[i][j] = 0
			} else {
				s.arr[i][j] = assignments[i][j]
			}
		}
	}

	return s
}

/**
 * Converts a sudoku into a CSP
 **/
func (s Sudoku) ToCSP() csp.CSP {
	var csp csp.CSP = csp.NewCSP()
	type Cell struct{ row, col int }
	var id int = 0
	var id_mapping = make(map[Cell]int)

	for i, row := range s.arr {
		for j, e := range row {
			id_mapping[Cell{i, j}] = id
			if e == 0 || (e < 0 || e > 9) {
				csp.Insert(id, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
			} else {
				csp.Insert(id, []int{e})
			}
		}
	}

	var constrained [9]int
	var temp []int
	for i := 0; i < SIZE; i++ {

		// row constraints
		for j := range constrained {
			constrained[j] = id_mapping[Cell{i, j}]
		}

		copy(temp, constrained[:])
		csp.Constrain(temp...)

		// column constraints
		for j := range constrained {
			constrained[j] = id_mapping[Cell{j, i}]
		}

		copy(temp, constrained[:])
		csp.Constrain(temp...)
	}

	// box constraints
	for i := 0; i < BOX_SIZE*BOX_SIZE; i += BOX_SIZE {
		for j := 0; j < BOX_SIZE*BOX_SIZE; j += BOX_SIZE {
			constrained[0] = id_mapping[Cell{i, j}] // im not proud of this
			constrained[1] = id_mapping[Cell{i + 1, j}]
			constrained[2] = id_mapping[Cell{i + 2, j}]
			constrained[3] = id_mapping[Cell{i, j + 1}]
			constrained[4] = id_mapping[Cell{i + 1, j + 1}]
			constrained[5] = id_mapping[Cell{i + 2, j + 1}]
			constrained[6] = id_mapping[Cell{i, j + 2}]
			constrained[7] = id_mapping[Cell{i + 1, j + 2}]
			constrained[8] = id_mapping[Cell{i + 2, j + 2}]

			copy(temp, constrained[:])
			csp.Constrain(temp...)
		}
	}

	return csp
}
