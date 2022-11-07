# Tile Puzzle AI

### CS 420 Project 1
### Michael Vanderloo

This project was done in golang. You will need to install golang and update GOPATH to include this directory. 

## Install Golang

https://go.dev/doc/install

## Run program

This program is built and run using the go command line tool.
You can use

```code
go run .
```

to run the program without generating files

or

```code
go build .
```

to create an executable.

## Configs

The input to the experiment is config.json located in the project folder. If the config isn't present, you can rename config_backup.json to config.json.

The config.json must be valid json.

For the values that are lists, the algorithm will run for each permutation of inputs.

Puzzle type is 1-3 for normal sudoku, overlap sudoku, and killer sudoku.

There are text files for each puzzle type specifying it's inputs. To create your own puzzle you can append it to the text file. The normal and overlap sudoku are a string of digits with 0 as a placeholder. The killer is a json.

## Logging

Everything is printed to the logfile specified in config.json. It will overwrite a file so be careful with this. 