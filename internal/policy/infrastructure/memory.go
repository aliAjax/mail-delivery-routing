package infrastructure

import (
	"context"
	"example.com/maildelivery/internal/policy/domain"
)

type Repository struct{ Items []domain.Policy }

func (r *Repository) Save(_ context.Context, p domain.Policy) error {
	r.Items = append(r.Items, p)
	return nil
}
