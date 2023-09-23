package ports

import "cricwatch/internal/core/domain"

type ScoreDisplayer interface {
	Display(score domain.Score) error
}
