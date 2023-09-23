package domain

import "fmt"

type Overs struct {
	overs int
	balls int
}

func NewOvers(overs, balls int) Overs {
	return Overs{overs, balls}
}

func (o Overs) String() string {
	return fmt.Sprintf("%d.%d", o.overs, o.balls)
}
