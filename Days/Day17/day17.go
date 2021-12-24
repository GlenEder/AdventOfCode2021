package Day17

import (
	"fmt"
	"regexp"
	"src/Utils"
)

func Run() {
	Utils.PrintDay(17)

	//get input
	input := Utils.ReadInputAsString("Days/Day17/input.txt")

	//find x and y's of target area
	expression := regexp.MustCompile(`[-]?\d[\d]*[\]?[\d{2}]*`)
	numbers := expression.FindAllString(input, 4)

	xMin := Utils.StringToInt(numbers[0])
	xMax := Utils.StringToInt(numbers[1])
	yMin := Utils.StringToInt(numbers[2])
	yMax := Utils.StringToInt(numbers[3])

	targetZone := Utils.CreateRectangle(xMin, yMax, xMax, yMin)
	fmt.Printf("Target: %s\n", targetZone)

	//Part 1 calculation
	n := (-1 * yMin) - 1
	part1 := (n * (n + 1)) / 2

	//Part 2 good old brute force
	part2 := make(map[Utils.Point]int)
	yRange := targetZone.GetHeightRange()
	xRange := targetZone.GetWidthRange()
	for i := -Utils.AbsInt(yMin); i <= Utils.AbsInt(yMin); i++ {
		ySteps := calcYSteps(i, yRange)
		//check for matching x values
		for _, velSteps := range ySteps {
			xMatches := matchingX(velSteps.X, xRange)

			//add start velocities to map
			for _, x := range xMatches {
				p := Utils.Point{X: x, Y: velSteps.Y}
				//fmt.Printf("Match found: %s\n", p)
				part2[p] = 1
			}
		}
	}

	// fmt.Printf("%v\n", part2)
	//Print results
	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1, len(part2))
}

//Calcs number of ySteps needed to land in target range
//@param yVel -- y veloctiy of projectile
//@param target -- y range of target zone
//@return slice of num steps points, yVelStart
func calcYSteps(yVel int, target Utils.Range) []Utils.Point {
	//Start y positon
	yPos := 0
	vel := yVel
	steps := 0
	possibleSteps := []Utils.Point{}
	for {
		//check if in range
		if target.InRange(yPos) {
			// Utils.PrintWithColorf(Utils.Green, "\t\tMatch found @ %d in %d steps with vel %d\n", yPos, steps, yVel)
			possibleSteps = append(possibleSteps, Utils.Point{X: steps, Y: yVel})
		}

		//Check if past lower
		if yPos < target.Lower {
			break
		}

		//move projectile
		yPos += vel
		//decrese velocity by 1
		vel--
		//increase step counter
		steps++
	}

	// 	fmt.Printf("Found %d poss steps for vel %d\n", len(possibleSteps), yVel)
	return possibleSteps
}

func matchingX(numSteps int, target Utils.Range) []int {

	matches := []int{}
	for xVelMax := Utils.AbsInt(target.Upper); xVelMax > 0; xVelMax-- {
		pos := 0
		xVel := xVelMax
		for i := 0; i < numSteps; i++ {
			pos += xVel
			xVel--
			if pos > target.Upper || xVel == 0 {
				break
			}
		}

		if target.InRange(pos) {
			//Utils.PrintWithColorf(Utils.Green, "\t\tMatch found @ vel %d\n", xVelMax)
			matches = append(matches, xVelMax)
		}
	}

	return matches
}
