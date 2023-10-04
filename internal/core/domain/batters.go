package domain

import "fmt"

type Batters struct {
	Name         Name
	CurrentScore Runs
	BallsFaced   Balls
	Facing       bool
	Out          bool
}

func (b Batters) String() string {
	s := fmt.Sprintf("%s %d (%d)", b.Name, b.CurrentScore, b.BallsFaced)

	if b.Facing {
		s = "*" + s
	}

	return s
}
