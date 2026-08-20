package application

import (
	"context"
	"example.com/maildelivery/internal/policy/domain"
	"sync"
)

type Service struct {
	mu   sync.RWMutex
	data map[string]domain.Policy
}

func New() *Service { return &Service{data: map[string]domain.Policy{}} }
func (s *Service) Put(_ context.Context, p domain.Policy) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[p.TenantID] = p
}
func (s *Service) Allows(ctx context.Context, tenant, address string) bool {
	s.mu.RLock()
	p, ok := s.data[tenant]
	s.mu.RUnlock()
	if !ok {
		return true
	}
	return p.Allows(address)
}
func (s *Service) Context(_ context.Context) string { return "policy-service" }
