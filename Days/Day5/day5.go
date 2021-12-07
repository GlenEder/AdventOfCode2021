package Day5

import (
	"fmt"
	"src/Utils"
	"strings"
)


func Run() {
	Utils.PrintDay(5)

	//get input
	input := Utils.ReadInputToSlice("Days/Day5/input.txt")
	part1Overlaps := 0
	part2Overlaps := 0

	var part1map = make(map[Utils.Point]int)
	var part2map = make(map[Utils.Point]int)

	//create lines
	for _, l := range input {
		//parse input line
		points := strings.Fields(l)
		firstPoint := strings.Split(points[0], ",")
		secPoint := strings.Split(points[2], ",")

		p1 := Utils.Point{
			X: Utils.StringToInt(firstPoint[0]),
			Y: Utils.StringToInt(firstPoint[1]),
		}
		p2 := Utils.Point{
			X: Utils.StringToInt(secPoint[0]),
			Y: Utils.StringToInt(secPoint[1]),
		}

		var newLine Utils.Line
		if p1.X < p2.X {
			newLine = Utils.CreateLineWithPoints(p1, p2)
		} else {
			newLine = Utils.CreateLineWithPoints(p2, p1)
		}


		//fmt.Printf("Created line: %s\n", newLine)
		//Check for overlaps in part 1
		if newLine.IsVert {
			start := newLine.Start.Y
			end := newLine.End.Y
			if start > end {
				start = newLine.End.Y
				end = newLine.Start.Y
			}
			for start <= end {
				p := Utils.Point{X: newLine.Start.X, Y: start}
				pointCount, _ := part1map[p]
				part1map[p] = pointCount + 1
				pc, _ := part2map[p]
				part2map[p] = pc + 1
				start++
			}

		} else if newLine.IsHorz {
			start := newLine.Start.X
			end := newLine.End.X
			for start <= end {
				p := Utils.Point{X: start, Y: newLine.Start.Y}
				pointCount, _ := part1map[p]
				part1map[p] = pointCount + 1
				pc, _ := part2map[p]
				part2map[p] = pc + 1
				start++
			}
		} else {
			//fmt.Printf("Line: %s\n", newLine)
			toAdd := Utils.Point{
				X: newLine.Slope.Numerator,
				Y: newLine.Slope.Denominator,
			}

			start := newLine.Start
			for start != newLine.End {
				//fmt.Printf("\tAdding point: %s\n", start)
				pointCount, _ := part2map[start]
				part2map[start] = pointCount + 1
				start = start.AddPoint(toAdd)
			}
			//add end point
			//fmt.Printf("\tAdding %s to map\n", start)
			pc, _ := part2map[start]
			part2map[start] = pc + 1
		}

	}

	//Check where there are more than 1 overlap
	for _, c := range part1map {
		if c > 1 {
			part1Overlaps++
		}
	}

	for _, c := range part2map {
		if c > 1 {
			part2Overlaps++
		}
	}

	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1Overlaps, part2Overlaps)
}
