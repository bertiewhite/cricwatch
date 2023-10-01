package ports

import "github.com/bertiewhite/cricwatch/internal/core/domain"

type UserInput interface {
	SelectMatch(matches domain.Matches) (domain.Match, error)
}
