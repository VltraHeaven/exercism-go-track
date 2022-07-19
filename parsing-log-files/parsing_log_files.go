package parsinglogfiles

import (
    "regexp"
    "fmt"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
    return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<([^a-z]*)>`)
    return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) (count int) {
	re := regexp.MustCompile(`"(.*)\b(?i)password\b*"`)
    for _, l := range lines {
        if re.MatchString(l) {
            count++
        }
    }
	return
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line[\d]+`)
    return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
    var newlines []string
	re := regexp.MustCompile(`(User)(\s+)([A-Z|a-z|0-9]+)`)
    for _, l := range lines {
        switch {
        case re.MatchString(l):
            u := re.FindStringSubmatch(l)
            newlines = append(newlines, fmt.Sprintf("%s %s %s", "[USR]", u[3], l))
        default:
        	newlines = append(newlines, l)
        }
    }
	return newlines
}
