package domain

type Team string

func NewTeam(t string) Team {
	return Team(t)
}
