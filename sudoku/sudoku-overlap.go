package sudoku

import (
	"Sudoku-CSP/csp"
	"Sudoku-CSP/util"
)

type OverlapSudoku struct {
	boards [3]Sudoku
}

// Returns a sudoku board with all tiles empty and
func NewOverlapSudoku(board1 Sudoku, board2 Sudoku, board3 Sudoku) OverlapSudoku {
	return OverlapSudoku{
		boards: [3]Sudoku{board1, board2, board3},
	}
}

func (os OverlapSudoku) ToCSP() csp.CSP {
	csp, id_mapping := os.boards[0].ToCSPWithIds()
	csp1, id_mapping1 := os.boards[1].ToCSPWithIds()
	csp2, id_mapping2 := os.boards[2].ToCSPWithIds()

	for variable := range csp1.GetVars() {
		coord, _ := util.FindKey(id_mapping1, int(variable))

		if coord.row >= 6 || coord.col >= 6 {
			id_mapping[Coord{coord.row + 3, coord.col + 3}] = len(csp.GetVars())
			csp.Insert(len(csp.GetVars()), initialDomain(os.boards[1].arr[coord.row][coord.col]))
		}
	}

	for _, constraint := range csp1.GetConstraints() {
		var cons_vars = []int{}
		for _, variable := range constraint.GetConstrained() {
			coord, _ := util.FindKey(id_mapping1, int(variable))
			cons_vars = append(cons_vars, id_mapping[Coord{coord.row + 3, coord.col + 3}])
		}
		csp.Constrain(cons_vars...)
	}

	for variable := range csp2.GetVars() {
		coord, _ := util.FindKey(id_mapping2, int(variable))

		if coord.row >= 6 || coord.col >= 6 {
			id_mapping[Coord{coord.row + 6, coord.col + 6}] = len(csp.GetVars())
			csp.Insert(len(csp.GetVars()), initialDomain(os.boards[2].arr[coord.row][coord.col]))
		}
	}

	for _, constraint := range csp2.GetConstraints() {
		var cons_vars = []int{}
		for _, variable := range constraint.GetConstrained() {
			coord, _ := util.FindKey(id_mapping2, int(variable))
			cons_vars = append(cons_vars, id_mapping[Coord{coord.row + 6, coord.col + 6}])
		}
		csp.Constrain(cons_vars...)
	}

	return csp
}
