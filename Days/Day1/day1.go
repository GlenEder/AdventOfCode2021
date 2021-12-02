package Day1

import (
	"fmt"
	"src/Utils"
)


func Day1() {
	//get input
	input := Utils.ReadInputToSlice("Days/Day1/input.txt")

	prev := -1
	count := -1
	for _, in := range input {
		num := Utils.StringToInt(in)
		if num > prev {count++}
		prev = num
	}
	fmt.Printf("Part 1: %d\n", count)

	//Reset for part 2
	count = 0
	fGroup := [3]int{0, 0, 0}
	sGroup := [3]int{0, 0, 0}

	for i, in := range input {
		num := Utils.StringToInt(in)
		fGroup = addNewNumber(fGroup, num)
		sGroup = addNewNumber(sGroup, fGroup[1])
		//fmt.Printf("f: %v\ts: %v\n", fGroup, sGroup)
		if i > 2 {
			if getSum(fGroup) - getSum(sGroup) > 0 { count ++ }
		}
	}

	fmt.Printf("Part 2: %d\n", count)
}

//Shuffles new number in to front of array
//@param numbers -- array to add new number in
//@param newNum -- number to add at front
//@return new array
func addNewNumber(numbers [3]int, newNum int) [3]int {
	numbers[2] = numbers[1]
	numbers[1] = numbers[0]
	numbers[0] = newNum

	return numbers
}

//Calcs sum of array
//@param numbers -- array to sum
//@return sum of array
func getSum(numbers [3]int) int {
	return numbers[0] + numbers [1] + numbers[2]
}