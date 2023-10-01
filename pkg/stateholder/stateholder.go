package state

import "sync"

type Holder[T any] struct {
	state       T
	subscribers []chan<- struct{}
	mu          sync.Mutex
}

func New[T any](init T) Holder[T] {
	return Holder[T]{
		state:       init,
		subscribers: []chan<- struct{}{},
	}
}

func (s *Holder[T]) Update(newState T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.state = newState

	for _, subscriber := range s.subscribers {
		select {
		case subscriber <- struct{}{}:
		default:
		}
	}
}

func (s *Holder[T]) Get() T {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.state
}

func (s *Holder[T]) Subscribe() <-chan struct{} {
	s.mu.Lock()
	defer s.mu.Unlock()

	ch := make(chan struct{}, 1)
	s.subscribers = append(s.subscribers, ch)

	return ch
}
