package store

import "context"

type Tx interface {
	Commit(context.Context) error
	Rollback(context.Context) error
}
type NoopTx struct{}

func (NoopTx) Commit(context.Context) error                            { return nil }
func (NoopTx) Rollback(context.Context) error                          { return nil }
func WithTx(ctx context.Context, fn func(context.Context) error) error { return fn(ctx) }
