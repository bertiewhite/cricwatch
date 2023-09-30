package poller

import (
	"errors"
	"time"
)

type Poller struct {
	action     func() error
	errHandler func(error)
	stop       chan struct{}
}

func NewPoller(action func() error, errHandler func(error)) *Poller {
	return &Poller{
		action:     action,
		errHandler: errHandler,
	}
}

func (p *Poller) Start(interval time.Duration) {
	ticker := time.NewTicker(interval)

	stopChan := make(chan struct{})

	p.stop = stopChan

	go func() {
		for {
			select {
			case <-ticker.C:
				err := p.action()
				if err != nil {
					p.errHandler(err)
				}
			case <-p.stop:
				return
			}
		}
	}()
}

func (p *Poller) Stop() error {
	if p.stop == nil {
		return errors.New("not started this poller silly billy")
	}

	close(p.stop)
	p.stop = nil

	return nil
}

func (p *Poller) Close() {
	_ = p.Stop() // doens't matter if this fails
	p.action = nil
	p.errHandler = nil
	p.stop = nil
}
