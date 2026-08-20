package application

import (
	"context"
	"example.com/maildelivery/internal/tenant/domain"
)

func ReservationSnapshot(s *Service, ctx context.Context, id string) (domain.Tenant, error) {
	if _, err := s.Reserve(ctx, id); err != nil {
		return domain.Tenant{}, err
	}
	t, ok := s.Get(ctx, id)
	if !ok {
		return domain.Tenant{}, context.Canceled
	}
	return t, nil
}
