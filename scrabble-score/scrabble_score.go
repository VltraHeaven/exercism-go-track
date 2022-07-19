package scrabble

import (
    "strings"
)

var letters = map[rune]int {
    'A': 1, 'E': 1, 'I': 1, 'O': 1, 'U': 1, 'L': 1, 'N': 1, 'R': 1, 'S': 1, 'T': 1,
    'F': 4, 'H': 4, 'V': 4, 'W': 4, 'Y': 4,
    'B': 3, 'C': 3, 'M': 3, 'P': 3,
    'D': 2, 'G': 2,
    'J': 8, 'X': 8,
    'Q': 10, 'Z': 10,
    'K': 5,
}

func letterNav(r rune, c chan int) {
	if _, ok := letters[r]; ok {
		c <- letters[r]
	}
}

func Score(word string) (score int) {
	runic := []rune(strings.ToUpper(word))
    scoreChan := make(chan int, len(runic))
    for _, v := range runic {
        go letterNav(v, scoreChan)
        score += <- scoreChan
    }
	return
}
