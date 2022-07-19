package booking

import (
    "time";
    "fmt"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
    a, err := time.Parse("1/2/2006 15:04:05", date)
    if err != nil {
        panic(err)
    }
	return a
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	n := time.Now()
    a, err := time.Parse("January 2, 2006 15:04:05", date)
    if err != nil {
        panic(err)
    }
    return n.After(a)
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	a, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
    if err != nil {
        panic(err)
    }
    t := a.Hour() * 100 + a.Minute()
    switch {
    case t >= 1200 && t < 1800:
    	return true
    }
	return false
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
    a := Schedule(date)
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.",
                       a.Weekday(),
                       a.Month(),
                       a.Day(),
                       a.Year(),
                       a.Hour(),
                       a.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	anniversary, err := time.Parse("2006-01-02", "2022-09-15")
    if err != nil {
        panic(err)
    }
    return anniversary
}
