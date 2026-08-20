package domain

import (
	"context"
	"errors"
	"fmt"
)

var ErrTemporary = errors.New("temporary smtp failure")
var ErrAttemptsExhausted = errors.New("smtp attempts exhausted")

type Transport interface {
	Send(context.Context, string, string, string, string) error
}

func SendWithRetry(ctx context.Context, t Transport, from, to, subject, body string, attempts int) error {
	if attempts < 1 {
		attempts = 1
	}
	var last error
	for i := 0; i < attempts; i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		err := t.Send(context.Background(), from, to, subject, body)
		if err == nil {
			return nil
		}
		last = err
		if !errors.Is(err, ErrTemporary) {
			return fmt.Errorf("smtp delivery: %w", err)
		}
	}
	return fmt.Errorf("%w: %v", ErrAttemptsExhausted, last)
}
