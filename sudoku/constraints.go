package sudoku

import . "Sudoku-CSP/csp"

const (
	NONE       Constraint_t = iota // 0 Value for constraint
	ALL_DIFF                       // All the listed values must be different
	ALL_SAME                       // All the listed values must be the same
	EQUALS                         // Both values must be the same
	NOT_EQUALS                     // Both values must be different
	SUM                            // Values must sum up to an amount
)

// func newRowConstraint(s Sudoku, row int) Constraint {
// 	var constraint = NWayConstraint{
// 		constraint: ALL_DIFF,
// 		xn:         make([]*Tile, SIZE),
// 	}

// 	for i := range constraint.xn {
// 		constraint.xn[i] = s.get(row, i)
// 	}

// 	return constraint
// }

// func newColConstraint(s Sudoku, col int) Constraint {
// 	var c = Constraint{
// 		constraint:  NOT_EQUAL,
// 		constrained: make([]*Tile, SIZE),
// 	}

// 	for i := range c.constrained {
// 		c.constrained[i] = s.get(i, col)
// 	}

// 	return c
// }

// func newBoxConstraint(s Sudoku, box int) Constraint {
// 	var c = Constraint{
// 		constraint:  NOT_EQUAL,
// 		constrained: make([]*Tile, SIZE),
// 	}

// 	for i := range c.constrained {
// 		// row and column of the box relative to other boxes
// 		var boxRow int = box / int(BOX_SIZE)
// 		var boxCol int = box % int(BOX_SIZE)

// 		// row and column of the ith element in the box
// 		var row int = i / int(BOX_SIZE)
// 		var col int = i % int(BOX_SIZE)

// 		c.constrained[i] = s.get(row+int(BOX_SIZE)*boxRow, col+int(BOX_SIZE)*boxCol)
// 	}

// 	return c
// }
