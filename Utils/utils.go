package Utils

import (
	"fmt"
)

//PrintDay prints the string ===Day n=== to console in color
//@param day -- day number to print
func PrintDay(day int) {
	PrintWithColor(Cyan, "===Day "+IntToString(day)+"===\n")
}

//Check validates the error passed to it and panics if one exists
//@param e -- error to check
func Check(e error) {
	if e != nil {
		panic(e)
	}
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
