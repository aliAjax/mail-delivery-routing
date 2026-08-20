package application

import (
	"context"
	"example.com/maildelivery/internal/queue/domain"
	"sync"
	"time"
)

type Handler func(context.Context, domain.Job) error

func (s *Service) Run(ctx context.Context, n int, h Handler) func() {
	if n < 1 {
		n = 1
	}
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ticker := time.NewTicker(100 * time.Millisecond)
			defer ticker.Stop()
			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					j, ok := s.Claim(ctx, time.Now())
					if ok {
						s.Finish(ctx, j.ID, h(ctx, j))
					}
				}
			}
		}()
	}
	return func() { wg.Wait() }
}
