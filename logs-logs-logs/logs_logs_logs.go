package logs

import (
    "unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, r := range log {
        switch r {
            case '❗':	// U+2757
        	return "recommendation"
            case '🔍':	// U+1F50D
        	return "search"
            case '☀':	// U+2600
        	return "weather"
        }
    }
    return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
    var modLog []rune
	for _, r := range log {
        switch r {
            case oldRune:
        	modLog = append(modLog, newRune)
            default:
        	modLog = append(modLog, r)
        }
    }
	return string(modLog)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
