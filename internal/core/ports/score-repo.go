package ports

import "cricwatch/internal/core/domain"

type ScoreRepo interface {
	GetScore(id int) (domain.Score, error)
}
