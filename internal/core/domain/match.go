package domain

type Match struct {
	ID MatchID
}

type MatchID int

func (m MatchID) Int() int {
	return int(m)
}
