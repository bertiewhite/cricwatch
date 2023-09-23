package scoreservice

import (
	"cricwatch/internal/core/domain"
	"cricwatch/internal/core/ports"
)

type ScoreService struct {
	ScoreRepo      ports.ScoreRepo
	ScoreDisplayer ports.ScoreDisplayer
	UserInput      ports.UserInput
}

func NewScoreService(scoreRepo ports.ScoreRepo, scoreDisplayer ports.ScoreDisplayer, userInput ports.UserInput) *ScoreService {
	return &ScoreService{
		ScoreRepo:      scoreRepo,
		ScoreDisplayer: scoreDisplayer,
		UserInput:      userInput,
	}
}

type ScoreApplication interface {
	SelectMatch() (domain.Match, error)
	GetAndDisplayScore(match domain.Match) error
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

func (s *ScoreService) SelectMatch() (domain.Match, error) {
	matches, err := s.ScoreRepo.GetMatches()
	if err != nil {
		return domain.Match{}, err
	}

	match, err := s.UserInput.SelectMatch(matches)
	if err != nil {
		return domain.Match{}, err
	}

	return match, nil
}
