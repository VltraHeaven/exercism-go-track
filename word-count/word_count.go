package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	re := regexp.MustCompile(`\w+('\w+)?`)
	count := make(Frequency)
	for _, w := range re.FindAllString(strings.ToLower(phrase), -1) {
		count[w]++
	}
	return count
}
