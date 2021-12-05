package Utils

import "fmt"

//Line struct that holds the two integer points
type Line struct {
	Start Point
	End Point
	Slope Fraction
	B int
	IsVert bool
	IsHorz bool
}

func CreateLine(x1 int, y1 int, x2 int, y2 int) Line {
	newLine := Line{
		Start: Point{
			X: x1,
			Y: y1,
		},
		End: Point{
			X: x2,
			Y: y2,
		},
		Slope: Fraction{0,0},
		B: 0,
		IsVert: x1 == x2,
		IsHorz: y1 == y2,
	}
	//Calculate line equation with two points
	newLine.SetLineEquation()

	return newLine
}

func CreateLineWithPoints(p1 Point, p2 Point) Line {
	return CreateLine(p1.X, p1.Y, p2.X, p2.Y)
}

//Calculates line equation (y = mx + b) based on start and end points
func (l *Line)SetLineEquation() {
	l.Slope = l.CalculateSlope()
	l.B = l.CalculateYIntercept()
}

//Calculates the y intercept of the line
//@return y intercept
func (l Line)CalculateYIntercept() int {
	return l.Start.Y - (l.Slope.ToInt() * l.Start.X)
}

//Returns the slope filed of the line based on its two points
//@return slope in reduced fraction form
func (l Line)CalculateSlope() Fraction {
	//calc slopes
	deltaX := l.End.X - l.Start.X
	deltaY := l.End.Y - l.Start.Y
	f := Fraction{
		Numerator: deltaX,
		Denominator: deltaY,
	}
	f = f.Reduce()
	return f
}

//Checks to see if the point provided is on the line
//@param p -- point to check
//@return true if on line, false otherwise
func (l Line) PointIsOnLine(p Point) bool {
	return p.Y == l.B + (l.Slope.ToInt() * p.X)
}

//Checks to see if the point is on the line segment of the two points
//@param p -- point to check
//@return true if on line, false otherwise
func (l Line) PointIsOnLineSegment(p Point) bool {
	ab := ((l.End.X - l.Start.X) * (l.End.X - l.Start.X)) + ((l.End.Y - l.Start.Y) * (l.End.Y - l.Start.Y))
	ap := ((p.X - l.Start.X) * (p.X - l.Start.X)) + ((p.Y - l.Start.Y) * (p.Y - l.Start.Y))
	pb := ((l.End.X - p.X) * (l.End.X - p.X)) + ((l.End.Y - p.Y) * (l.End.Y - p.Y))
	return ab == ap + pb
}

//Returns a formatted string of line data
func (l Line) String() string {
	return fmt.Sprintf("%d,%d -> %d,%d\ty = %dx + %d\tm=%d/%d",
		l.Start.X, l.Start.Y, l.End.X, l.End.Y, l.Slope.ToInt(), l.B, l.Slope.Numerator, l.Slope.Denominator)
}