package Utils

import "fmt"

//Struct to hold point stack data
type PointStack struct {
	stack []Point
}

//Returns the point stack's slice
//@return point stack's slice
func (ps *PointStack) GetStackSlice() []Point {
	return ps.stack
}

//Pushes new point onto stack
//@param p -- point to add
func (ps *PointStack) Push(p Point) {
	ps.stack = append([]Point{p}, ps.stack...)
}

//Pops the first point off the stack
//@return first point on stack, error if stack is empty
func (ps *PointStack) Pop() (Point, error) {
	if len(ps.stack) < 1 { return Point{}, fmt.Errorf("cannot pop from empty stack")}
	toReturn := ps.stack[0]
	ps.stack = ps.stack[1:]
	return toReturn, nil
}

//Removes the point from the stack
//@param p -- point to remove
func (ps *PointStack) Remove(p Point) {
	for i, point := range ps.stack {
		if p == point {
			ps.stack = append(ps.stack[:i], ps.stack[i+1:]...)
		}
	}
}

//Checks if stack is empty
//@return true if stack is empty, otherwise false
func (ps *PointStack) IsEmpty() bool {
	return len(ps.stack) < 1
}

//Checks if stack contains the point
//@param p -- point to check for
//@return true if stack has point, otherwise false
func (ps *PointStack) Contains(p Point) bool {
	for _, point := range ps.stack {
		if point == p {
			return true
		}
	}
	return false
}
