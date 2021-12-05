package Utils

import (
	"fmt"
	"math"
)


// Fraction represents ratio.
type Fraction struct {
	Numerator int
	Denominator int
}

//Returns string rep of fraction
func (f Fraction) String() string {
	return fmt.Sprintf("%d/%d", f.Numerator, f.Denominator)
}

//calculates the given fraction's reduced form.
//@return reduced fraction
func (f Fraction) Reduce() Fraction {
	gcd := f.CalcGCD()
	f.Numerator /= gcd
	f.Denominator /= gcd
	return f
}

//coverts the fraction to an int
//@return int form of fraction
func (f Fraction) ToInt() int {
	if f.Denominator == 0 {
		return 0
	}
	return f.Numerator / f.Denominator
}

//converts the fraction to a float64
//@return float form of fraction
func (f Fraction) ToFloat() float64 {
	return float64((f.Numerator * 1.0) / f.Denominator)
}

//Calculates greatest common denominator
//@return gcd of fraction
func (f Fraction) CalcGCD() int {
	a := int(math.Abs(float64(f.Numerator)))
	b := int(math.Abs(float64(f.Denominator)))
	r := 0
	if b == 0 { return a }
	for a % b > 0 {
		r = a % b
		a = b
		b = r
	}
	return b
}
