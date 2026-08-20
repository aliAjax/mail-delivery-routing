package application

import (
	"context"
	"example.com/maildelivery/internal/tenant/domain"
	"fmt"
	"sync"
)

type Service struct {
	mu      sync.RWMutex
	tenants map[string]domain.Tenant
}

func New() *Service { return &Service{tenants: map[string]domain.Tenant{}} }
func (s *Service) Put(_ context.Context, t domain.Tenant) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tenants[t.ID] = t
}
func (s *Service) Get(_ context.Context, id string) (domain.Tenant, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	t, ok := s.tenants[id]
	return t, ok
}
func (s *Service) Pause(_ context.Context, id string, p bool) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	t, ok := s.tenants[id]
	if !ok {
		return false
	}
	t.Paused = p
	s.tenants[id] = t
	return true
}

// Reserve atomically reserves one unit of a tenant's daily quota. The entire
// read-modify-write is performed under the write lock so that concurrent
// reservations for the same tenant cannot interleave: they are serialized, the
// per-tenant quota is never over-consumed, and the tenants map is never read
// and written concurrently.
func (s *Service) Reserve(_ context.Context, id string) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	t, ok := s.tenants[id]
	if !ok {
		return 0, fmt.Errorf("tenant %s not found", id)
	}
	if err := t.Consume(); err != nil {
		return t.Used, err
	}
	s.tenants[id] = t
	return t.Used, nil
}

func (s *Service) Snapshot(_ context.Context) []domain.Tenant {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]domain.Tenant, 0, len(s.tenants))
	for _, tenant := range s.tenants {
		out = append(out, tenant)
	}
	return out
}
