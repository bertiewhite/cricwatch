package scoreservice

import (
	"cricwatch/internal/core/domain"
	"cricwatch/internal/core/ports"
	"os"
	"os/signal"
	"time"
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
	defer s.ScoreDisplayer.Close()
	t := time.Now()

	// The rest of this function is dumb code for testing purposes when there's not a live match
	sigintChan := make(chan os.Signal, 1)
	signal.Notify(sigintChan, os.Interrupt)
	for {
		select {
		case <-sigintChan:
			return nil
		default:
			if time.Since(t) < 10*time.Second {
				time.Sleep(1 * time.Second)
				continue
			}
			t = time.Now()
			score, err = s.ScoreRepo.GetScore(match)
			if err != nil {
				return err
			}

			err = s.ScoreDisplayer.Update(score)
			if err != nil {
				return err
			}
		}
	}
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
