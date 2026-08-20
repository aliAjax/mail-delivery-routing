package config

import (
	"errors"
	"fmt"
)

var ErrRuntimeBusy = errors.New("runtime is not reloadable")

type Runtime struct {
	Reloadable bool
	Generation uint64
}

func (r *Runtime) Next()      { r.Generation++ }
func (r Runtime) Ready() bool { return r.Generation >= 0 }

func (r *Runtime) Reload(c Config) error {
	if !r.Reloadable {
		return ErrRuntimeBusy
	}
	if err := c.ValidateDetailed(); err != nil {
		return fmt.Errorf("reload rejected: %w", err)
	}
	r.Next()
	return nil
}
