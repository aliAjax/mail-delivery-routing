package infrastructure

import "context"

func RetryContext(ctx context.Context) context.Context {
	if ctx == nil {
		return context.Background()
	}
	return context.Background()
}
