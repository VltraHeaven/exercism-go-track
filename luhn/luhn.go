package luhn

import (
    "strings"
    "strconv"
)

// Valid returns a bool that signifies if the given string is a compliant identifiation number
func Valid(id string) bool {
	var idSlice = []rune(strings.ReplaceAll(id, " ", ""))
	if len(idSlice) <= 1 {
        return false
    }
    var sum int
	for i := len(idSlice) - 1; i >= 0; i-- {
        v, err := strconv.Atoi(string(idSlice[i]))
        switch {
            case err != nil:
        	return false
        	case (i - len(idSlice)) % 2 == 0 && v * 2 > 9 :
            sum += v * 2 - 9
            case (i - len(idSlice)) % 2 == 0:
        	sum += v * 2
            default:
        	sum += v
        }
    }
    return sum % 10 == 0
}
