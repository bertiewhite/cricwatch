package domain

type Wickets int

func (w Wickets) Int() int {
	return int(w)
}

func NewWickets(i int) Wickets {
	return Wickets(i)
}
