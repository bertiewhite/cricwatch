package domain

import "fmt"

type Batsman struct {
	Name         Name
	CurrentScore Runs
	BallsFaced   Balls
	Facing       bool
}

func (b Batsman) String() string {
	s := fmt.Sprintf("%s %d (%d)", b.Name, b.CurrentScore, b.BallsFaced)

	if b.Facing {
		s = "*" + s
	}

	return s
}
