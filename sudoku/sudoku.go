package sudoku

import "fmt"

// defines the size of the sudoku board and the values each tile can take on
// Values range from 1-SIZE inclusive
// the type of Value must be large enough to hold it's size
const BOX_SIZE Value = 3
const SIZE Value = BOX_SIZE * BOX_SIZE

// Sudoku board is a flat array with length SIZE*SIZE projected as a square 2d array length SIZE
type Sudoku struct {
	arr         [SIZE * SIZE]Tile
	constraints [3 * SIZE]SudokuConstraint
}

// Returns a sudoku board with all tiles empty and
func NewSudoku() Sudoku {
	var s Sudoku

	for i := range s.arr {
		s.arr[i] = EmptyTile()
	}

	for i := 0; i < int(SIZE); i++ {
		s.constraints[i] = newColConstraint(s, i)
		s.constraints[i+int(SIZE)] = newRowConstraint(s, i)
		s.constraints[i+2*int(SIZE)] = newBoxConstraint(s, i)
	}

	return s
}

// Returns a sudoku board with
func NewSudokuPartial(assignments [][]Value) Sudoku {
	var s = NewSudoku()

	for rIdx, row := range assignments {
		for cIdx, assignment := range row {
			if assignment != EMPTY {
				s.get(rIdx, cIdx).assign(assignment)
			}
		}
	}

	return s
}

func (s Sudoku) get(rIdx int, cIdx int) *Tile {
	return &s.arr[9*rIdx+cIdx]
}

func (s Sudoku) print() {
	fmt.Println("sudoku print function")
}
