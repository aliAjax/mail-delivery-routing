package domain

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var ErrInvalid = errors.New("invalid message")

type Message struct {
	ID, TenantID, From, To, ReplyTo, Subject, Text, HTML, IdempotencyKey string
	Headers                                                              map[string]string
	Status                                                               string
	CreatedAt                                                            time.Time
}

func (m Message) Validate() error {
	if m.TenantID == "" || m.From == "" || m.To == "" || !strings.Contains(m.From, "@") || !strings.Contains(m.To, "@") {
		return ErrInvalid
	}
	if strings.ContainsAny(m.From+m.To+m.ReplyTo, "\r\n") {
		return errors.New("header injection")
	}
	if len(m.Subject) > 998 || len(m.Text)+len(m.HTML) > 1<<20 {
		return errors.New("message exceeds limits")
	}
	return nil
}
func (m Message) EffectiveBody() string {
	if m.HTML != "" {
		return m.HTML
	}
	return m.Text
}

func (m *Message) SetStatus(status string) error {
	if m == nil {
		return fmt.Errorf("%w: nil message", ErrInvalid)
	}
	if !CanTransition(m.Status, status) {
		return fmt.Errorf("%w: %s -> %s", ErrInvalidTransition, m.Status, status)
	}
	return ApplyTransition(m, status)
}
