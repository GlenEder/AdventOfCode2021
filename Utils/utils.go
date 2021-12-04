package Utils

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	colorReset = "\033[0m"
	colorRed = "\033[31m"
	colorGreen = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan = "\033[36m"
	colorWhite = "\033[37m"

	ColorReset = 0
	Red = 1
	Green = 2
	Yellow = 3
	Blue = 4
	Purple = 5
	Cyan = 6
	White = 7
)

//Prints the string provided in color provided
//@param c -- color to print in
//@param s -- string to print
func PrintWithColor(c int, s string) {
	SetOutputColor(c)
	fmt.Printf("%s", s)
	ResetOutputColor()
}

//Sets output color in terminal
//@param c -- color int to use (see const defined at top of utils.go)
func SetOutputColor(c int)  {
	switch c {
	case Red:
		fmt.Print(colorRed)
		break
	case Green:
		fmt.Print(colorGreen)
		break
	case Yellow:
		fmt.Print(colorYellow)
		break
	case Blue:
		fmt.Print(colorBlue)
		break
	case Purple:
		fmt.Print(colorPurple)
		break
	case Cyan:
		fmt.Print(colorCyan)
		break
	case White:
		fmt.Print(colorWhite)
		break
	}
}

//Returns color code for setting terminal output color
//@param c -- color int to use (see const defined at top of utils.go)
//@return color code for terminal output
func GetColorCode(c int) string {
	switch c {
	case ColorReset:
		return colorReset
	case Red:
		return colorRed
	case Green:
		return colorGreen
	case Yellow:
		return colorYellow
	case Blue:
		return colorBlue
	case Purple:
		return colorPurple
	case Cyan:
		return colorCyan
	case White:
		return colorWhite
	}

	return ""
}

func ResetOutputColor() {
	fmt.Print(colorReset)
}

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

//Returns int rep of string
//@param s -- string to convert
//@return int rep of string
func StringToInt(s string) int {
	num, err := strconv.Atoi(s)
	Check(err)
	return num
}

func IntToString(i int) string  {
	return strconv.Itoa(i)
}

//Coverts binary string into float rep form
//@param b -- binary string to convert
//@return float64 rep of binary number
func BinaryStringToFloat(b string) float64 {
	//Calculate number
	toReturn := 0.0
	numBits := len(b) - 1
	for i, c := range b {
		if c == '1' {
			pos := float64(numBits - i)
			toReturn += math.Exp2(pos)
		}
	}

	return toReturn
}

