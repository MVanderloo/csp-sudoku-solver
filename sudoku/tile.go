package sudoku

import (
	"Sudoku-CSP/util"
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
	var domain = make([]Value, 9)
	for i := range domain {
		domain[i] = Value(i + 1)
	}

	return Tile{
		assignment: 0,
		domain:     domain,
	}
}

/**
 * Assigns a value to the tile if it is in it's domain
 **/
func (t *Tile) assign(val Value) {
	if util.Contains(t.domain, val) {
		t.assignment = val
		t.domain = nil
	}
}

func (t Tile) domainContains(val Value) bool {
	return util.Contains(t.domain, val)
}

func (t Tile) isAssigned() bool {
	return t.assignment != EMPTY
}

func (t Tile) getDomain() []Value {
	return t.domain
}

func (t Tile) removeVal(val Value) {
	t.domain = util.Remove(t.domain, val)
}
