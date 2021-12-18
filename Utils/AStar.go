package Utils

import "math"

func AStar(grid [][]int, start Point, end Point) []Point {

	openSet := make(map[Point]int)
	closedSet := make(map[Point]int)
	maxXandY := Point{X: len(grid[0]), Y: len(grid)}
	gScores := make(map[Point]int)
	//fScores := make(map[Point]int)
	cameFrom := make(map[Point]Point)


	//setup starting point data
	openSet[start] = 1
	gScores[start] = grid[start.Y][start.X]
	//[start] = start.ManhattanDistanceFromPoint(end) + gScores[start]

	//A star loop
	for len(openSet) > 0 {
		//find best fit for next point
		next := bestF(grid, gScores, openSet, end)
		//fmt.Printf("Next point: %s\n", next)

		//Check for done
		if next == end {
			//PrintWithColorf(Green, "A Star done.\n")
			return recreatePath(cameFrom, next)
		}
		//Remove next point from open list
		delete(openSet, next)
		//Add current to closed set
		closedSet[next] = 1
		//Check neighbors of next and add to openSet
		neighbors := getNeighbors(next, maxXandY)
		//calc neighbors g scores
		for _, neighbor := range neighbors {
			if _, exists := closedSet[neighbor]; exists { continue }
			tempGscore := gScores[next] + grid[next.Y][next.X]
			if _, exists := openSet[neighbor]; exists {
				if tempGscore < gScores[neighbor] {
					gScores[neighbor] = tempGscore
					cameFrom[neighbor] = next
				}
			} else {
				gScores[neighbor] = tempGscore
				cameFrom[neighbor] = next
				openSet[neighbor] = 1
			}
		}
	}

	//No solution
	PrintError("AStar path not found.\n")
	return []Point{}
}

//Returns slice of points for all neighbors around point
//@param home -- point to look around
//@param maxXY -- point struct holding the max x and y values
//@return slice of neighbor points
func getNeighbors(home Point, maxXY Point) []Point {
	var neighbors []Point
	if home.X - 1 >= 0 {
		left := home.Add(-1, 0)
		neighbors = append(neighbors, left)
	}
	if home.X + 1 < maxXY.X {
		right := home.Add(1, 0)
		neighbors = append(neighbors, right)
	}
	if home.Y - 1 >= 0 {
		up := home.Add(0, -1)
		neighbors = append(neighbors, up)
	}
	if home.Y + 1 < maxXY.Y {
		down := home.Add(0, 1)
		neighbors = append(neighbors, down)
	}

	return neighbors
}

//Walks back up the path creating a point slice to return
//@param cameFrom -- map of points and where they came from
//@param current -- last point in path
//@return path in reverse order
func recreatePath(cameFrom map[Point]Point, current Point) []Point {
	path := []Point{current}

	next, exists := cameFrom[current]
	for exists {
		//fmt.Printf("Adding %s to path\n", next)
		path = append([]Point{next}, path...)
		next, exists = cameFrom[next]
	}

	return path
}

//Calculates the best point in the open set to go to next
//@param grid -- grid to navigate
//@param openSet -- current set of points to evaluate
//@param end -- end point
//@return best point to go to
func bestF(grid [][]int, gScores map[Point]int, openSet map[Point]int, end Point) Point {
	bestPoint := Point{}
	bestF := math.MaxInt32
 	for p, _ := range openSet {
 		f := p.ManhattanDistanceFromPoint(end) + grid[p.Y][p.X]
		if score, exists := gScores[p]; exists {
			f += score
		}
		if f < bestF {
			bestF = f
			bestPoint = p
		}
	}

	return bestPoint
}