package Utils

import "strings"

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
func NumMatchingChars(a string, b string) int  {
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