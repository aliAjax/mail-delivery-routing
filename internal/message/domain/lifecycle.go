package domain

import (
	"errors"
	"fmt"
	"time"
)

var Transitions = map[string]map[string]bool{"queued": {"running": true, "cancelled": true, "suppressed": true}, "running": {"delivered": true, "retry": true, "failed": true}, "retry": {"running": true, "failed": true}, "delivered": {}, "failed": {}, "cancelled": {}, "suppressed": {}}

func CanTransition(from, to string) bool { next, ok := Transitions[from]; return ok && next[to] }

var ErrInvalidTransition = errors.New("invalid message transition")

func ApplyTransition(m *Message, to string) error {
	if m == nil {
		return ErrInvalidTransition
	}
	if !CanTransition(m.Status, to) {
		return fmt.Errorf("%w: %s -> %s", ErrInvalidTransition, m.Status, to)
	}
	m.Status = to
	Stamp(m)
	return nil
}
func Stamp(m *Message) {
	if m.CreatedAt.IsZero() {
		m.CreatedAt = time.Now().UTC()
	}
}
func (m Message) IsTerminal() bool {
	return m.Status == "delivered" || m.Status == "failed" || m.Status == "cancelled" || m.Status == "suppressed"
}
