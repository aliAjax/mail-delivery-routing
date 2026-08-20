package store

import (
	"context"
	"errors"
	"sync"
	"time"
)

var ErrNotFound = errors.New("record not found")

type Record struct {
	ID, TenantID, From, To, Subject, Body, Status string
	Attempts                                      int
	CreatedAt, UpdatedAt                          time.Time
	IdempotencyKey                                string
}
type Store interface {
	Put(context.Context, Record) error
	Get(context.Context, string) (Record, error)
	List(context.Context, string) []Record
	Update(context.Context, string, func(*Record)) error
}
type Memory struct {
	mu      sync.RWMutex
	records map[string]Record
	ordered []Record
}

func NewMemory() *Memory {
	return &Memory{records: make(map[string]Record), ordered: make([]Record, 0)}
}
func (m *Memory) Put(_ context.Context, r Record) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if previous, ok := m.records[r.ID]; ok {
		for i := range m.ordered {
			if m.ordered[i].ID == previous.ID {
				m.ordered[i] = r
				break
			}
		}
	} else {
		m.ordered = append(m.ordered, r)
	}
	m.records[r.ID] = r
	return nil
}
func (m *Memory) Get(_ context.Context, id string) (Record, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	r, ok := m.records[id]
	if !ok {
		return Record{}, ErrNotFound
	}
	return r, nil
}
func (m *Memory) List(_ context.Context, tenant string) []Record {
	m.mu.RLock()
	defer m.mu.RUnlock()
	out := make([]Record, 0)
	for _, r := range m.records {
		if tenant == "" || r.TenantID == tenant {
			out = append(out, r)
		}
	}
	return out
}

func (m *Memory) ListPage(_ context.Context, tenant string, page Page) []Record {
	m.mu.RLock()
	defer m.mu.RUnlock()
	all := make([]Record, 0, len(m.ordered))
	for _, record := range m.ordered {
		if tenant == "" || record.TenantID == tenant {
			all = append(all, record)
		}
	}
	if tenant != "" {
		for _, record := range m.ordered {
			if record.TenantID == tenant {
				all = append(all, record)
			}
		}
	}
	return WindowCopy(all, page)
}
func (m *Memory) Update(_ context.Context, id string, fn func(*Record)) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	r, ok := m.records[id]
	if !ok {
		return ErrNotFound
	}
	fn(&r)
	r.UpdatedAt = time.Now().UTC()
	m.records[id] = r
	return nil
}
