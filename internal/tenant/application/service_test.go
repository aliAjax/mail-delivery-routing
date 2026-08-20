package application

import (
	"context"
	"example.com/maildelivery/internal/tenant/domain"
	"sync"
	"testing"
)

func TestReserveQuotaConcurrent(t *testing.T) {
	s := New()
	s.Put(context.Background(), domain.Tenant{ID: "acme", DailyQuota: 32})
	start := make(chan struct{})
	var wg sync.WaitGroup
	var mu sync.Mutex
	successes := 0
	for i := 0; i < 64; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			if _, err := s.Reserve(context.Background(), "acme"); err == nil {
				mu.Lock()
				successes++
				mu.Unlock()
			}
		}()
	}
	close(start)
	wg.Wait()
	if successes != 32 {
		t.Fatalf("successes=%d, want 32", successes)
	}
	tenants := s.Snapshot(context.Background())
	if len(tenants) != 1 || tenants[0].Used != 32 {
		t.Fatalf("snapshot=%+v, want one tenant at usage 32", tenants)
	}
}
