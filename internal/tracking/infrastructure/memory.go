package infrastructure

import (
	"context"
	"example.com/maildelivery/internal/tracking/domain"
)

type Repository struct{ Events []domain.Event }

func (r *Repository) Save(_ context.Context, e domain.Event) error {
	r.Events = append(r.Events, e)
	return nil
}
func (r *Repository) Count(_ context.Context) int { return len(r.Events) }
