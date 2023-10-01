package ports

import "github.com/bertiewhite/cricwatch/internal/core/domain"

type ScoreRepo interface {
	GetScore(match domain.Match) (domain.Score, error)
	GetMatches() (domain.Matches, error)
}
