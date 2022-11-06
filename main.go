package main

import (
	. "Sudoku-CSP/csp"
	"fmt"
)

func main() {
	// var s sudoku.Sudoku = sudoku.NewSudoku([][]int{
	// 	{1, 2, 3, 4, 5, 6, 7, 8, 9},
	// 	{1, 2, 3, 4, 5, 6, 7, 8, 9},
	// 	{1, 2, 3, 4, 5, 6, 7, 8, 9},
	// 	{1, 2, 3, 4, 5, 6, 7, 8, 9},
	// 	{1, 2, 3, 4, 5, 6, 7, 8, 9},
	// 	{1, 2, 3, 4, 5, 6, 7, 8, 9},
	// 	{1, 2, 3, 4, 5, 6, 7, 8, 9},
	// 	{1, 2, 3, 4, 5, 6, 7, 8, 9},
	// 	{1, 2, 3, 4, 5, 6, 7, 8, 9},
	// })

	// fmt.Println(s)

	var csp = NewCSP()
	csp.Insert(0, []int{1, 2, 3, 4})
	csp.Insert(1, []int{1})
	csp.Insert(2, []int{1, 2, 3, 4})
	csp.Insert(3, []int{1, 2, 3, 4})

	csp.Constrain(0, 1)
	csp.Constrain(0, 2)
	csp.Constrain(1, 3)
	csp.Constrain(2, 3)

	csp.Print()
	fmt.Println()
	csp.AC3()
	csp.Print()

}
