package isbn

import (
	"strconv"
	"strings"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.Replace(isbn, "-", "", -1)
	if len(isbn) != 10 {
		return false
	}
	var num int
	for i, r := range []rune(isbn) {
		var n int
		switch s := string(r); {
		case s == "X" && i == 9:
			n = 10
		default:
			d, err := strconv.Atoi(s)
			if err != nil {
				return false
			}
			n = d
		}
		num += n * (10 - i)
	}
	return num%11 == 0
}
