package application

import (
	"context"
	"example.com/maildelivery/internal/domain/domain"
	"sync"
)

type Service struct {
	mu   sync.RWMutex
	data map[string]domain.Domain
}

func New() *Service { return &Service{data: map[string]domain.Domain{}} }
func (s *Service) Register(_ context.Context, d domain.Domain) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[d.ID] = d
}
func (s *Service) Get(_ context.Context, id string) (domain.Domain, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	v, ok := s.data[id]
	return v, ok
}
func (s *Service) Verify(_ context.Context, id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	v, ok := s.data[id]
	if !ok {
		return false
	}
	v.Verify()
	s.data[id] = v
	return true
}
