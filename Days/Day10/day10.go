package Day10

import (
	"fmt"
	"sort"
	"src/Utils"
	"strings"
)


func Run() {
	Utils.PrintDay(10)

	//get input
	input := Utils.ReadInputToSlice("Days/Day10/input.txt")

	openings := []string{"(", "{", "[", "<"}
	closers := []string{")", "}", "]", ">"}

	part1total := 0
	var part2scores []int

	for _, line := range input {
		var stack Utils.StringStack
		corrupt := false
		for _, char := range strings.Split(line, "") {
			if Utils.SliceContainsString(openings, char) {
				stack.Push(char)
			} else {
				firstToClose, err := stack.Pop()
				//If stack empty have syntax error
				if err != nil {
					corrupt = true
					part1total += part1Calc(char)
				} else {
					indexOfClose := Utils.SliceIndexOfStr(closers, char)
					indexOfOpen := Utils.SliceIndexOfStr(openings, firstToClose)
					if indexOfOpen != indexOfClose {
						part1total += part1Calc(char)
						corrupt = true
					}
				}
			}
		}

		//Part 2 -- complete the line
		if !corrupt {
			score := 0
			for !stack.IsEmpty() {
				toClose, _ := stack.Pop()
				indexOfOpen := Utils.SliceIndexOfStr(openings, toClose)
				score = part2Calc(closers[indexOfOpen], score)
			}
			part2scores = append(part2scores, score)
		}
	}

	//sort scores
	sort.Ints(part2scores)

	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1total, part2scores[len(part2scores) / 2])

}

func part1Calc(s string) int {
	switch s {
	case ")": return 3
	case "]": return 57
	case "}": return 1197
	case ">": return 25137
	default:
		return -1
	}
}

func part2Calc(s string, currScore int) int {
	currScore *= 5
	switch s {
	case ")":
		currScore += 1
		break
	case "]":
		currScore += 2
		break
	case "}":
		currScore += 3
		break
	case ">":
		currScore += 4
		break
	}
	return currScore
}