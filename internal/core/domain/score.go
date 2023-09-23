package domain

import "fmt"

type TeamScore struct {
	Team    Team
	Runs    Runs
	Wickets Wickets
	Overs   Overs
}

type Score struct {
	Home TeamScore
	Away TeamScore
}

func (s *Score) String() string {
	return fmt.Sprintf("%s %d-%d (%s) vs %s %d-%d (%s)", s.Home.Team, s.Home.Runs, s.Home.Wickets, s.Home.Overs.String(), s.Away.Team, s.Away.Runs, s.Away.Wickets, s.Away.Overs.String())
}
