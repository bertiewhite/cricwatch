package scoreservice

import "cricwatch/internal/core/ports"

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

func (s *ScoreService) GetAndDisplayScore(id int) error {
	score, err := s.ScoreRepo.GetScore(id)
	if err != nil {
		return err
	}

	err = s.ScoreDisplayer.Display(score)
	if err != nil {
		return err
	}

	return nil
}
