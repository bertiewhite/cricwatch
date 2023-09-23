package domain

type Runs int

func (r Runs) Int() int {
	return int(r)
}

func NewRuns(i int) Runs {
	return Runs(i)
}
