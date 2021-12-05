package Utils

import (
	"bufio"
	"os"
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

//Opens and reads in file to slice
//@param filepath -- path to file to read in
//@return slice of strings, each without \n at end
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