package Day7

import (
	"fmt"
	"sort"
	"src/Utils"
	"strings"
)


func Run() {
	Utils.PrintDay(7)

	//get input
	input := Utils.ReadInputAsString("Days/Day7/input.txt")
	//convert to int slice
	numbers := Utils.ToIntSlice(strings.Split(input, ","))
	//sort list
	sort.Ints(numbers)
	//find median for part 1
	median := Utils.Median(numbers)
	part1 := 0
	//find mean for part 2
	mean := Utils.LowMean(numbers)
	part2 := 0

	for _, num := range numbers {
		part1 += Utils.AbsInt(num - median)
		part2 += part2FuelCost(Utils.AbsInt(num - mean))
	}

	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1, part2)
}

//Calcs the linear fuel cost for part 2
//@param dist -- distance crab must travel
//@return total fuel cost 
func part2FuelCost(dist int) int {
	cost := 0
	for i := 1; i <= dist; i++ {
		cost += i
	}
	return cost
}