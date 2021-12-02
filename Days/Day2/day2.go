package Day2

import (
	"fmt"
	"src/Utils"
	"strings"
)

func Run() {
	fmt.Printf("===Day 2===\n")
	//read input file
	input := Utils.ReadInputToSlice("Days/Day2/input.txt")
	x, y := 0, 0

	for _, line := range input {
		splits := strings.Split(line, " ")
		i := Utils.StringToInt(splits[1])
		switch line[0] {
		case 'u':
			y -= i
			break
		case 'd':
			y += i
			break
		case 'f':
			x += i
			break
		default:
			panic(fmt.Errorf("Input not expected: %s\n", line))
		}
	}

	fmt.Printf("Part 1: %d\n", x * y)

	//Reset for part two
	x, y = 0, 0
	aim := 0

	for _, line := range input {
		splits := strings.Split(line, " ")
		i := Utils.StringToInt(splits[1])
		switch line[0] {
		case 'u':
			aim -= i
			break
		case 'd':
			aim += i
			break
		case 'f':
			x += i
			y += aim * i
			break
		default:
			panic(fmt.Errorf("Input not expected: %s\n", line))
		}
	}

	fmt.Printf("Part 2: %d\n", y * x)
}