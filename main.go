package main

import (
	. "Sudoku-CSP/csp"
	"Sudoku-CSP/sudoku"
	"fmt"
)

func main() {
	var s sudoku.Sudoku = sudoku.NewSudoku([][]int{
		{0, 0, 0, 2, 6, 0, 7, 0, 1},
		{6, 8, 0, 0, 7, 0, 0, 9, 0},
		{1, 9, 0, 0, 0, 4, 5, 0, 0},
		{8, 2, 0, 1, 0, 0, 0, 4, 0},
		{0, 0, 4, 6, 0, 2, 9, 0, 0},
		{0, 5, 0, 0, 0, 3, 0, 2, 8},
		{0, 0, 9, 3, 0, 0, 0, 7, 4},
		{0, 4, 0, 0, 5, 0, 0, 3, 6},
		{7, 0, 3, 0, 1, 8, 0, 0, 0},
	})

	fmt.Println(s)

	var csp = NewCSP()
	csp.Insert(0, []int{2})
	csp.Insert(1, []int{1})
	csp.Insert(2, []int{3})
	csp.Insert(3, []int{2})

	csp.Constrain(0, 1)
	csp.Constrain(0, 2)
	csp.Constrain(1, 3)
	csp.Constrain(2, 3)
	//csp.ConstrainSum(8, 0, 1, 2, 3)

	csp.Print()

	res := csp.BacktrackingSearch(false, false, false)
	fmt.Println(res)

}
