package domain

import (
	"errors"
	"testing"
)

func TestMessageTransitionRejectsTerminalRollback(t *testing.T) {
	m := Message{Status: "delivered"}
	if err := m.SetStatus("retry"); !errors.Is(err, ErrInvalidTransition) {
		t.Fatalf("err=%v, want invalid transition", err)
	}
	if m.Status != "delivered" {
		t.Fatalf("terminal message changed to %q", m.Status)
	}
}

func TestMessageTransitionAcceptsRetryPath(t *testing.T) {
	m := Message{Status: "queued"}
	for _, status := range []string{"running", "retry", "running", "delivered"} {
		if err := m.SetStatus(status); err != nil {
			t.Fatalf("transition to %s: %v", status, err)
		}
	}
	if !m.IsTerminal() {
		t.Fatal("delivered message is not terminal")
	}
}
