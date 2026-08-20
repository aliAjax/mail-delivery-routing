package domain

import (
	"context"
	"errors"
	"time"
	"testing"
)

type exhaustedTransport struct{}

func (exhaustedTransport) Send(context.Context, string, string, string, string) error {
	return ErrTemporary
}

func TestSendWithRetryPreservesExhaustedErrors(t *testing.T) {
	err := SendWithRetry(context.Background(), exhaustedTransport{}, "a", "b", "s", "body", 2)
	if !errors.Is(err, ErrAttemptsExhausted) || !errors.Is(err, ErrTemporary) {
		t.Fatalf("err=%v, want both exhaustion and temporary errors", err)
	}
}

type contextTransport struct{ started chan struct{} }

func (contextTransport) Send(ctx context.Context, _, _, _, _ string) error {
	// The transport reports success if the caller fails to propagate cancellation.
	// That makes the context boundary observable without relying on timing in production code.
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(20 * time.Millisecond):
		return nil
	}
}

func (t contextTransport) SendStarted(ctx context.Context, _, _, _, _ string) error {
	close(t.started)
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(20 * time.Millisecond):
		return nil
	}
}

func TestSendWithRetryPassesContextToTransport(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	transport := contextTransport{started: make(chan struct{})}
	result := make(chan error, 1)
	go func() { result <- SendWithRetry(ctx, startedTransport{transport}, "a", "b", "s", "body", 2) }()
	<-transport.started
	cancel()
	err := <-result
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("err=%v, want context cancellation", err)
	}
}

type startedTransport struct{ contextTransport }

func (t startedTransport) Send(ctx context.Context, from, to, subject, body string) error {
	return t.contextTransport.SendStarted(ctx, from, to, subject, body)
}
