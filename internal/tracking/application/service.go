package application

import (
	"context"
	"example.com/maildelivery/internal/tracking/domain"
	"sync"
)

type Service struct {
	mu     sync.RWMutex
	events []domain.Event
}

func New() *Service { return &Service{events: make([]domain.Event, 0)} }
func (s *Service) Append(_ context.Context, e domain.Event) {
	if !e.Valid() {
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.events = append(s.events, e)
}
func (s *Service) Query(_ context.Context, f domain.Filter) []domain.Event {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]domain.Event, 0)
	for _, e := range s.events {
		if f.TenantID != "" && e.TenantID != f.TenantID {
			continue
		}
		if f.Status != "" && e.Status != f.Status {
			continue
		}
		if !f.From.IsZero() && e.At.Before(f.From) {
			continue
		}
		if !f.To.IsZero() && e.At.After(f.To) {
			continue
		}
		out = append(out, e)
	}
	return out
}
