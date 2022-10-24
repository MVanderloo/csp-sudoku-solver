package sudoku

import "fmt"

// defines the size of the sudoku board and the values each tile can take on
// Values range from 1-SIZE inclusive
// the type of Value must be large enough to hold it's size
const BOX_SIZE Value = 3
const SIZE Value = BOX_SIZE * BOX_SIZE

// Sudoku board is a flat array with length SIZE*SIZE projected as a square 2d array length SIZE
type Sudoku struct {
	arr         [SIZE * SIZE]Tile
	constraints [3 * SIZE]SudokuConstraint
}

// Returns a sudoku board with all tiles empty and
func NewSudoku() Sudoku {
	var s Sudoku

	for i := range s.arr {
		s.arr[i] = EmptyTile()
	}

	for i := 0; i < int(SIZE); i++ {
		s.constraints[i] = newColConstraint(s, i)
		s.constraints[i+int(SIZE)] = newRowConstraint(s, i)
		s.constraints[i+2*int(SIZE)] = newBoxConstraint(s, i)
	}

	return s
}

// Returns a sudoku board with the specified assignments
func NewSudokuPartial(assignments [][]int) Sudoku {
	var s = NewSudoku()

	if assignments == nil {
		return s
	}

	for rIdx, row := range assignments {
		if row == nil {
			continue
		}

		for cIdx, assignment := range row {
			if Value(assignment) != EMPTY && assignment <= int(SIZE) {
				// fmt.Printf("row: %v, col: %v = %v\n", rIdx, cIdx, assignment)
				s.get(rIdx, cIdx).assign(Value(assignment))
			}
		}
	}

	return s
}

func (s *Sudoku) get(rIdx int, cIdx int) *Tile {
	return &s.arr[int(SIZE)*rIdx+cIdx]
}

func (s *Sudoku) getFromBoxIdx(boxRow, boxCol, row, col int) *Tile {
	return &s.arr[int(SIZE)*(boxRow*int(BOX_SIZE)+row)+(boxCol*int(BOX_SIZE)+col)]
}

// func (s Sudoku) toStr() string {
// 	var sb strings.Builder
// 	tilePrintWidth := 2*int(BOX_SIZE) - 1
// 	horzBoxDivider := "+" + strings.Repeat(strings.Repeat("+"+strings.Repeat("=", tilePrintWidth), int(BOX_SIZE))+"+", int(BOX_SIZE)) + "+\n"
// 	horzRowDivider := "+" + strings.Repeat(strings.Repeat("+"+strings.Repeat("-", tilePrintWidth), int(BOX_SIZE))+"+", int(BOX_SIZE)) + "+\n"
// 	sb.WriteString(horzBoxDivider)
// 	for boxRow := 0; boxRow < int(BOX_SIZE); boxRow++ { // for each row of boxes
// 		for boxCol := 0; boxCol < int(BOX_SIZE); boxCol++ { // for each box in the row
// 			for col := 0; col < int(BOX_SIZE); col++ { // for each row of the tile
// 				sb.WriteString("|")
// 				for row := 0; row < int(BOX_SIZE); row++ { // for each tile in the box row
// 					sb.WriteString("|")
// 					tile := s.getFromBoxIdx(boxRow, boxCol, row, col)

// 					if tile.isAssigned() {
// 						sb.WriteString(strings.Repeat(" ", tilePrintWidth/2) + fmt.Sprint(tile.assignment) + strings.Repeat(" ", tilePrintWidth/2))
// 					} else {
// 						sb.WriteString("     ")
// 					}

// 				}

// 			}
// 			sb.WriteString("|\n")
// 			sb.WriteString(horzRowDivider)
// 		}
// 		// if boxRow != int(BOX_SIZE)-1 {
// 		// 	sb.WriteString(horzRowDivider)
// 		// } else {
// 		// 	sb.WriteString(horzBoxDivider)
// 		// }
// 		sb.WriteString(horzBoxDivider)
// 	}
// 	// 	for box := 0; box < int(BOX_SIZE); box++ { // for each box in the row
// 	// 		for tileIdx := 0; tileIdx < int(BOX_SIZE); tileIdx++ { // for each tile in the box
// 	// 			sb.WriteString("|")
// 	// 			tile := s.get(i, k+j*int(BOX_SIZE))
// 	// 			if tile.isAssigned() {
// 	// 				sb.WriteString("     ")
// 	// 			} else {
// 	// 				for i := Value(1); i < BOX_SIZE; i++ {
// 	// 					if tile.domainContains(i) {
// 	// 						sb.WriteString(fmt.Sprint(i) + " ")
// 	// 					}
// 	// 				}
// 	// 				if tile.domainContains(BOX_SIZE) {
// 	// 					sb.WriteString(fmt.Sprint(i))
// 	// 				} else {
// 	// 					sb.WriteString(fmt.Sprint(" "))
// 	// 				}
// 	// 			}
// 	// 		}
// 	// 		sb.WriteString("\n")

