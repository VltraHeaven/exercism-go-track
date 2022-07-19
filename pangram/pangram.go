package pangram

import (
    "regexp"
    "strings"
)

func IsPangram(input string) bool {
	alphabet := map[rune]int{}
	re := regexp.MustCompile(`[a-z]`)
    for _, r := range strings.ToLower(input) {
        if re.MatchString(string(r)) {
        	alphabet[r]++
        }
    }
	return len(alphabet) == 26
}
