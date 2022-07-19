package robotname

import (
    "math/rand"
    "strings"
    "time"
    "strconv"
    "errors"
)

const (
    chars = "QWERTYUIOPASDFGHJKLZXCVBNM"
    max = 26 * 26 * 10 * 10 * 10
)

// Define the Robot type here.
type Robot struct {
    name string
}

// NamedRobots keeps a list of named robots with assigned names
type NamedRobots struct {
    robots map[string]*Robot
}

// NewNamedRobots returns a pointer to an instance of type NamedRobots
func NewNamedRobots() *NamedRobots {
    return &NamedRobots{robots: map[string]*Robot{}}
}

var named = NewNamedRobots()

// Name assigns a name to an instance of type Robot  
func (r *Robot) Name() (name string, err error) {
    if r.name != "" {
        name = r.name
        return
    } else if len(named.robots) >= max {
    	err = errors.New("maximum limit of possible names has been reached")
        return
    }
    name = buildName()
	switch check := named.robots[name] != nil; {
        case check:
    	name, err = r.Name()
        default:
    	r.name = name
        named.robots[name] = r
    }
    return
}

// buildName returns a randomized 5 hexidecimal-character length string 
func buildName() string {
    rand.Seed(time.Now().UnixNano())
	var n strings.Builder
    for i := 0; i <= 4; i++ {
        switch {
            case i < 2:
        	n.WriteString(string(chars[rand.Intn(len(chars))]))
            default:
        	n.WriteString(strconv.Itoa(rand.Intn(9)))
    	}
    }
	return n.String()
}

// Reset removes an instance of type Robot from a list of named robots and clears it's name
func (r *Robot) Reset() {
    delete(named.robots, r.name)
    r.name = ""
}
