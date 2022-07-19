// romannumerals package converts arabic numbers to roman numerals
package romannumerals

import (
    "errors"
    )

// ToRomanNumeral converts an integer to a roman numeral formatted string 
func ToRomanNumeral(input int) (string, error) {
	var romanMap = map[string]int {
        "M":	1000,
        "CM":	900,
        "D": 	500,
        "CD":	400,
        "C":	100,
        "XC":	90,
        "L":	50,
        "XL":	40,
        "X":	10,
        "IX":	9,
    	"V": 	5,
        "IV":	4,
		"I":	1,
    }
    var err error
    var converted string
    if input <= 0 || input > 3000 {
        err = errors.New("value must be a positive integer including and between 1 and 3000")
    	return converted, err
    }
    for input > 0 {
    	var high int
        var roman string
		for k, v := range romanMap {
            if input >= v && v > high {
             	roman, high = k, v
            }
        }
    	converted += roman
        input -= high
	}
    return converted, err
}