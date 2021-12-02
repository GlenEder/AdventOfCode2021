package Utils

import (
	"bufio"
	"os"
	"strconv"
)

//Check validates the error passed to it and panics if one exists
//@param e -- error to check
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

//Opens the file at path provided
//@param filepath -- path to file to open
//@return pointer to open file
func OpenFile(filepath string) *os.File {
	f, err := os.Open(filepath)
	Check(err)
	return f
}

func ReadFileToSlice(file *os.File) []string {
	var input []string
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		input = append(input, scan.Text())
	}
	return input
}

func ReadInputToSlice(filepath string) []string {
	f := OpenFile(filepath)
	return ReadFileToSlice(f)
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

//Returns int rep of string
//@param s -- string to convert
//@return int rep of string
func StringToInt(s string) int {
	num, err := strconv.Atoi(s)
	Check(err)
	return num
}
