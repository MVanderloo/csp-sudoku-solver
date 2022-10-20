package main

import (
	"Sudoku-CSP/sudoku"
)

func main() {
	var s sudoku.Sudoku = sudoku.NewSudokuPartial([][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		make([]int, 9),
		make([]int, 9),
		make([]int, 9),
		make([]int, 9),
		make([]int, 9),
		make([]int, 9),
		make([]int, 9),
		make([]int, 9),
	})
	s.Print()
}
