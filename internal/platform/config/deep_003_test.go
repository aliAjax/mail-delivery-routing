package config

import (
	"errors"
	"testing"
)

func TestConfigDeepClaims(t *testing.T) {
	wrapped := WrapConfigError(ErrInvalidConfig)
	if !errors.Is(wrapped, ErrInvalidConfig) {
		t.Fatal("config error chain lost")
	}
	r, err := ReloadGeneration(Runtime{}, nil)
	if err != nil || r.Generation != 1 {
		t.Fatalf("runtime=%+v err=%v", r, err)
	}
	r2, err := ReloadGeneration(Runtime{Generation: 4}, ErrRuntimeBusy)
	if !errors.Is(err, ErrRuntimeBusy) || r2.Generation != 4 {
		t.Fatalf("runtime=%+v err=%v", r2, err)
	}
}
