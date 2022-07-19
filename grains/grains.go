package grains

import (
    "errors"
)

func Square(number int) (uint64, error) {
    if number <= 0 || number > 64 {
        return 0, errors.New("number must be an integer between 1 and 64")
    }
    return uint64(1) << uint(number-1), nil
}

func Total() (grainTotal uint64) {
	for b := 64; b > 0; b-- {
        s, err := Square(b)
        if err != nil {
            panic(err)
        }
    	grainTotal += s
    }
	return
}
