package isogram

import (
    "strings"
)

func IsIsogram(word string) bool {
    // Empty strings are always isograms
	if len(word) <= 0 {
        return true
    }
	// Remove spaces and hypens
	var yeet = []string{" ", "-"}
    for i := range yeet {
        word = strings.ReplaceAll(word, yeet[i], "")
    }
	// All letters to lowercase then check for duplicate runes
	w := []rune(strings.ToLower(word))
    record := make(map[rune]int) 
	for _, v := range w {
        _, ok := record[v]
        if ok {
        	return false
        }
    	record[v]++
    }
	return true
}
