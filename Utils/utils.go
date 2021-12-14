package Utils

import (
	"fmt"
	"math"
	"strconv"
)

//PrintDay prints the string ===Day n=== to console in color
//@param day -- day number to print
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

//IntToString converts given int to a string rep of the int
//@param i -- int to convert
//@return string rep of int
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

//Calculates the mean of the int slice
//@param ints -- int slice to calc on
//@return mean of int slice
func Mean(ints []int) int {
	sum := 0
	for _, n := range ints {
		sum += n
	}
	return sum / len(ints)
}

//Calculates the mean of the int slice, rounds down
//@param ints -- int slice to calc on
//@return mean of int slice
func LowMean(ints []int) int {
	sum := 0
	for _, n := range ints {
		sum += n
	}
	return int(math.Floor(float64(sum) / float64(len(ints))))
}

//Calculates the mean of the int slice, rounds up
//@param ints -- int slice to calc on
//@return mean of int slice
func HighMean(ints []int) int {
	sum := 0
	for _, n := range ints {
		sum += n
	}
	return int(math.Ceil(float64(sum) / float64(len(ints))))
}

//Finds the median of the integer slice
//@param ints -- sorted integer slice
//@return median of slice
func Median(ints []int) int {
	length := len(ints)
	middle := length / 2
	if length % 2 != 0 {
		return ints[middle]
	}
	return (ints[middle - 1] + ints[middle]) / 2
}

//AbsInt returns the absolute value of an int
//@param i -- int to calc abs on
//@return abs of i
func AbsInt(i int) int {
	if i >= 0 { return i }
	return i * -1
}

//Returns the lower of two numbers
//@param a -- first num to compare
//@param b -- second num to compare
//@return lower of a and b
func MinInt(a int, b int) int {
	if a <= b { return a }
	return b
}

//Returns the higher of two numbers
//@param a -- first num to compare
//@param b -- second num to compare
//@return higher of a and b
func MaxInt(a int, b int) int {
	if a >= b { return a }
	return b
}


//Calculates the product of the ints in the slice
//@param slice -- slice to calc with
//@return product of slice values
func ProductOfIntSlice(slice []int) int {
	product := 1
	for _, num := range slice {
		product *= num
	}
	return product
}

//Prints with desired tabs placed in front of string
//@param depth -- number of tabs to add
//@param s -- string formatting
//@param a -- args for string
func PrintRecursionLogf(depth int, s string, a ...interface{}) {
	for i := 0; i < depth; i++ {
		fmt.Printf("\t")
	}
	fmt.Printf(s, a...)
}
