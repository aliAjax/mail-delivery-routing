package infrastructure

import "context"

// RetryContext returns the context the retry loop should use. It must keep
// the caller's cancellation channel wired through so a cancelled send stops
// both the in-flight attempt and any remaining retries; only a nil context
// is replaced with a fresh Background.
func RetryContext(ctx context.Context) context.Context {
	if ctx == nil {
		return context.Background()
	}
	return ctx
}
