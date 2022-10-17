package sudoku

import "fmt"
import util

type Sudoku struct {
	variables [81]Tile
}

func new() Sudoku {
	var s Sudoku

	return s
}

func newPartial() Sudoku {
	var s Sudoku = new()

}

func (Sudoku s) print() {
	fmt.Println("sudoku print function")
}
