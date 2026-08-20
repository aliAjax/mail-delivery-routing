package application

import (
	"context"
	"example.com/maildelivery/internal/template/domain"
	"fmt"
	"sync"
)

type Service struct {
	mu   sync.RWMutex
	data map[string]domain.Template
}

func New() *Service { return &Service{data: map[string]domain.Template{}} }
func (s *Service) Put(_ context.Context, t domain.Template) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[t.ID] = t
}
func (s *Service) Render(_ context.Context, id string, v map[string]string) (string, error) {
	s.mu.RLock()
	t, ok := s.data[id]
	s.mu.RUnlock()
	if !ok {
		return "", fmt.Errorf("template %s not found", id)
	}
	return t.Render(v)
}

func (s *Service) RenderWithFallback(ctx context.Context, id string, vars, defaults map[string]string) (string, error) {
	if ctx == nil {
		return "", fmt.Errorf("template context required")
	}
	if err := ctx.Err(); err != nil {
		return "", err
	}
	s.mu.RLock()
	t, ok := s.data[id]
	s.mu.RUnlock()
	if !ok {
		return "", fmt.Errorf("template %s not found", id)
	}
	return t.RenderWithDefaults(vars, defaults)
}
