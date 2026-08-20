package application

import (
	"context"
	"example.com/maildelivery/internal/queue/domain"
	"sync"
	"testing"
	"time"
)

func TestClaimBatchDoesNotDuplicateJobs(t *testing.T) {
	s := New()
	if err := s.Enqueue(context.Background(), domain.Job{ID: "j1", Status: "queued"}); err != nil {
		t.Fatal(err)
	}
	start := make(chan struct{})
	results := make(chan []domain.Job, 2)
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			results <- s.ClaimBatch(context.Background(), time.Now(), 2)
		}()
	}
	close(start)
	wg.Wait()
	close(results)
	seen := map[string]bool{}
	count := 0
	for batch := range results {
		for _, job := range batch {
			if seen[job.ID] {
				t.Fatalf("job %s claimed twice", job.ID)
			}
			seen[job.ID] = true
			count++
		}
	}
	if count != 1 {
		t.Fatalf("claimed=%d, want 1", count)
	}
}

func TestClaimBatchHonorsLimit(t *testing.T) {
	s := New()
	for _, id := range []string{"j1", "j2"} {
		if err := s.Enqueue(context.Background(), domain.Job{ID: id, Status: "queued"}); err != nil {
			t.Fatal(err)
		}
	}
	if got := s.ClaimBatch(context.Background(), time.Now(), 2); len(got) != 2 {
		t.Fatalf("claimed=%d, want 2", len(got))
	}
}

func TestClaimBatchPersistsState(t *testing.T) {
	s := New()
	if err := s.Enqueue(context.Background(), domain.Job{ID: "j1", Status: "queued"}); err != nil {
		t.Fatal(err)
	}
	if got := s.ClaimBatch(context.Background(), time.Now(), 2); len(got) != 1 {
		t.Fatalf("first claim=%d, want 1", len(got))
	}
	if got := s.ClaimBatch(context.Background(), time.Now(), 2); len(got) != 0 {
		t.Fatalf("second claim=%d, want 0", len(got))
	}
}
