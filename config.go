package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Config struct {
	Log_file string `json:"log file"`
	Inputs   []struct {
		Puzzle_type      int  `json:"puzzle type"`
		Puzzle_id        int  `json:"ID"`
		Ac3              bool `json:"ac3"`
		Forward_checking bool `json:"forward checking"`
		Mrv              bool `json:"mrv"`
		Lcv              bool `json:"lcv"`
	}
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
