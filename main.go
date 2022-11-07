package main

import (
	"Sudoku-CSP/sudoku"
	"fmt"
)

func main() {
	var s sudoku.Sudoku = sudoku.NewSudoku([][]int{
		{0, 0, 0, 6, 0, 0, 4, 0, 0},
		{7, 0, 0, 0, 0, 3, 6, 0, 0},
		{0, 0, 0, 0, 9, 1, 0, 8},
		nil,
		{0, 5, 0, 1, 8, 0, 0, 0, 3},
		{0, 0, 0, 3, 0, 6, 0, 4, 5},
		{0, 4, 0, 2, 0, 0, 0, 6, 0},
		{9, 0, 3},
		{0, 2, 0, 0, 0, 0, 1},
	})

	// var s sudoku.Sudoku = sudoku.NewSudoku([][]int{
	// 	{5, 8, 1, 6, 7, 2, 4, 3, 9},
	// 	{7, 0, 0, 0, 0, 3, 6, 0, 0},
	// 	{0, 0, 0, 0, 9, 1, 0, 8},
	// 	{4},
	// 	{0, 5, 0, 1, 8, 0, 0, 0, 3},
	// 	{0, 0, 0, 3, 0, 6, 0, 4, 5},
	// 	{0, 4, 0, 2, 0, 0, 0, 6, 0},
	// 	{9, 0, 3},
	// 	{0, 2, 0, 0, 0, 0, 1},
	// })

	// var s sudoku.Sudoku = sudoku.NewSudoku([][]int{
	// 	{4, 0, 0, 0, 0, 0, 8, 0, 5},
	// 	{0, 3, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 7, 0, 0, 0, 0, 0},
	// 	{0, 2, 0, 0, 0, 0, 0, 6, 0},
	// 	{0, 0, 0, 0, 8, 0, 4, 0, 0},
	// 	{0, 0, 0, 0, 1, 0, 0, 0, 0},
	// 	{0, 0, 0, 6, 0, 3, 0, 7, 0},
	// 	{5, 0, 0, 2, 0, 0, 0, 0, 0},
	// 	{1, 0, 4, 0, 0, 0, 0, 0, 0},
	// })

	// s.Print()

	// // var csp1 = s.ToCSP()
	// // csp1.AC3()
	// // sudoku.NewSudokuFromAssignment(csp1.BacktrackingSearch(true, false, false)).Print()

	// var csp1 = s.ToCSP()
	// csp1.Print()
	// csp1.AC3()
	// sudoku.NewSudokuFromAssignment(csp1.BacktrackingSearch(true, true, true, true)).Print()

	// var s sudoku.OverlapSudoku = sudoku.NewOverlapSudoku(sudoku.NewSudoku(nil), sudoku.NewSudoku(nil), sudoku.NewSudoku(nil))

	// var s sudoku.KillerSudoku = sudoku.NewKillerSudoku([][]int{
	// 	{0, 2, 0, 6, 0, 8, 0, 0, 0},
	// 	{5, 8, 0, 0, 0, 9, 7, 0, 0},
	// 	{0, 0, 0, 0, 4, 0, 0, 0, 0},

	// 	{9, 7, 0, 0, 0, 0, 5, 0, 0},
	// 	{6, 0, 0, 0, 0, 0, 0, 0, 4},
	// 	{0, 0, 8, 0, 0, 0, 0, 1, 3},

	// 	{0, 0, 0, 0, 2, 0, 0, 0, 0},
	// 	{0, 0, 9, 8, 0, 0, 0, 3, 6},
	// 	{0, 0, 0, 3, 0, 6, 0, 9, 0},
	// }, []sudoku.Cage{
	// 	sudoku.NewCage(12, [][2]int{
	// 		{0, 0},
	// 		{0, 1},
	// 		{0, 2},
	// 	}),
	// })

	csp1 := s.ToCSP()
	assignment := csp1.BacktrackingSearch(true, true, true, true)
	fmt.Println(assignment)
	fmt.Println(csp1.IsSatisfied(assignment))

	// for _, p1 := range [2]bool{true, false} {
	// 	for _, p2 := range [2]bool{true, false} {
	// 		for _, p3 := range [2]bool{true, false} {
	// 			for _, p4 := range [2]bool{true, false} {
	// 				csp1 := s.ToCSP()
	// 				fmt.Println("ac3:", p1, "forward checking:", p2, "mrv:", p3, "lcv:", p4)
	// 				assignment := csp1.BacktrackingSearch(p1, p2, p3, p4)
	// 				if !csp1.IsSatisfied(assignment) {
	// 					sudoku.NewSudokuFromAssignment(assignment).Print()
	// 				}
	// 			}
	// 		}
	// 	}
	// }
}

func readPuzzle(puzzle_type int, index int) {

}
