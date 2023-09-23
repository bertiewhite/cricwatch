package ports

import "cricwatch/internal/core/domain"

type ScoreRepo interface {
	GetScore(match domain.Match) (domain.Score, error)
}
