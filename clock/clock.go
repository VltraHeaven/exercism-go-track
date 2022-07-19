package clock

import (
    "fmt"
)

// Define the Clock type here.
type Clock struct {
    minute int
}

const (
    dayInMin int = 1440
)

func New(h, m int) Clock {
    return Clock{
        minute: ((h*60+m) % dayInMin + dayInMin) % dayInMin,
    }
}

func (c Clock) Add(m int) Clock {
    return New(0, c.minute+m)
}

func (c Clock) Subtract(m int) Clock {
    return New(0, c.minute-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minute / 60 , c.minute % 60)
}
