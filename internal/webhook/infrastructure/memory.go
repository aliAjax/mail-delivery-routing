package infrastructure

import (
	"context"
	"example.com/maildelivery/internal/webhook/domain"
	"sync"
)

type Repository struct {
	mu    sync.Mutex
	items []domain.Delivery
}

func (r *Repository) Save(_ context.Context, d domain.Delivery) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.items = append(r.items, d)
	return nil
}
func (r *Repository) List(_ context.Context) []domain.Delivery {
	r.mu.Lock()
	defer r.mu.Unlock()
	return append([]domain.Delivery(nil), r.items...)
}
