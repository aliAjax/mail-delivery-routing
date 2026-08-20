package infrastructure

import (
	"context"
	"fmt"
	"sync"
)

type MemoryTransport struct {
	mu       sync.Mutex
	Sent     []string
	FailNext bool
}

func (m *MemoryTransport) Send(_ context.Context, from, to, subject, body string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.FailNext {
		m.FailNext = false
		return fmt.Errorf("simulated temporary smtp failure")
	}
	m.Sent = append(m.Sent, from+" -> "+to+" | "+subject+" | "+body)
	return nil
}
