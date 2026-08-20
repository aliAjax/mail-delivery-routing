package infrastructure

import (
	"context"
	"example.com/maildelivery/internal/bounce/domain"
	"sync"
)

type Repository struct {
	mu     sync.Mutex
	events []domain.Event
}

func (r *Repository) Save(_ context.Context, e domain.Event) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.events = append(r.events, e)
	return nil
}
func (r *Repository) List(_ context.Context) []domain.Event {
	r.mu.Lock()
	defer r.mu.Unlock()
	return append([]domain.Event(nil), r.events...)
}
