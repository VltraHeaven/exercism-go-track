package tournament

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
	"text/tabwriter"
)

// Team provides team name and match records
type Team struct {
	Name                             string
	Matches, win, lose, draw, points int
}

// NewTeam constructs Team objects
func NewTeam(name string) *Team {
	return &Team{Name: name}
}

// Tally receives an io.Reader with team records and writes processed data to an io.Writer
func Tally(reader io.Reader, writer io.Writer) (err error) {
	teams := map[string]*Team{}

	t, err := ioutil.ReadAll(reader)
	if err != nil {
		return
	}

	matches := strings.Split(strings.TrimSpace(string(t)), "\n")
	for _, m := range matches {
		regOk, _ := regexp.Match("#", []byte(m))
		if len(m) == 0 || regOk {
			continue
		}
		match := strings.Split(m, ";")

		if len(match) != 3 {
			err = errors.New("invalid team name")
			return
		}

		for i := 0; i < 2; i++ {
			_, ok := teams[match[i]]
			if !ok {
				teams[match[i]] = NewTeam(match[i])
			}
		}

		teams[match[0]].Matches += 1
		teams[match[1]].Matches += 1

		switch match[2] {
		case "win":
			teams[match[0]].win += 1
			teams[match[0]].points += 3
			teams[match[1]].lose += 1
		case "loss":
			teams[match[1]].win += 1
			teams[match[1]].points += 3
			teams[match[0]].lose += 1
		case "draw":
			teams[match[0]].draw += 1
			teams[match[0]].points += 1
			teams[match[1]].draw += 1
			teams[match[1]].points += 1
		default:
			err = errors.New("invalid match result")
			return
		}
	}

	var records []*Team
	for _, team := range teams {
		records = append(records, team)
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].Name < records[j].Name
	})

	sort.SliceStable(records, func(i, j int) bool {
		return records[i].points > records[j].points
	})

	table := tabwriter.NewWriter(writer, 0, 2, 1, ' ', tabwriter.Debug)
	fmt.Fprintf(table, "%-30s\t%3s\t%3s\t%3s\t%3s\t%3s\n", "Team", "MP", "W", "D", "L", "P")
	for _, t := range records {
		fmt.Fprintf(table, "%-30s\t%3d\t%3d\t%3d\t%3d\t%3d\n", t.Name, t.Matches, t.win, t.draw, t.lose, t.points)
	}
	table.Flush()
	return
}
