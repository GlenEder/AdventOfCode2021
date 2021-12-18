package Utils

import "fmt"

//2d point struct
type Point struct {
	X int
	Y int
}

//Adds two points together
//@param p2 -- point to add to point object
//@return point with added data
func (p Point)AddPoint(p2 Point) Point {
	return Point{
		X: p.X + p2.X,
		Y: p.Y + p2.Y,
	}
}

//Adds x and y values to point
//@param p2 -- point to add to point object
//@return point with added data
func (p Point)Add(x int, y int) Point {
	return Point{
		X: p.X + x,
		Y: p.Y + y,
	}
}


//To string method for point struct
//@return string representation of point
func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

//Checks if the point is in the point slice
//@param s -- point slice to check in
//@return true if in, otherwise false
func (p Point) InPointSlice(s []Point) bool {
	for _, point := range s {
		if p == point { return true }
	}
	return false
}

//Checks if point lies within the bounds (inclusive)
//@param xMin -- min x value
//@param xMax -- max x value
//@param yMin -- min y value
//@param yMax -- max y value
//@return true if in bounds
func (p Point) InBounds(xMin int, xMax int, yMin int, yMax int) bool {
	if p.X < xMin || p.X > xMax || p.Y < yMin || p.Y > yMax {
		return false
	}
	return true
}

//Checks if point lies within the bounds of (0,0) to (x, y)
//@param x -- max x value
//@param y -- max y value
//@return true if in bounds
func (p Point) InPositiveBounds(x int, y int) bool {
	return p.InBounds(0, x, 0, y)
}

//Calculates the manhattan distance between the point and the x,y cord provided
//@param x -- x value of end point
//@param y -- y value of end point
//@return manhattan distance
func (p Point) ManhattanDistance(x int, y int) int {
	return AbsInt(p.X - x) + AbsInt(p.Y - y)
}

//Calculates the manhattan distance between the point and the x,y cord provided
//@param p2 -- end point
//@return manhattan distance
func (p Point) ManhattanDistanceFromPoint(p2 Point) int {
	return p.ManhattanDistance(p2.X, p2.Y)
}
