package hamming

import (
	"errors"
)

// Distance function first determines if two strings are the same length and
// if true, calculates the differences for each character in the string.
func Distance(a, b string) (distance int, err error) {
	if len(a) != len(b) {
		err = errors.New("strings are of unequal length")
        return
	}
	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}
	return
}
