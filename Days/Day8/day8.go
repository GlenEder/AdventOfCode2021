package Day8

import (
	"fmt"
	"src/Utils"
	"strings"
)

type pattern struct {
	chars string
	length int
	number int 
}

func Run() {
	Utils.PrintDay(8)

	//get input
	input := Utils.ReadInputToSlice("Days/Day8/input.txt")

	part1total := 0
	part2total := 0

	for _, l := range input {
		sections := strings.Split(l, "|")
		patternDef := sections[0]
		output := sections[1]

		//Part 2
		patterns := calcPattern(patternDef)
		//printPattern(patterns)
		outputNumber := ""

		//Part 1
		outputNumbers := strings.Fields(output)
		for _, num := range outputNumbers {
			if checkFor1478(num) { part1total++ }

			//Part 2 calc
			toAdd := getOutputNumber(num, patterns)
			if toAdd < 0 { Utils.PrintError("No matching pattern found\n")}
			outputNumber += Utils.IntToString(toAdd)
		}

		//fmt.Printf("%s: %s\n", output, outputNumber)
		part2total += Utils.StringToInt(outputNumber)
	}

	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1total, part2total)
}

//Checks to see if s reps the number 1, 4, 7, 8
//@param s -- string to check
//@return true if matches pattern for 1, 4, 7, 8
func checkFor1478(s string) bool {
	length := len(s)
	if length == 2 || length == 3 || length == 4 || length == 7 {
		return true
	}
	return false
}

//Returns the matching number for the string pattern
//@param s -- string to find pattern match
//@param pats -- patterns to match with
//@return matching pattern number, -1 if no match found
func getOutputNumber(s string, pats map[int]string) int {
	for num, chars := range pats {
		if len(s) != len(chars) { continue }
		if Utils.NumMatchingChars(s, chars) == len(s) {
			return num
		}
	}
	return -1
}

//Calculates the pattern for part 2
//@param s -- string with all pattern codes
//@return map of number to string with matching chars
func calcPattern(s string) map[int]string {
	pats := make(map[int]string)
	parts := strings.Fields(s)

	//Find 1, 4, 7, 8
	for _, p := range parts {
		switch len(p) {
		case 2:
			pats[1] = p
			break
		case 3:
			pats[7] = p
			break
		case 4:
			pats[4] = p
			break
		case 7:
			pats[8] = p
			break
		}
	}

	// find others off pattern knowns
	for _, p := range parts {
		switch len(p) {
		//6, 9, or 0
		case 6:
			if Utils.NumMatchingChars(p, pats[1]) == 1 {
				pats[6] = p
			} else if Utils.NumMatchingChars(p, pats[4]) == 3 {
				pats[0] = p
			} else {
				pats[9] = p
			}
			break
		//2, 3, 5
		case 5:
			if Utils.NumMatchingChars(p, pats[1]) == 2 {
				pats[3] = p
			} else if Utils.NumMatchingChars(p, pats[4]) == 3 {
				pats[5] = p
			} else {
				pats[2] = p
			}
		}
	}

	return pats
}

func printPattern(m map[int]string) {
	fmt.Printf("--Pattern--\n")
	for num, chars := range m {
		fmt.Printf("\t%s: %d\n", chars, num)
	}
}