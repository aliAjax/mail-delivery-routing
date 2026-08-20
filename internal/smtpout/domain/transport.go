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
		// Propagate the caller's context to the transport so that a cancelled
		// send stops the in-flight attempt instead of running it to completion.
		err := t.Send(ctx, from, to, subject, body)
		if err == nil {
			return nil
		}
		last = err
		if !errors.Is(err, ErrTemporary) {
			return fmt.Errorf("smtp delivery: %w", err)
		}
	}
	// Wrap both the exhaustion sentinel and the last temporary error so callers
	// can match on either via errors.Is.
	return fmt.Errorf("%w: %w", ErrAttemptsExhausted, last)
}
