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
	// Keep the caller's cancellation wired through so a cancelled send stops the
	// in-flight attempt and skips remaining retries, and derive the attempt
	// budget from the canonical rule rather than a local off-by calc.
	return domain.SendWithRetry(RetryContext(ctx), r.Transport, from, to, subject, body, domain.RetryBudget(r.Attempts))
}
