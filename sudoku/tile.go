package sudoku

type Tile struct {
	assignment int
	domain     []int
}

func (t Tile) assign(val int) {
	if contains(t.domain, val) {
		t.assignment = val
	}
}

func (Tile t) getDomain() []Value {
	return domain
}

func (Tile t) removeVal(val Value) {
	t.domain = remove(t.domain, val)
}
