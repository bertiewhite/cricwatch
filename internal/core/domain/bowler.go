package domain

import "fmt"

type Bowler struct {
	Name    Name
	Over    Overs
	Wickets Wickets
	Runs    Runs
}

func (b Bowler) String() string {
	return fmt.Sprintf("%s %d-%d %s", b.Name, b.Runs.Int(), b.Wickets.Int(), b.Over.String())
}
