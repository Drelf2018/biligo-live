package utils

import (
	"encoding/json"
	"sync"
)

type Slice[T any] struct {
	rw   sync.RWMutex
	data []T
}

func (s *Slice[T]) Append(data ...T) {
	s.rw.Lock()
	defer s.rw.Unlock()
	s.data = append(s.data, data...)
}

func (s *Slice[T]) Index(i int) T {
	s.rw.RLock()
	defer s.rw.RUnlock()
	return s.data[i]
}

func (s *Slice[T]) Data() []T {
	s.rw.RLock()
	defer s.rw.RUnlock()
	return s.data
}

func (s *Slice[T]) Len() int {
	s.rw.RLock()
	defer s.rw.RUnlock()
	return len(s.data)
}

func (s *Slice[T]) Iter() <-chan T {
	ch := make(chan T)
	s.rw.RLock()
	go func() {
		defer s.rw.RUnlock()
		for _, t := range s.data {
			ch <- t
		}
		close(ch)
	}()
	return ch
}

func (s *Slice[T]) MarshalJSON() ([]byte, error) {
	s.rw.RLock()
	defer s.rw.RUnlock()
	return json.Marshal(s.data)
}

func NewSlice[T any](data ...T) *Slice[T] {
	return &Slice[T]{data: data}
}
