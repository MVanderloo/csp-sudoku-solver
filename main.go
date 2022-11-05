package main

import (
	"Sudoku-CSP/sudoku"
	"fmt"
)

func main() {
	var s sudoku.Sudoku = sudoku.NewSudokuPartial([][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
	})

	fmt.Println(s)

	s.ToCSP().Print()

	// var c = csp.NewCSP()
	// csp.Print(c)
	// c.Insert(4, []int{1, 2})
	// csp.Print(c)
}
