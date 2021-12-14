package Day13

import (
	"fmt"
	"src/Utils"
	"strings"
)


func Run() {
	Utils.PrintDay(13)

	//get input
	input := Utils.ReadInputToSlice("Days/Day13/input.txt")

	var folds []string
	board := make(map[Utils.Point]int)

	//create map of points
	for _, line := range input {
		point, err := createPoint(line)
		//fold command
		if err != nil {
			if len(line) > 1 {
				folds = append(folds, line)
			}
		} else {
			//add point to map
			board[point] = 1
		}
	}

	//fmt.Printf("Fold commands: %v\n", folds)
	lastXFold := 0
	lastYFold := 0
	//Do the folding
	for i, fold := range folds {
		if strings.Contains(fold, "x") {
			//Fold left
			x := Utils.StringToInt(strings.Split(fold, "=")[1])
			lastXFold = x
			for point, _ := range board {
				if point.X > x {
					dif := point.X - x
					newPoint := point.Add(dif * -2, 0)
					board[newPoint] = 1
					delete(board, point)
				}
			}
		} else {
			//Fold up
			y :=  Utils.StringToInt(strings.Split(fold, "=")[1])
			lastYFold = y
			for point, _ := range board {
				if point.Y > y {
					dif := point.Y - y
					newPoint := point.Add(0, dif * -2)
					board[newPoint] = 1
					delete(board, point)
				}
			}
		}

		if i == 0 {
			fmt.Printf("Part 1: %d\n", len(board))
		}
	}

	//print part 2's code
	fmt.Printf("Part 2:\n")
	for y := 0; y < lastYFold; y++ {
		for x := 0; x < lastXFold; x++ {
			if _, exists := board[Utils.Point{X: x, Y: y}]; exists {
				fmt.Printf("#")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}


func createPoint(s string) (Utils.Point, error){
	parts := strings.Split(s, ",")
	if len(parts) == 1 {
		return Utils.Point{}, fmt.Errorf("string not in point format")
	}

	return Utils.Point{
		X: Utils.StringToInt(parts[0]),
		Y: Utils.StringToInt(parts[1]),
	}, nil
}