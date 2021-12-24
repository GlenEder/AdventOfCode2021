package Utils

import "fmt"

type Range struct {
	Upper int
	Lower int
}

//Checks if number is in range
//@param i -- int to check
//@return true if in range
func (r *Range) InRange(i int) bool {
	return i >= r.Lower && i <= r.Upper
}

//To string override
//@return string rep of range
func (r Range) String() string {
	return fmt.Sprintf("{%d --> %d}", r.Lower, r.Upper)
}
