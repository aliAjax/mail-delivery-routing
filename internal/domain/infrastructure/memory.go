package infrastructure

import (
	"context"
	"example.com/maildelivery/internal/domain/domain"
	"sync"
)

type Repository struct {
	mu   sync.RWMutex
	data map[string]domain.Domain
}

func New() *Repository { return &Repository{data: map[string]domain.Domain{}} }
func (r *Repository) Save(_ context.Context, d domain.Domain) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[d.ID] = d
	return nil
}
func (r *Repository) Find(_ context.Context, id string) (domain.Domain, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	d, ok := r.data[id]
	return d, ok
}
