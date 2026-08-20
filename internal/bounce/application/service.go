package application

import (
	"context"
	"example.com/maildelivery/internal/bounce/domain"
	"sync"
	"time"
)

type Service struct {
	mu         sync.RWMutex
	events     []domain.Event
	suppressed map[string]time.Time
}

func New() *Service { return &Service{suppressed: map[string]time.Time{}} }
func (s *Service) Record(_ context.Context, e domain.Event) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.events = append(s.events, e)
	if e.Kind == domain.Hard || e.Kind == domain.Complaint {
		s.suppressed[e.Address] = time.Time{}
	}
}
func (s *Service) Suppressed(_ context.Context, address string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, ok := s.suppressed[address]
	return ok
}
func (s *Service) Events(_ context.Context) []domain.Event {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return append([]domain.Event(nil), s.events...)
}
