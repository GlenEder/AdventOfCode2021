package Utils

import "fmt"

type Rectangle struct {
	topLeft     Point
	bottomRight Point
}

//Creates a rectangle based off the points provided
//@param x1 -- left most value of rectangle
//@param y1 -- top most value of rectangle
//@param x2 -- right most value of rectangle
//@praam y2 -- bottom most value of rectangle
//@return Rectangle object with points provided
func CreateRectangle(x1 int, y1 int, x2 int, y2 int) Rectangle {
	return Rectangle{
		topLeft: Point{
			X: x1,
			Y: y1,
		},
		bottomRight: Point{
			X: x2,
			Y: y2,
		},
	}
}

//To string override for rectangle
//@return string reep of rectangle
func (rect Rectangle) String() string {
	return fmt.Sprintf("{(%d, %d), (%d, %d)}", rect.topLeft.X, rect.topLeft.Y, rect.bottomRight.X, rect.bottomRight.Y)
}

//Gets the rectangles max y value
//@return max y value
func (rect *Rectangle) GetTop() int {
	return rect.topLeft.Y
}

//Gets the rectangles min y value
//@return min y value
func (rect *Rectangle) GetBottom() int {
	return rect.bottomRight.Y
}

//Gets the rectangles min x value
//@return min x value
func (rect *Rectangle) GetLeft() int {
	return rect.topLeft.X
}

//Gets the rectangles max x value
//@return max x value
func (rect *Rectangle) GetRight() int {
	return rect.bottomRight.X
}

//Returns a range of the rectangles y values
//@return Range of y values
func (rect *Rectangle) GetHeightRange() Range {
	return Range{
		Upper: rect.topLeft.Y,
		Lower: rect.bottomRight.Y,
	}
}

//Returns a range of the rectangles x values
//@return Range of x values
func (rect *Rectangle) GetWidthRange() Range {
	return Range{
		Upper: rect.bottomRight.X,
		Lower: rect.topLeft.X,
	}
}

//Checks if the point is within the bounds of the rectangle
//@param p -- point to check
//@return true if point lies within bounds, otherwise false
func (rect *Rectangle) PointInBounds(p Point) bool {
	return p.X >= rect.topLeft.X &&
		p.X <= rect.bottomRight.X &&
		p.Y >= rect.bottomRight.Y &&
		p.Y <= rect.topLeft.Y
}
