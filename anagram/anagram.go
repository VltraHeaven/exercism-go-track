package anagram

import (
	"strings"
)

func Detect(subject string, candidates []string) (anagrams []string) {
	isAnagram := func(s string, c string) bool {
		// words are not anagrams of themselves
		if s == c {
			return false
		}
		for _, r := range c {
			i := strings.IndexRune(s, r)
			// does not detect anagram subsets
			if i == -1 {
				return false
			}
			s = s[:i] + s[i+1:]
		}
		return s == ""
	}
	for _, candidate := range candidates {
		if isAnagram(strings.ToLower(subject), strings.ToLower(candidate)) {
			anagrams = append(anagrams, candidate)
		}
	}
	return
}
