// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import (
	"math"
)
// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind string

const (
    // Pick values for the following identifiers used by the test program.
    NaT = Kind("NaT") // not a triangle
    Equ = Kind("Equ") // equilateral
    Iso = Kind("Iso") // isosceles
    Sca = Kind("Sca") // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	if isValidTriangle(a, b, c) {
		if a == b && b == c {
			return Equ
		} else if a == b || a == c || b == c {
			return Iso
		} else if a != b && b != c {
			return Sca
		}
	}
	return NaT
}

func isValidTriangle(a, b, c float64) bool {
	if !validateSide(a) || !validateSide(b) || !validateSide(c) {
		return false
	} else if a+b < c || a+c < b || c+b < a || a <= 0 || b <= 0 || c <= 0 {
			return false
		} else {
			return true
		}
}

func validateSide(side float64) bool {
	return !math.IsNaN(side) && !math.IsInf(side, 1) && side > 0
}
