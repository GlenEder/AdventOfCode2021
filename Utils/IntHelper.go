package Utils

import (
	"fmt"
	"math"
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

//Clamps the int to be within the range provided
//@param i -- int to clamp
//@param r -- range to enforce
//@return i if in range, otherwise upper or lower of range depending on overflow direction
func ClampInt(i int, r Range) int {
	if i < r.Lower {
		return r.Lower
	} else if i > r.Upper {
		return r.Upper
	} else {
		return i
	}
}

//Checks if i is negative
//@param i -- int to force
//@return 0 if i is negative, otherwise i
func ForcePositiveInt(i int) int {
	if i < 0 {
		return 0
	}
	return i
}

//Checks if i is positive
//@param i -- int to force
//@return 0 if i is positive, otherwise i
func ForceNegativeInt(i int) int {
	if i > 0 {
		return 0
	}
	return i
}

//Calculates the quadratic forumla for ax^2 + bx + c = 0
//@param a -- a int in forumla
//@param b -- b int in forumla
//@param c -- c int in forumla
//@return two answers for solution, 0,0,error if error found
func Quadratic(a int, b int, c int) (int, int, error) {
	if a == 0 {
		return 0, 0, fmt.Errorf("a cannot be 0")
	}
	bottom := 2 * a
	negB := -1 * b
	toRoot := (b * b) - (4 * a * c)
	if toRoot < 0 {
		return 0, 0, fmt.Errorf("result is imaginary")
	}
	root := int(math.Sqrt(float64(toRoot)))
	return (negB + root) / bottom, (negB - root) / bottom, nil
}
