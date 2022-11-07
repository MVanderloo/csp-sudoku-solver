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

func NewOverlapSudokuFromString(str string) OverlapSudoku {
	var rows [15]string
	rows[0] = str[0:9]
	rows[1] = str[9:18]
	rows[2] = str[18:27]

	rows[3] = str[27:39]
	rows[4] = str[39:51]
	rows[5] = str[51:63]

	rows[6] = str[63:78]
	rows[7] = str[78:93]
	rows[8] = str[93:108]

	rows[9] = str[108:120]
	rows[10] = str[120:132]
	rows[11] = str[132:144]

	rows[12] = str[144:153]
	rows[13] = str[153:162]
	rows[14] = str[162:171]

	var s1, s2, s3 Sudoku
	s1 = NewSudokuFromString(rows[0] + rows[1] + rows[2] + rows[3][:9] + rows[4][:9] + rows[5][:9] + rows[6][:9] + rows[7][:9] + rows[8][:9])
	s2 = NewSudokuFromString(rows[3][3:12] + rows[4][3:12] + rows[5][3:12] + rows[6][3:12] + rows[7][3:12] + rows[8][3:12] + rows[9][:9] + rows[10][:9] + rows[11][:9])
	s3 = NewSudokuFromString(rows[6][6:15] + rows[7][6:15] + rows[8][6:15] + rows[9][3:12] + rows[10][3:12] + rows[11][3:12] + rows[12] + rows[13] + rows[14])

	return NewOverlapSudoku(s1, s2, s3)
}

func (os OverlapSudoku) ToCSP() csp.CSP {
	csp, id_mapping := os.boards[0].ToCSPWithIds()
	csp1, id_mapping1 := os.boards[1].ToCSPWithIds()
	csp2, id_mapping2 := os.boards[2].ToCSPWithIds()

	for _, variable := range csp1.GetVariables() {
		coord, _ := util.FindKey(id_mapping1, int(variable))

		if coord.row >= 6 || coord.col >= 6 {
			id_mapping[Coord{row: coord.row + 3, col: coord.col + 3}] = int(len(csp.GetVariables()))
			csp.Insert(int(len(csp.GetVariables())), initialDomain(os.boards[1].arr[coord.row][coord.col]))
		}
	}

	for _, constraint := range csp1.GetConstraints() {
		var cons_vars = []int{}
		for _, variable := range constraint.GetConstrained() {
			coord, _ := util.FindKey(id_mapping1, int(variable))
			cons_vars = append(cons_vars, id_mapping[Coord{row: coord.row + 3, col: coord.col + 3}])
		}
		csp.Constrain(cons_vars...)
	}

	for variable := range csp2.GetVariables() {
		coord, _ := util.FindKey(id_mapping2, int(variable))

		if coord.row >= 6 || coord.col >= 6 {
			id_mapping[Coord{row: coord.row + 6, col: coord.col + 6}] = int(len(csp.GetVariables()))
			csp.Insert(int(len(csp.GetVariables())), initialDomain(os.boards[2].arr[coord.row][coord.col]))
		}
	}

	for _, constraint := range csp2.GetConstraints() {
		var cons_vars = []int{}
		for _, variable := range constraint.GetConstrained() {
			coord, _ := util.FindKey(id_mapping2, int(variable))
			cons_vars = append(cons_vars, id_mapping[Coord{row: coord.row + 6, col: coord.col + 6}])
		}
		csp.Constrain(cons_vars...)
	}

	return csp
}
