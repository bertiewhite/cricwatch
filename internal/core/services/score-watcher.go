package scoreservice

import (
	"github.com/bertiewhite/cricwatch/internal/core/domain"
	"github.com/bertiewhite/cricwatch/internal/core/ports"
	state "github.com/bertiewhite/cricwatch/pkg/stateholder"
)

type ScoreWatcher struct {
	scoreHolder    *state.Holder[domain.Score]
	stop           chan struct{}
	scoreDisplayer ports.ScoreDisplayer
}

func NewScoreWatcher(scoreDisplayer ports.ScoreDisplayer, scoreHolder *state.Holder[domain.Score]) *ScoreWatcher {
	stop := make(chan struct{})

	return &ScoreWatcher{
		scoreHolder:    scoreHolder,
		scoreDisplayer: scoreDisplayer,
		stop:           stop,
	}
}

func (s *ScoreWatcher) Start() {
	notification := s.scoreHolder.Subscribe()
	go func() {
		for {
			select {
			case <-notification:
				s.scoreDisplayer.Update(s.scoreHolder.Get())
			case <-s.stop:
				return
			}
		}
	}()
}

func (s *ScoreWatcher) Close() {
	close(s.stop)
}
