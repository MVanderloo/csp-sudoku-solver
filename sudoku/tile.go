package sudoku

import (
	"Sudoku-CSP/util"
	"fmt"
)

/**
 * The tile implements the variable interface for the Sudoku puzzle CSP
 **/
type Tile struct {
	assignment Value
	domain     []Value
}

func EmptyTile() Tile {
	// initial domain of the empty tile is values 1-9
	domain := make([]Value, 9)
	for i := range domain {
		domain[i] = Value(i + 1)
	}

	return Tile{
		0,
		domain,
	}
}

func (t Tile) assign(val Value) {
	if util.Contains(t.domain, val) {
		t.assignment = val
	}
}

func (t Tile) getDomain() []Value {
	return t.domain
}

func (t Tile) removeVal(val Value) {
	t.domain = util.Remove(t.domain, val)
}

func (t Tile) print() {
	fmt.Println("printing tile")
}
