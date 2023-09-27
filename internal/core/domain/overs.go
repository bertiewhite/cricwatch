package domain

import "fmt"

type Overs struct {
	overs int
	balls int
}

type Balls int

func NewBalls(balls int) Balls {
	return Balls(balls)
}

func NewOvers(overs, balls int) Overs {
	return Overs{overs, balls}
}

func (o Overs) String() string {
	return fmt.Sprintf("%d.%d", o.overs, o.balls)
}
