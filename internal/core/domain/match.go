package domain

type Name string

func (n Name) String() string {
	return string(n)
}

type League string

func (l League) String() string {
	return string(l)
}

type Match struct {
	Name   Name
	League League
	ID     MatchID
	Live   bool
	Score  Score
}

type MatchID int

func (m MatchID) Int() int {
	return int(m)
}

type Matches map[League][]Match
