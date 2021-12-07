package Day6

import (
	"fmt"
	"src/Utils"
	"strings"
)


func Run() {
	Utils.PrintDay(6)

	//get input
	input := Utils.ReadInputToSlice("Days/Day6/input.txt")

	maxLife := 8
	part1Days := 80
	part2Days := 256


	//slice of lanternfish lifespans
	var lanternfish []int
	for i := 0; i <= maxLife; i++ {
		lanternfish = append(lanternfish, 0)
	}

	//load input to lanternfish array
	numbers := strings.Split(input[0], ",")
	for _, num := range numbers {
		lanternfish[Utils.StringToInt(num)]++
	}

	//Day sims
	for i := 0; i < part2Days; i++ {
		if i == part1Days {
			fmt.Printf("Part 1: %d\n", calcTotal(lanternfish))
		}
		simDay(lanternfish)
	}

	fmt.Printf("Part 2: %d\n", calcTotal(lanternfish))
}

//Sims the day by shuffling down the numbers in the tracker
//@param lanternfish -- slice of fish counts to simulate
func simDay(lanternfish []int) {
	//fmt.Printf("s: %v\t", lanternfish)
	numRespawns := lanternfish[0]
	for i := 1; i < len(lanternfish); i++ {
		lanternfish[i - 1] = lanternfish[i]
	}
	//Add new lantern fish
	lanternfish[8] = numRespawns
	//Respawn dying lantern fish
	lanternfish[6] += numRespawns

	//fmt.Printf("e: %v\n", lanternfish)
}

//calcTotal sums the fishes in the slice
//@param lanternfish -- slice of fish counts
func calcTotal(lanternfish []int) int {
	toReturn := 0
	for _, fish := range lanternfish {
		toReturn += fish
	}
	return toReturn
}