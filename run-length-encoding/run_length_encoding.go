package encode

import (
	"regexp"
	"strconv"
	"strings"
)

func RunLengthEncode(input string) (encoded string) {
	runes := []rune(input)
	var last string
	var count int
	appender := func(old string, char string, num int) string {
		if num > 1 {
			return old + strconv.Itoa(num) + char
		}
		return old + char
	}
	for i, r := range runes {
		c := string(r)
		if c != last {
			encoded = appender(encoded, last, count)
			last, count = "", 0
		}
		last = c
		count++
		if len(runes) == i+1 {
			encoded = appender(encoded, c, count)
		}
	}
	return
}

func RunLengthDecode(input string) (decoded string) {
	re := regexp.MustCompile(`(\d*)([A-Za-z]|\s)`)
	char := re.FindAllString(input, -1)
	for _, l := range char {
		re = regexp.MustCompile(`([\d]+)|([A-Za-z]|\s)`)
		split := re.FindAllString(l, -1)
		switch {
		case len(split) > 1:
			rep, _ := strconv.Atoi(split[0])
			decoded = decoded + strings.Repeat(split[1], rep)
		default:
			decoded = decoded + split[0]
		}
	}
	return
}
