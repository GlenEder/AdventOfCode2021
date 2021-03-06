package Utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Opens the file at path provided
//@param filepath -- path to file to open
//@return pointer to open file
func OpenFile(filepath string) *os.File {
	f, err := os.Open(filepath)
	Check(err)
	return f
}

//Reads in file to slice
//@param file -- opened file
//@return slice of strings, each without \n at end
func ReadFileToSlice(file *os.File) []string {
	var input []string
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		input = append(input, scan.Text())
	}
	return input
}

//Reads in file to slice
//@param file -- opened file
//@return slice of integers
func ReadFileToIntegerSlice(file *os.File) []int {
	var input []int
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		input = append(input, StringToInt(scan.Text()))
	}
	return input
}

//Opens and reads in file to slice
//@param filepath -- path to file to read in
//@return slice of strings, each without \n at end
func ReadInputToSlice(filepath string) []string {
	f := OpenFile(filepath)
	return ReadFileToSlice(f)
}

//Opens and reads in file to slice
//@param filepath -- path to file to read in
//@return slice of integers
func ReadInputToIntegerSlice(filepath string) []int {
	f := OpenFile(filepath)
	return ReadFileToIntegerSlice(f)
}

//Reads the input provided
//@param filepath -- path to file to read
//@return array of byte data
func ReadInput(filepath string) []byte {
	//Read the file
	data, err := os.ReadFile(filepath)
	Check(err)

	return data
}

//Reads the input provided
//@param filepath -- path to file to read
//@return single string of input
func ReadInputAsString(filepath string) string {
	return string(ReadInput(filepath))
}

//Takes string slice and coverts it to an int slice
//@param strs -- string slice to convert
//@return int slice
func ToIntSlice(strs []string) []int {
	var ints []int
	for _, s := range strs {
		ints = append(ints, StringToInt(s))
	}
	return ints
}

//Reads the file to an double int slice
//@param filepath -- path to input file
//@param delimit -- how the numbers are separated
//@return double int slice of input values
func ReadInputToIntGrid(filepath string, delimit string) [][]int {
	f := OpenFile(filepath)
	scan := bufio.NewScanner(f)
	var input [][]int
	for scan.Scan() {
		inLine := scan.Text()
		intstrings := strings.Split(inLine, delimit)
		var ints []int
		for _, num := range intstrings {
			ints = append(ints, StringToInt(num))
		}
		input = append(input, ints)
	}
	return input
}

func PrintIntGrid(grid [][]int) {
	fmt.Printf("--Printed Grid--\n")
	for _, line := range grid {
		for _, num := range line {
			fmt.Printf("%d ", num)
		}
		fmt.Printf("\n")
	}
}