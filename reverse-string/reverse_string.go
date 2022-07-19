package reverse

func Reverse(input string) (output string) {
	runes := []rune(input)
	for i := range runes {
		output += string(runes[len(runes) - i - 1])
	}
	return
}
