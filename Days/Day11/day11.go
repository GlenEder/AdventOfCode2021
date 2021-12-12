package Day11

import (
	"fmt"
	"src/Utils"
)


func Run() {
	Utils.PrintDay(11)

	//get input
	input := Utils.ReadInputToIntGrid("Days/Day11/input.txt", "")


	part1total := 0
	part1Steps := 100
	currStep := 0
	for true {
		currStep++
		explosions := 0
		for y, line := range input {
			for x, eLevel := range line {
				input[y][x]++
				if eLevel == 9 {
					explosions += explode(Utils.Point{X: x, Y: y}, input)
				}
			}
		}

		//Set all numbers over 9 to 0
		for y, line := range input {
			for x, num := range line {
				if num > 9 {
					input[y][x] = 0
				}
			}
		}
		//fmt.Printf("Num Explosions Expericenced: %d\n", part1total)
		//Utils.PrintIntGrid(input
		if currStep <= part1Steps {
			part1total += explosions
		}
		if explosions == len(input) * len(input[0]) {
			fmt.Printf("Part 1: %d\nPart 2: %d\n", part1total, currStep)
			return
		}
	}

}

func explode(start Utils.Point, grid [][]int) int {

	x := start.X
	y := start.Y
	yMax := len(grid) - 1
	xMax := len(grid[0]) - 1
	explosions := 1
	//increment surrounding points
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			point := Utils.Point{X: x + j, Y: y + i}
			if point.InPositiveBounds(xMax, yMax) {
					if i == 0 && j == 0 { continue }
					grid[y + i][x + j]++
					if grid[y + i][x + j] == 10 {
						explosions += explode(Utils.Point{X:x + j, Y: y + i}, grid)
					}
			}
		}
	}
	return explosions
}

