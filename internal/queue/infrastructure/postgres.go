package infrastructure

import (
	"context"
	"time"
)

type Lease struct {
	ID, Owner string
	ExpiresAt time.Time
}
type Repository struct{}

func (Repository) Acquire(_ context.Context, id, owner string, ttl time.Duration) (Lease, bool) {
	return Lease{ID: id, Owner: owner, ExpiresAt: time.Now().Add(ttl)}, true
}
func (Repository) Release(_ context.Context, _ Lease) error { return nil }
