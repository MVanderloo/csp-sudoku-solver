package main

import (
	. "Sudoku-CSP/csp"
	. "Sudoku-CSP/sudoku"
	"Sudoku-CSP/util"
	"fmt"
	"log"
	"os"
	"time"
)

const config_file string = "config.json"

var config Config
var ksConfig KillerSudokuConfig
var logger log.Logger
var logfile *os.File
var sudoku_strs []string
var overlap_strs []string

func init() {
	config = ReadConfig(config_file)
	ksConfig = ReadKillerSudokuConfig(config.Input_files.Directory + config.Input_files.Killer_sudoku)
	sudoku_strs = util.GetFileLines(util.OpenFileRead(config.Input_files.Directory + config.Input_files.Sudoku))
	overlap_strs = util.GetFileLines(util.OpenFileRead(config.Input_files.Directory + config.Input_files.Overlap_sudoku))
}

func main() {
	logfile = util.OpenLogFile(config.Log_file)
	defer logfile.Close()
	logger.SetOutput(logfile)
	logger.Println(time.Now().Format("15:04:05 02/01/06"))
	logger.Println("Time limit:", config.Time_limit, "(ms)")
	logger.Print(util.LogFileSpacer())

	var csp CSP

	for _, input := range config.Inputs {
	puzzleIdLoop:
		for _, puzzle_id := range input.Presets {
			for _, ac3 := range input.Ac3 {
				for _, forward_checking := range input.Forward_checking {
					for _, mrv := range input.Mrv_heuristic {
						for _, lcv := range input.Lcv_heuristic {
							switch input.Type {
							case 1:
								if puzzle_id > len(sudoku_strs) {
									continue puzzleIdLoop
								}
								csp = NewSudokuFromString(sudoku_strs[puzzle_id-1]).ToCSP()
							case 2:
								if puzzle_id > len(overlap_strs) {
									continue puzzleIdLoop
								}
								csp = NewOverlapSudokuFromString(overlap_strs[puzzle_id-1]).ToCSP()
							case 3:
								if puzzle_id > len(ksConfig.Inputs) {
									continue puzzleIdLoop
								}
								var cages []Cage
								for _, cage := range ksConfig.Inputs[puzzle_id-1].Cages {
									cages = append(cages, NewCage(cage.Sum, cage.Coords))
								}
								csp = NewKillerSudokuFromString(ksConfig.Inputs[puzzle_id-1].Sudoku, cages).ToCSP()
							default:
								fmt.Println("Invalid puzzle type:", input.Type)
							}
							logger.Println("type:", input.Type, "\npuzzle_id:", puzzle_id, "\ntime limit:", config.Time_limit, "\nac3:", ac3, "\nforward checking:", forward_checking, "\nmrv:", mrv, "\nlcv:", lcv)
							start := time.Now()
							assignment, rec_calls := csp.BacktrackingSearch(ac3, forward_checking, mrv, lcv, time.Duration(config.Time_limit*1e9))
							duration := time.Since(start)
							logger.Println("Backtracking duration:", duration.Milliseconds(), "(ms)")
							logger.Println("Recursive calls:", rec_calls)
							if csp.IsSatisfied(assignment) {
								logger.Println("Success")
							} else {
								logger.Println("No solution found")
							}
							logger.Println(assignment)
							logger.Print(util.LogFileSpacer())
						}
					}
				}
			}
		}
	}
}
