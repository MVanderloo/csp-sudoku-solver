package sudoku

import (
	"Sudoku-CSP/csp"
	"Sudoku-CSP/util"
)

type Tile struct {
	value csp.Variable
}

func EmptyTile() Tile {
	// initial domain of the empty tile is values 1-9
	var domain = make([]Value, 9)
	for i := range domain {
		domain[i] = Value(i + 1)
	}

	return Tile{
		value: csp.Variable{
			assignment: EMPTY,
			domain:     domain,
		},
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
