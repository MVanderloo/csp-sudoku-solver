package util

import (
	"bufio"
	"log"
	"os"
)

func Contains[T comparable](arr []T, val T) bool {
	if arr == nil {
		return false
	}

	for _, v := range arr {
		if v == val {
			return true
		}
	}

	return false
}

func Remove[T comparable](arr []T, val T) []T {
	if len(arr) == 0 {
		return arr
	}

	for i, v := range arr {
		if v == val {
			arr[i] = arr[len(arr)-1]
			arr = arr[:len(arr)-1]
			return arr
		}
	}

	return arr
}

func RemoveOrdered[T comparable](arr []T, val T) []T {
	var ret = make([]T, 0)

	for _, e := range arr {
		if e != val {
			ret = append(ret, e)
		}
	}

	return ret
}

func RemoveFirst[T any](arr []T) (T, []T) {
	if len(arr) == 0 {
		var zero_val T
		return zero_val, arr
	}

	return arr[0], arr[1:]
}

func RemoveLast[T any](arr []T) (T, []T) {
	if len(arr) == 0 {
		var zero_val T
		return zero_val, arr
	}

	idx := len(arr) - 1

	return arr[idx], arr[:idx]
}

func FindKey[T1 comparable, T2 comparable](m map[T1]T2, value T2) (key T1, ok bool) {
	for k, v := range m {
		if v == value {
			key = k
			ok = true
			return
		}
	}
	return
}

func MinSlice[T int](arr []T) T {
	if len(arr) == 0 {
		panic("")
	}

	var min T = arr[0]
	for _, e := range arr {
		if e < min {
			min = e
		}
	}

	return min
}

func FileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	return err == nil
}

/**
 * Reads a file line by line
 **/
func GetFileLines(f *os.File) []string {
	var lines = make([]string, 0)
	sc := bufio.NewScanner(f)

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func OpenFileRead(filepath string) *os.File {
	f, err := os.OpenFile(filepath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return nil
	}
	return f
}

func OpenLogFile(filename string) *os.File {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)

	if err != nil {
		panic(err)
	}

	return f
}

func LogFileSpacer() string {
	return "\n+----------------------------------+\n\n"
}
