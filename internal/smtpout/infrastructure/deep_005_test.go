package infrastructure

import (
	"context"
	"example.com/maildelivery/internal/smtpout/domain"
	"testing"
)

func TestRetryDeepClaims(t *testing.T) {
	if got := domain.RetryBudget(0); got != 1 {
		t.Fatalf("budget=%d", got)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if RetryContext(ctx).Err() == nil {
		t.Fatal("retry context lost cancellation")
	}
}
