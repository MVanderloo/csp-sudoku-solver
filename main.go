package main

import (
	"Sudoku-CSP/sudoku"
)

func main() {
	// var s sudoku.Sudoku = sudoku.NewSudoku([][]int{
	// 	{0, 0, 0, 6, 0, 0, 4, 0, 0},
	// 	{7, 0, 0, 0, 0, 3, 6, 0, 0},
	// 	{0, 0, 0, 0, 9, 1, 0, 8},
	// 	nil,
	// 	{0, 5, 0, 1, 8, 0, 0, 0, 3},
	// 	{0, 0, 0, 3, 0, 6, 0, 4, 5},
	// 	{0, 4, 0, 2, 0, 0, 0, 6, 0},
	// 	{9, 0, 3},
	// 	{0, 2, 0, 0, 0, 0, 1},
	// })

	var s sudoku.Sudoku = sudoku.NewSudoku([][]int{
		{5, 8, 1, 6, 7, 2, 4, 3, 9},
		{7, 0, 0, 0, 0, 3, 6, 0, 0},
		{0, 0, 0, 0, 9, 1, 0, 8},
		{4},
		{0, 5, 0, 1, 8, 0, 0, 0, 3},
		{0, 0, 0, 3, 0, 6, 0, 4, 5},
		{0, 4, 0, 2, 0, 0, 0, 6, 0},
		{9, 0, 3},
		{0, 2, 0, 0, 0, 0, 1},
	})

	// s.Print()

	// var csp1 = s.ToCSP()
	// csp1.AC3()
	// sudoku.NewSudokuFromAssignment(csp1.BacktrackingSearch(true, false, false)).Print()

	var csp2 = s.ToCSP()
	csp2.AC3()
	sudoku.NewSudokuFromAssignment(csp2.BacktrackingSearch(true, true, false)).Print()
}
