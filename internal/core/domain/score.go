package domain

import "fmt"

type TeamScore struct {
	Team    Team
	Runs    Runs
	Wickets Wickets
	Overs   Overs
	Batting bool
}

type CurrentState struct {
	BatterA    Batters
	BatterB    Batters
	Bowler     Bowler
	PrevBowler Bowler
}

type Score struct {
	Home         TeamScore
	Away         TeamScore
	CurrentState CurrentState
	AllBowlers   []Bowler
	AllBatters   []Batters
}

func (s *Score) String() string {
	headerLine := fmt.Sprintf("%s %d-%d (%s) vs %s %d-%d (%s)", s.Home.Team, s.Home.Runs, s.Home.Wickets, s.Home.Overs.String(), s.Away.Team, s.Away.Runs, s.Away.Wickets, s.Away.Overs.String())
	battingLine := fmt.Sprintf("%s %d (%d), %s %d (%d)", s.CurrentState.BatterA.Name, s.CurrentState.BatterA.CurrentScore, s.CurrentState.BatterA.BallsFaced, s.CurrentState.BatterB.Name, s.CurrentState.BatterB.CurrentScore, s.CurrentState.BatterB.BallsFaced)
	bowlingLine := fmt.Sprintf("%s: %s-%d-%d", s.CurrentState.Bowler.Name, s.CurrentState.Bowler.Over.String(), s.CurrentState.Bowler.Wickets, s.CurrentState.Bowler.Runs)
	prevBowlingLine := fmt.Sprintf("%s: %s-%d-%d", s.CurrentState.PrevBowler.Name, s.CurrentState.PrevBowler.Over.String(), s.CurrentState.PrevBowler.Wickets, s.CurrentState.PrevBowler.Runs)

	return fmt.Sprintf("%s\n%s\n%s\n%s", headerLine, battingLine, bowlingLine, prevBowlingLine)
}
