package Day15

import (
	"fmt"
	"src/Utils"
)


func Run() {
	Utils.PrintDay(15)

	//get input
	input := Utils.ReadInputToIntGrid("Days/Day15/input.txt", "")

	part1total := 0
	start := Utils.Point{X: 0, Y: 0}
	end := Utils.Point{X: len(input[0]) - 1, Y: len(input) - 1}
	path1 := Utils.AStar(input, start, end)

	//Utils.PrintGridWithHighlights(input, path)
	//Calc part 1
	for _, point := range path1 {
		part1total += input[point.Y][point.X]
	}
	part1total -= input[0][0]

	//creat larger grid for part 2
	scaling := 5
	var newInput [][]int
	for j := 0; j < scaling; j++ {
		for _, line := range input {
			var inputline []int
			for i := 0; i < scaling; i++ {
				for _, num := range line {
					newNum := num + i + j
					for newNum > 9 { newNum -= 9 }
					inputline = append(inputline, newNum)
				}
			}
			newInput = append(newInput, inputline)
		}
	}

	end2 := Utils.Point{X: len(newInput[0]) - 1, Y: len(newInput) - 1}
	path2 := Utils.AStar(newInput, start, end2)
	//Utils.PrintGridWithHighlights(newInput, path2)
	part2total := 0
	//Calc part 2
	for _, point := range path2 {
		part2total += newInput[point.Y][point.X]
	}
	part2total -= newInput[0][0]

	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1total, part2total)
}
