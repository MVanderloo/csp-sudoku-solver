package sudoku

// all the types of constraints to define sudoku
type SudokuConstraint_t int8

const (
	NONE      SudokuConstraint_t = iota // 0 Value for constraint
	NOT_EQUAL                           // all the assignments must be unique
)

type SudokuConstraint struct {
	constraint  SudokuConstraint_t
	constrained []*Tile
}

func newRowConstraint(s Sudoku, row int) SudokuConstraint {
	var c = SudokuConstraint{
		constraint:  NOT_EQUAL,
		constrained: make([]*Tile, SIZE),
	}

	for i := range c.constrained {
		c.constrained[i] = s.get(row, i)
	}

	return c
}

func newColConstraint(s Sudoku, col int) SudokuConstraint {
	var c = SudokuConstraint{
		constraint:  NOT_EQUAL,
		constrained: make([]*Tile, SIZE),
	}

	for i := range c.constrained {
		c.constrained[i] = s.get(i, col)
	}

	return c
}

func newBoxConstraint(s Sudoku, box int) SudokuConstraint {
	var c = SudokuConstraint{
		constraint:  NOT_EQUAL,
		constrained: make([]*Tile, SIZE),
	}

	for i := range c.constrained {
		// row and column of the box relative to other boxes
		var boxRow int = box / int(BOX_SIZE)
		var boxCol int = box % int(BOX_SIZE)

		// row and column of the ith element in the box
		var row int = i / int(BOX_SIZE)
		var col int = i % int(BOX_SIZE)

		c.constrained[i] = s.get(row+int(BOX_SIZE)*boxRow, col+int(BOX_SIZE)*boxCol)
	}

	return c
}
