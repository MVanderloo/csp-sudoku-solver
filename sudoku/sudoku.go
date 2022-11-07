package sudoku

import (
	"Sudoku-CSP/csp"
	"fmt"
	"strconv"
	"strings"
)

const BOX_SIZE int8 = 3
const SIZE int8 = BOX_SIZE * BOX_SIZE

type Sudoku struct {
	arr [SIZE][SIZE]int8
}

type Coord struct{ row, col int8 }

// Returns a sudoku board with the specified assignments
func NewSudoku(arr [][]int8) Sudoku {
	var s = Sudoku{} // all values are 0

	if arr == nil {
		return s
	}

	// input validation
	for i, e1 := range arr {
		if e1 == nil { // if a row is nil it will be 0's
			continue
		}

		for j, e2 := range e1 {
			if int8(i) >= SIZE || int8(j) >= SIZE { // ignore values in array past 9 elements
				continue
			}

			if e2 < 0 || e2 > SIZE { // ignore values not 0-9
				s.arr[i][j] = 0
			} else {
				s.arr[i][j] = arr[i][j]
			}
		}
	}

	return s
}

func NewSudokuFromAssignment(assignments map[int8]int8) Sudoku {
	var ctr int8 = 0
	var s = Sudoku{}
	for i := int8(0); i < SIZE; i++ {
		for j := int8(0); j < SIZE; j++ {
			s.arr[i][j] = assignments[ctr]
			ctr++
		}
	}
	return s
}

func NewSudokuFromString(str string) Sudoku {
	str = strings.Trim(str, " \t\n")
	var ctr int = 0
	var s = NewSudoku([][]int8{})
	for i, row := range s.arr {
		for j := range row {
			intVal, err := strconv.Atoi(str[ctr : ctr+1])
			if err != nil || intVal < 0 || intVal > 9 {
				s.arr[i][j] = 0
			}
			s.arr[i][j] = int8(intVal)
			ctr++
		}
	}
	return s
}

func NewSudokuVariable() csp.Domain {
	return csp.NewDomain([]int8{1, 2, 3, 4, 5, 6, 7, 8, 9})
}

/**
 * Converts a sudoku into a CSP
 **/
func (s Sudoku) ToCSP() csp.CSP {
	csp, _ := s.ToCSPWithIds()
	return csp
}

func (s Sudoku) ToCSPWithIds() (csp.CSP, map[Coord]int8) {
	var csp csp.CSP = csp.NewCSP()

	var id int8 = 0
	var id_mapping = make(map[Coord]int8)

	for i, row := range s.arr {
		for j, e := range row {
			id_mapping[Coord{int8(i), int8(j)}] = id
			if e == 0 || (e < 0 || e > 9) {
				csp.Insert(id, []int8{1, 2, 3, 4, 5, 6, 7, 8, 9})
			} else {
				csp.Insert(id, []int8{e})
			}
			id++
		}
	}

	var constrained [9]int8
	for i := int8(0); i < SIZE; i++ {

		// row constraints
		for j := range constrained {
			constrained[j] = id_mapping[Coord{i, int8(j)}]
		}

		csp.Constrain(constrained[:]...)
		csp.ConstrainSum(45, constrained[:]...)

		// column constraints
		for j := range constrained {
			constrained[j] = id_mapping[Coord{int8(j), i}]
		}

		csp.Constrain(constrained[:]...)
		csp.ConstrainSum(45, constrained[:]...)
	}

	// box constraints
	for i := int8(0); i < BOX_SIZE*BOX_SIZE; i += BOX_SIZE {
		for j := int8(0); j < BOX_SIZE*BOX_SIZE; j += BOX_SIZE {
			constrained[0] = id_mapping[Coord{i, j}] // im not proud of this
			constrained[1] = id_mapping[Coord{i, j + 1}]
			constrained[2] = id_mapping[Coord{i, j + 2}]
			constrained[3] = id_mapping[Coord{i + 1, j}]
			constrained[4] = id_mapping[Coord{i + 1, j + 1}]
			constrained[5] = id_mapping[Coord{i + 1, j + 2}]
			constrained[6] = id_mapping[Coord{i + 2, j}]
			constrained[7] = id_mapping[Coord{i + 2, j + 1}]
			constrained[8] = id_mapping[Coord{i + 2, j + 2}]

			csp.Constrain(constrained[:]...)
			csp.ConstrainSum(45, constrained[:]...)
		}
	}

	return csp, id_mapping
}

func (s Sudoku) Print() {
	for i, row := range s.arr {
		for j, e := range row {
			if j%3 == 0 {
				fmt.Print(" ")
			}
			fmt.Print(e, " ")

		}
		if i%3 == 2 {
			fmt.Println()
		}
		fmt.Println()
	}
}

func initialDomain(value int8) []int8 {
	if value >= 1 && value <= 9 {
		return []int8{value}
	} else {
		return []int8{1, 2, 3, 4, 5, 6, 7, 8, 9}
	}
}
