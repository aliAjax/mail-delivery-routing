package application

import (
	"context"
	"example.com/maildelivery/internal/tracking/domain"
)

type Summary struct{ Queued, Running, Delivered, Failed, Retry int }

func (s *Service) Summarize(ctx context.Context, tenant string) Summary {
	var out Summary
	for _, e := range s.Query(ctx, domain.Filter{TenantID: tenant}) {
		switch e.Status {
		case "queued":
			out.Queued++
		case "running":
			out.Running++
		case "delivered":
			out.Delivered++
		case "failed":
			out.Failed++
		case "retry":
			out.Retry++
		}
	}
	return out
}
