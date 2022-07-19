package thefarm

import (
    "math"
    "fmt"
    "errors"
)
// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
    negativeCows int
}

func (e *SillyNephewError) Error() string {
    return fmt.Sprintf("silly nephew, there cannot be %d cows", e.negativeCows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodder, err := weightFodder.FodderAmount()
    switch {
        case cows == 0:
    	return 0.0, errors.New("division by zero")
        case cows < 0:
    	return 0.0, &SillyNephewError{negativeCows: cows}
        case err != ErrScaleMalfunction && err != nil:
    	return 0.0, err
        case math.Signbit(fodder):
    	return 0.0, errors.New("negative fodder")
        case err == ErrScaleMalfunction && !math.Signbit(fodder):
		fodder *= 2.0
    }
	return fodder / float64(cows), nil
}
