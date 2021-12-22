package Utils

import (
	"strconv"
)

//IntToString converts given int to a string rep of the int
//@param i -- int to convert
//@return string rep of int
func IntToString(i int) string {
	return strconv.Itoa(i)
}

//Coverts int to binary string
//@param i -- int to covert
//@return binary string rep of i
func ToBinaryString(i int) string {
	return strconv.FormatInt(int64(i), 2)
}

//AbsInt returns the absolute value of an int
//@param i -- int to calc abs on
//@return abs of i
func AbsInt(i int) int {
	if i >= 0 {
		return i
	}
	return i * -1
}

//Returns the lower of two numbers
//@param a -- first num to compare
//@param b -- second num to compare
//@return lower of a and b
func MinInt(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}

//Returns the higher of two numbers
//@param a -- first num to compare
//@param b -- second num to compare
//@return higher of a and b
func MaxInt(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

//Converts hex string to int
//@param hex -- hex string to covert
//@return int rep of hex string
func HexToInt(hex string) int {
	i, err := strconv.ParseInt(hex, 16, 64)
	Check(err)
	return int(i)
}
