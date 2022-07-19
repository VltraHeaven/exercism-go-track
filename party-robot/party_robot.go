package partyrobot

import (
    "fmt";
    "strings"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	var assignment = []string{
        Welcome(name),
        fmt.Sprintf("\nYou have been assigned to table %03d.", table),
        fmt.Sprintf(" Your table is %s, ", direction),
		fmt.Sprintf("exactly %2.1f meters from here.\n", distance),
		fmt.Sprintf("You will be sitting next to %s.", neighbor),
    }
	return strings.Join(assignment, "")
}
