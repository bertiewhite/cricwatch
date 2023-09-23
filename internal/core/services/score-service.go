package scoreservice

import (
	"cricwatch/internal/core/domain"
	"cricwatch/internal/core/ports"
)

type ScoreService struct {
	ScoreRepo      ports.ScoreRepo
	ScoreDisplayer ports.ScoreDisplayer
}

func NewScoreService(scoreRepo ports.ScoreRepo, scoreDisplayer ports.ScoreDisplayer) *ScoreService {
	return &ScoreService{
		ScoreRepo:      scoreRepo,
		ScoreDisplayer: scoreDisplayer,
	}
}

func (s *ScoreService) GetAndDisplayScore(match domain.Match) error {
	score, err := s.ScoreRepo.GetScore(match)
	if err != nil {
		return err
	}

	err = s.ScoreDisplayer.Display(score)
	if err != nil {
		return err
	}

	return nil
}
