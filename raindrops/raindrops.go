package raindrops

import (
    "strconv"
)

func Convert(number int) (converted string) {
	if number % 3 == 0 {
    	converted += "Pling"
    }
	if number % 5 == 0 {
    	converted += "Plang"
	}
    if number % 7 == 0 {
    	converted += "Plong"
    }
    if converted == "" {
		converted = strconv.Itoa(number)
    }
	return
}
