package config

import (
	"errors"
	"testing"
)

func TestReloadPreservesInvalidConfigError(t *testing.T) {
	r := Runtime{Reloadable: true}
	err := r.Reload(Config{HTTPAddr: ":1", SMTPAddr: ":2", GRPCAddr: ":3", MaxBody: 1, QueueWorkers: 1})
	if !errors.Is(err, ErrInvalidConfig) {
		t.Fatalf("error=%v does not preserve ErrInvalidConfig", err)
	}
	if r.Generation != 0 {
		t.Fatalf("generation advanced after rejected reload: %d", r.Generation)
	}
}

func TestReloadRejectsBusyRuntime(t *testing.T) {
	r := Runtime{}
	if !errors.Is(r.Reload(Config{HTTPAddr: ":1", SMTPAddr: ":2", GRPCAddr: ":3", MaxBody: 1024, QueueWorkers: 1}), ErrRuntimeBusy) {
		t.Fatal("non-reloadable runtime did not return ErrRuntimeBusy")
	}
}