// 	// 		for k := 0; k < int(BOX_SIZE); k++ { // for each tile in the box
// 	// 			sb.WriteString("|")
// 	// 			tile := s.get(i, k+j*int(BOX_SIZE))
// 	// 			if tile.isAssigned() {
// 	// 				sb.WriteString("  " + fmt.Sprint(tile.assignment) + "  ")
// 	// 			} else {
// 	// 				for i := Value(4); i < 3+BOX_SIZE; i++ {
// 	// 					if tile.domainContains(i) {
// 	// 						sb.WriteString(fmt.Sprint(i) + " ")
// 	// 					}
// 	// 				}
// 	// 				if tile.domainContains(BOX_SIZE) {
// 	// 					sb.WriteString(fmt.Sprint(i))
// 	// 				} else {
// 	// 					sb.WriteString(fmt.Sprint(" "))
// 	// 				}
// 	// 			}
// 	// 		}
// 	// 		sb.WriteString("\n")

// 	// 		for k := 0; k < int(BOX_SIZE); k++ { // for each tile in the box
// 	// 			sb.WriteString("|")
// 	// 			tile := s.get(i, k+j*int(BOX_SIZE))
// 	// 			if tile.isAssigned() {
// 	// 				sb.WriteString("     ")
// 	// 			} else {
// 	// 				for i := Value(7); i < 6+BOX_SIZE; i++ {
// 	// 					if tile.domainContains(i) {
// 	// 						sb.WriteString(fmt.Sprint(i) + " ")
// 	// 					}
// 	// 				}
// 	// 				if tile.domainContains(BOX_SIZE) {
// 	// 					sb.WriteString(fmt.Sprint(i))
// 	// 				} else {
// 	// 					sb.WriteString(fmt.Sprint(" "))
// 	// 				}
// 	// 			}
// 	// 		}
// 	// 		sb.WriteString("\n")

// 	// 		// sb.WriteString(fmt.Sprint(tile.assignment))
// 	// 		// if tile.assignment == 0 {
// 	// 		// 	sb.WriteString("  ")
// 	// 		// } else {
// 	// 		// 	sb.WriteString(fmt.Sprint(tile))
// 	// 		// }

// 	// 	}

// 	// 	// 	for j := 0; j < int(SIZE); j++ {
// 	// 	// 		sb.WriteString("| ")

// 	// 	// 		tile := s.get(i, j)
// 	// 	// 		sb.WriteString(fmt.Sprint(tile.assignment))
// 	// 	// 		// if tile.assignment == 0 {
// 	// 	// 		// 	sb.WriteString("  ")
// 	// 	// 		// } else {
// 	// 	// 		// 	sb.WriteString(fmt.Sprint(tile))
// 	// 	// 		// }

// 	// 	// 		sb.WriteString(" ")
// 	// 	// 	}

// 	// 	// 	for j := 0; j < int(SIZE); j++ {
// 	// 	// 		sb.WriteString("| ")

// 	// 	// 		tile := s.get(i, j)
// 	// 	// 		sb.WriteString(fmt.Sprint(tile.assignment))
// 	// 	// 		// if tile.assignment == 0 {
// 	// 	// 		// 	sb.WriteString("  ")
// 	// 	// 		// } else {
// 	// 	// 		// 	sb.WriteString(fmt.Sprint(tile))
// 	// 	// 		// }

// 	// 	// 		sb.WriteString(" ")
// 	// 	// 	}

// 	// 	sb.WriteString("\n")
// 	// }

// 	return sb.String()
// }

func (s Sudoku) Print() {
	fmt.Println("no")
}
