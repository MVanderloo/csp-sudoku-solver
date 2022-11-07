package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Config struct {
	Log_file    string `json:"log file"`
	Time_limit  int    `json:"time limit"`
	Input_files struct {
		Directory      string `json:"directory"`
		Sudoku         string `json:"sudoku"`
		Killer_sudoku  string `json:"killer sudoku"`
		Overlap_sudoku string `json:"overlap sudoku"`
	} `json:"input files"`
	Inputs []struct {
		Type             int    `json:"type"`
		Presets          []int  `json:"presets"`
		Ac3              []bool `json:"ac3"`
		Forward_checking []bool `json:"forward checking"`
		Mrv_heuristic    []bool `json:"mrv heuristic"`
		Lcv_heuristic    []bool `json:"lcv heuristic"`
	} `json:"inputs"`
}

func ConfigExists(config_file string) bool {
	if _, err := os.Stat(config_file); err == nil {
		return true
	} else {
		return false
	}
}

func ReadConfig(config_file string) Config {
	jsonFile, err := os.Open(config_file)
	if err != nil {
		panic(err)
	}

	defer jsonFile.Close()

	byteValue, err2 := io.ReadAll(jsonFile)
	if err2 != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var config Config
	if err3 := json.Unmarshal(byteValue, &config); err3 != nil {
		fmt.Printf("Invalid %v has invalid format for config file\n", config_file)
		os.Exit(1)
	}

	return config
}

type KillerSudokuConfig struct {
	Inputs []struct {
		Sudoku string `json:"sudoku"`
		Cages  []struct {
			Sum    int8      `json:"sum"`
			Coords [][2]int8 `json:"coords"`
		} `json:"cages"`
	} `json:"inputs"`
}

func ReadKillerSudokuConfig(config_file string) KillerSudokuConfig {
	jsonFile, err := os.Open(config_file)
	if err != nil {
		panic(err)
	}

	defer jsonFile.Close()

	byteValue, err2 := io.ReadAll(jsonFile)
	if err2 != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var config KillerSudokuConfig
	if err3 := json.Unmarshal(byteValue, &config); err3 != nil {
		fmt.Printf("Invalid %v has invalid format for config file\n", config_file)
		os.Exit(1)
	}

	return config
}
