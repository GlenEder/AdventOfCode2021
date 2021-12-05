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

