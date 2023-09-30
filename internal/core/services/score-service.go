package scoreservice

import (
	"cricwatch/internal/core/domain"
	"cricwatch/internal/core/ports"
	"cricwatch/pkg/poller"
	"fmt"
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

	getAndDisplay := func() error {
		score, err = s.ScoreRepo.GetScore(match)
		if err != nil {
			return err
		}

		err = s.ScoreDisplayer.Update(score)
		if err != nil {
			return err
		}
		return nil
	}

	errHandler := func(err error) {
		fmt.Println(fmt.Sprintf("Uh oh something went wrong: %s", err.Error()))
		_ = s.ScoreDisplayer.Close()

		os.Exit(1)
	}

	poller := poller.NewPoller(getAndDisplay, errHandler)

	poller.Start(30 * time.Second)
	sigintChan := make(chan os.Signal, 1)
	signal.Notify(sigintChan, os.Interrupt)

	for {
		select {
		case <-sigintChan:
			poller.Stop()
			os.Exit(0)
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
