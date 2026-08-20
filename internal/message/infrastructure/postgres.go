package infrastructure

import (
	"context"
	"example.com/maildelivery/internal/platform/store"
)

type Repository struct{ Store store.Store }

func New(s store.Store) *Repository                  { return &Repository{Store: s} }
func (r *Repository) Health(_ context.Context) error { return nil }
func (r *Repository) Name() string                   { return "postgres-compatible-message-store" }
