package Day14

import (
	"fmt"
	"math"
	"src/Utils"
	"strings"
)


func Run() {
	Utils.PrintDay(14)

	//get input
	input := Utils.ReadInputToSlice("Days/Day14/input.txt")

	//setup
	part1Steps := 10
	part1Total := 0
	part2Steps := 40

	//Create initial mapping of start word
	startWord := input[0]
	lastLetter := startWord[len(startWord)-1:]
	pairMap := make(map[string]int)
	for i, _ := range strings.Split(startWord, "") {
		if i + 1 == len(startWord) { break }
		pair := startWord[i:i+2]
		addPair(pair, pairMap, 1)
	}

	//Read in rules
	rules := make(map[string]string)
	for _, line := range input[2:] {
		parts := strings.Split(line, " -> ")
		rules[parts[0]] = parts[1]
	}

	//Run through steps
	for i := 0; i < part2Steps; i++ {
		newPairs := make(map[string]int)
		for pair, count := range pairMap {
			letterToAdd := rules[pair]
			//add first new pair
			addPair(pair[:1] + letterToAdd, newPairs, count)
			//add second new pair
			addPair(letterToAdd + pair[1:], newPairs, count)
		}
		//fmt.Printf("%v\n", newPairs)
		pairMap = newPairs

		if i == part1Steps - 1 {
			part1Total = calcMinMaxDif(pairMap, lastLetter)
		}
	}

	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1Total, calcMinMaxDif(pairMap, lastLetter))
}

//Calculates the diff between the max char appearing and min char appearing
//@param pairMap -- map of pairs with their counts
//@param lastLetter -- last letter of string to add at end of counting step
//@return diff between max and min letter counts
func calcMinMaxDif(pairMap map[string]int, lastLetter string) int {
	letterCounts := countLetters(pairMap)
	letterCounts[lastLetter]++
	max := -1
	min := math.MaxInt64
	for _, count := range letterCounts {
		if count > max { max = count }
		if count < min { min = count }
	}

	//fmt.Printf("max: %d, min: %d\n", max, min)
	return max - min
}

//Counts the number of times each letter appears in the map
//@param pairMap -- map of pairs with their counts
//@return map of each letter to its count (not including last letter)
func countLetters(pairMap map[string]int) map[string]int {
	letterCounts := make(map[string]int)
	for pair, count := range pairMap {
		addPair(pair[:1], letterCounts, count)
	}
	return letterCounts
}

//Adds pair to map and increments the count
//@param pair -- pair to add
//@param pairMap -- map to add pair to
//@param num -- number to increment count by
func addPair(pair string, pairMap map[string]int, num int) {
	//fmt.Printf("Adding pair: %s\n", pair)
	count, _ := pairMap[pair]
	pairMap[pair] = count + num
}