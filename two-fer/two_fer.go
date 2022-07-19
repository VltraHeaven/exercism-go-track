package twofer
import "fmt"
// ShareWith stores a given string in the 'name' variable.
// If no string is submitted the string 'you' is stored.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}