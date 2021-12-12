package Day9

import (
	"fmt"
	"sort"
	"src/Utils"
)

type basinValue struct {
	basinId int
	fillLevel int
	filled bool
}

func Run() {
	Utils.PrintDay(9)

	part1total := 0
	//get input
	numInput := Utils.ReadInputToIntGrid("Days/Day9/input.txt", "")

	//Part 2 setup
	var visited [][]int
	for y, yline := range numInput{
		visited = append(visited, []int{})
		for _ = range yline {
			visited[y] = append(visited[y], 0)
		}
	}
	basinIds := make(map[int]int)
	currId := 1

	for y, yLine := range numInput {
		for x, num := range yLine {
			islowest := true

			if up := y - 1; up >= 0 {
				upNum := numInput[up][x]
				if upNum <= num { islowest = false }
			}
			if down := y + 1; down < len(numInput) {
				if numInput[down][x] <= num { islowest = false }
			}
			if left := x - 1; left >= 0 {
				leftNum := yLine[left]
				if leftNum <= num { islowest = false }
			}
			if right := x + 1; right < len(yLine) {
				rightNum := yLine[right]
				if rightNum <= num { islowest = false }
			}
			if islowest { part1total += num + 1 }

			//Part 2
			point := Utils.Point{X: x, Y: y}
			if num != 9 && visited[y][x] == 0 {
				basinIds[currId] = flood(point, numInput, visited)
				currId++
			}
		}
	}

	var maxes []int
	for _, num := range basinIds {
		maxes = append(maxes, num)
		sort.Ints(maxes)
		if len(maxes) > 3 {
			maxes = maxes[1:4]
		}
	}

	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1total, Utils.ProductOfIntSlice(maxes))
}

func flood(start Utils.Point, inputMap [][]int, visited [][]int) int {
	if start.Y < 0 || start.Y >= len(inputMap) { return 0 }
	if start.X < 0 || start.X >= len(inputMap[start.Y]) { return 0 }
	if inputMap[start.Y][start.X] == 9 {
		visited[start.Y][start.X] = 8
		return 0
	}
	if visited[start.Y][start.X] > 0 { return 0 }
	visited[start.Y][start.X] = 1
	size := 1
	size += flood(start.Add(0, 1), inputMap, visited)
	size += flood(start.Add(0, -1), inputMap, visited)
	size += flood(start.Add(1, 0), inputMap, visited)
	size += flood(start.Add(-1, 0), inputMap, visited)
	return size
}







