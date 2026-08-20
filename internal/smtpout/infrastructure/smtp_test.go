package infrastructure

import (
	"context"
	"errors"
	"example.com/maildelivery/internal/smtpout/domain"
	"sync"
	"testing"
	"time"
)

type retryTransport struct {
	mu      sync.Mutex
	calls   int
	started chan struct{}
	mode    string
}

func (t *retryTransport) Send(ctx context.Context, _, _, _, _ string) error {
	t.mu.Lock()
	t.calls++
	if t.started != nil && t.calls == 1 {
		close(t.started)
	}
	t.mu.Unlock()
	if t.mode == "cancel" {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(30 * time.Millisecond):
			return domain.ErrTemporary
		}
	}
	if t.calls == 1 {
		return domain.ErrTemporary
	}
	return nil
}

func TestRetryingTransportHonorsCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	fake := &retryTransport{started: make(chan struct{}), mode: "cancel"}
	done := make(chan error, 1)
	go func() { done <- (Retrying{Transport: fake, Attempts: 3}).Send(ctx, "a", "b", "s", "body") }()
	<-fake.started
	cancel()
	select {
	case err := <-done:
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("error=%v, want context.Canceled", err)
		}
	case <-time.After(200 * time.Millisecond):
		t.Fatal("retry did not stop after cancellation")
	}
}

func TestRetryingTransportRetriesTemporaryFailure(t *testing.T) {
	fake := &retryTransport{}
	if err := (Retrying{Transport: fake, Attempts: 3}).Send(context.Background(), "a", "b", "s", "body"); err != nil {
		t.Fatal(err)
	}
	if fake.calls != 2 {
		t.Fatalf("calls=%d, want 2", fake.calls)
	}
}
