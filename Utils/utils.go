package Utils

import (
	"math"
	"strconv"
)

func PrintDay(day int) {
	PrintWithColor(Cyan, "===Day " + IntToString(day) + "===\n")
}

//Check validates the error passed to it and panics if one exists
//@param e -- error to check
func Check(e error) {
	if e != nil {
		panic(e)
	}
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
