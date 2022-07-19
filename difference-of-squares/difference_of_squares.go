package diffsquares

import (
    "math"
)

func SquareOfSum(n int) (sqsum int) {
	for n > 0{
        sqsum+=n
        n--
    }
	sqsum = int(math.Pow(float64(sqsum), 2.0))
    return
}

func SumOfSquares(n int) (sumsq int) {
	for n > 0 {
        sumsq += int(math.Pow(float64(n), 2.0))
        n--
    }
	return
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
