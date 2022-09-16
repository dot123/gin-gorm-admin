package store

import (
	"strings"
	"sync"
)

type Store[T any] struct {
	mux  sync.RWMutex
	data map[string]T
}

func New[T any](data map[string]T) *Store[T] {
	return &Store[T]{data: data}
}

func (s *Store[T]) Remove(key string) {
	s.mux.Lock()
	defer s.mux.Unlock()

	delete(s.data, key)
}

func (s *Store[T]) Has(key string) bool {
	s.mux.Lock()
	defer s.mux.Unlock()

	_, ok := s.data[key]

	return ok
}

func (s *Store[T]) Get(key string) T {
	s.mux.Lock()
	defer s.mux.Unlock()

	return s.data[key]
}

func (s *Store[T]) Set(key string, value T) {
	s.mux.Lock()
	defer s.mux.Unlock()

	if s.data == nil {
		s.data = make(map[string]T)
	}

	s.data[key] = value
}

func (s *Store[T]) SetIfLessThanLimit(key string, value T, maxAllowedElements int) bool {
	s.mux.Lock()
	defer s.mux.Unlock()

	if s.data == nil {
		s.data = make(map[string]T)
	}

	_, ok := s.data[key]

	if !ok && len(s.data) >= maxAllowedElements {
		return false
	}

	s.data[key] = value

	return true
}

func (s *Store[T]) LikeDeletes(likeKey string) {
	s.mux.Lock()
	defer s.mux.Unlock()

	for key := range s.data {
		if strings.Contains(key, likeKey) {
			delete(s.data, key)
		}
	}
}
