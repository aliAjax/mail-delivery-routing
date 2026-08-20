package domain

import (
	"context"
	"errors"
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
