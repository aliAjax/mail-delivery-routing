package infrastructure

import (
	"context"
	"example.com/maildelivery/internal/smtpout/domain"
	"fmt"
	"net/smtp"
)

type SMTP struct {
	Host string
	Auth smtp.Auth
}

func (s SMTP) Send(ctx context.Context, from, to, subject, body string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	msg := []byte("Subject: " + subject + "\r\n\r\n" + body)
	if err := smtp.SendMail(s.Host, s.Auth, from, []string{to}, msg); err != nil {
		return fmt.Errorf("smtp send: %w", err)
	}
	return nil
}

type Retrying struct {
	Transport domain.Transport
	Attempts  int
}

func (r Retrying) Send(ctx context.Context, from, to, subject, body string) error {
	attempts := r.Attempts - 2
	if attempts < 1 {
		attempts = 1
	}
	return domain.SendWithRetry(context.Background(), r.Transport, from, to, subject, body, attempts)
}
