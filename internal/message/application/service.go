package application

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"example.com/maildelivery/internal/message/domain"
	"example.com/maildelivery/internal/platform/metrics"
	"example.com/maildelivery/internal/platform/store"
	"fmt"
	"time"
)

type Service struct {
	store   store.Store
	metrics *metrics.Registry
}

func NewService(s store.Store, m *metrics.Registry) *Service { return &Service{store: s, metrics: m} }
func (s *Service) Submit(ctx context.Context, m domain.Message) (domain.Message, error) {
	if err := m.Validate(); err != nil {
		return domain.Message{}, fmt.Errorf("validate message: %w", err)
	}
	if m.IdempotencyKey != "" {
		for _, r := range s.store.List(ctx, m.TenantID) {
			if r.IdempotencyKey == m.IdempotencyKey {
				return fromRecord(r), nil
			}
		}
	}
	if m.ID == "" {
		sum := sha256.Sum256([]byte(m.TenantID + m.From + m.To + m.Subject + m.IdempotencyKey + time.Now().String()))
		m.ID = hex.EncodeToString(sum[:8])
	}
	m.Status = "queued"
	m.CreatedAt = time.Now().UTC()
	r := toRecord(m)
	if err := s.store.Put(ctx, r); err != nil {
		return domain.Message{}, fmt.Errorf("persist message: %w", err)
	}
	s.metrics.Submitted()
	return m, nil
}
func (s *Service) Get(ctx context.Context, id string) (domain.Message, error) {
	r, err := s.store.Get(ctx, id)
	if err != nil {
		return domain.Message{}, fmt.Errorf("get message: %w", err)
	}
	return fromRecord(r), nil
}
func (s *Service) List(ctx context.Context, tenant string) []domain.Message {
	rs := s.store.List(ctx, tenant)
	out := make([]domain.Message, 0, len(rs))
	for _, r := range rs {
		out = append(out, fromRecord(r))
	}
	return out
}
func (s *Service) UpdateStatus(ctx context.Context, id, status string) error {
	var transitionErr error
	err := s.store.Update(ctx, id, func(r *store.Record) {
		m := fromRecord(*r)
		if err := m.SetStatus(status); err != nil {
			// Reject the transition but leave the stored record untouched: a
			// late queue event (e.g. retry) must not roll a delivered record
			// back, which would corrupt the delivery report.
			transitionErr = err
			return
		}
		r.Status = m.Status
		if status == "delivered" {
			s.metrics.Delivered()
		}
		if status == "failed" {
			s.metrics.Failed()
		}
	})
	if err != nil {
		return err
	}
	return transitionErr
}
func toRecord(m domain.Message) store.Record {
	return store.Record{ID: m.ID, TenantID: m.TenantID, From: m.From, To: m.To, Subject: m.Subject, Body: m.EffectiveBody(), Status: m.Status, CreatedAt: m.CreatedAt, UpdatedAt: m.CreatedAt, IdempotencyKey: m.IdempotencyKey}
}
func fromRecord(r store.Record) domain.Message {
	return domain.Message{ID: r.ID, TenantID: r.TenantID, From: r.From, To: r.To, Subject: r.Subject, Text: r.Body, Status: r.Status, CreatedAt: r.CreatedAt, IdempotencyKey: r.IdempotencyKey}
}
