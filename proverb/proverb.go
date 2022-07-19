// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb takes a list of strings and generates a proverbial rhyme
// https://golang.org/doc/effective_go.html#commentary
package proverb

import (
    "fmt"
)

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	var proverb []string
    if len(rhyme) > 0 {
        var s0 = rhyme[0]
    	for i, v := range rhyme {
    		var s1 = v
        	if len(rhyme) > 1 && i+1 < len(rhyme) {
                var s2 = rhyme[i+1]
            	joinstr := "For want of a "+ s1 +" the "+ s2 +" was lost."
            	fmt.Println(joinstr)
        		proverb = append(proverb, joinstr)
        	} else {
        		proverb = append(proverb, "And all for the want of a "+ s0 +".")
                break
        	}
    	}
    }
    return proverb
}
