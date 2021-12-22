package Utils

import (
	"math"
	"strconv"
	"strings"
)

//Checks for the string in the given slice
//@param slice -- string slice to look in
//@param s -- string to look for
//@return true if in slice, false otherwise
func SliceContainsString(slice []string, s string) bool {
	for _, a := range slice {
		if a == s {
			return true
		}
	}
	return false
}

//Finds index of string in slice
//@param slice -- slice to look in
//@param s -- string to look for
//@return index of string in slice, -1 if not found
func SliceIndexOfStr(slice []string, s string) int {
	for i, a := range slice {
		if a == s {
			return i
		}
	}
	return -1
}

//Removes white space from string
//@param str -- string to remove whitespace in
//@return string without whitespace
func RemoveWhiteSpace(str string) string {
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\t", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	return str
}

//Checks if the string is in all lowercase letters
//@param s -- string to check
//@return true if all lowercase, otherwise false
func IsLowerCase(s string) bool {
	return s == strings.ToLower(s)
}

//Finds the number of matching chars between the two strings
//@param a -- first string
//@param b -- second string
//@return number of matching chars
func NumMatchingChars(a string, b string) int {
	matches := 0
	bChars := []rune(b)
	for i := range bChars {
		if strings.Contains(a, string(bChars[i])) {
			matches++
		}
	}
	return matches
}

//Counts the number of occurrences of the needle string in the slice
//@param slice -- slice to look in
//@param needle -- string to look for
//@return number of matches in slice
func NumStringMatches(slice []string, needle string) int {
	numMatches := 0
	for _, s := range slice {
		if s == needle {
			numMatches++
		}
	}
	return numMatches
}

//Returns int rep of string
//@param s -- string to convert
//@return int rep of string
func StringToInt(s string) int {
	num, err := strconv.Atoi(s)
	Check(err)
	return num
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

//Coverts binary string into int rep form
//@param b -- binary string to convert
//@return int rep of binary number
func BinaryToInt(b string) int {
	//Calculate number
	toReturn := 0.0
	numBits := len(b) - 1
	for i, c := range b {
		if c == '1' {
			pos := float64(numBits - i)
			toReturn += math.Exp2(pos)
		}
	}

	return int(toReturn)
}

func HexCharToBinary(c rune) string {
	switch c {
	case '0':
		return "0000"
	case '1':
		return "0001"
	case '2':
		return "0010"
	case '3':
		return "0011"
	case '4':
		return "0100"
	case '5':
		return "0101"
	case '6':
		return "0110"
	case '7':
		return "0111"
	case '8':
		return "1000"
	case '9':
		return "1001"
	case 'A', 'a':
		return "1010"
	case 'B', 'b':
		return "1011"
	case 'C', 'c':
		return "1100"
	case 'D', 'd':
		return "1101"
	case 'E', 'e':
		return "1110"
	case 'F', 'f':
		return "1111"
	}
	return ""
}

func HexToBinary(s string) string {
	binary := ""
	for _, c := range s {
		binary += HexCharToBinary(c)
	}
	return binary
}
