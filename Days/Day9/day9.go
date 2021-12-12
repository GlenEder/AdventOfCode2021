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
	numInput := Utils.ReadInputToIntGrid("Days/Day9/testInput.txt", "")

	//Part 2
	var maxes []int
	var basinIds Utils.IntStack
	basin := make(map[Utils.Point]int)
	basinFillCount := make(map[int]int)

	for y, yLine := range numInput {
		for x, num := range yLine {
			islowest := true
			point := Utils.Point{X: x, Y: y}
			basinId := -1
			if up := y - 1; up >= 0 {
				upNum := numInput[up][x]
				if upNum <= num { islowest = false }
				if upNum != 9 {
					if bVal, exists := basin[point.Add(0, -1)]; exists {
						basinId = bVal
					}
				}
			}
			if down := y + 1; down < len(numInput) {
				if numInput[down][x] <= num { islowest = false }
			}
			if left := x - 1; left >= 0 {
				leftNum := yLine[left]
				if leftNum <= num { islowest = false }
				if leftNum != 9 {
					if bVal, exists := basin[point.Add(-1, 0)]; exists {
						basinId = bVal
					}
				}

			}
			if right := x + 1; right < len(yLine) {
				rightNum := yLine[right]
				if rightNum <= num { islowest = false }
				for right < len(yLine) {
					if yLine[right] != 9 {
						if bVal, exists := basin[Utils.Point{X: right, Y: y - 1}]; exists {
							basinId = bVal
							break
						}
					} else { break }
					right++
				}
			}
			if islowest { part1total += num + 1 }

			//Part 2
			//create new basinId if needed
			if num == 9 { continue }
			if basinId == -1 {
				max, err := basinIds.Max()
				if err == nil { basinId = max + 1 } else { basinId = 1}
				basinIds.Push(basinId)
			}
			//add point to map
			basin[point] = basinId
			//add count to basin fill count
			prevVal, _ := basinFillCount[basinId]
			basinFillCount[basinId] = prevVal + 1
			//fmt.Printf("Point: %s\tBasin Id: %d\tFill Level: %d\n", point,  basinId, basinFillCount[basinId])
		}
	}

	for _, fillVal := range basinFillCount {
		//fmt.Printf("Attempting to add %d...", fillVal)
		maxes = append(maxes, fillVal)
		sort.Ints(maxes)
		if len(maxes) > 3 {
			maxes = maxes[1:len(maxes)]
		}
		//fmt.Printf("\t%v\n", maxes)
	}

	product := 1
	for _, val := range maxes { product *= val }

	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1total, product)
}






