package ports

import "cricwatch/internal/core/domain"

type ScoreDisplayer interface {
	Display(score domain.Score) error
	Update(score domain.Score) error
	Close() error
}
