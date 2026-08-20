package application

import (
	"bytes"
	"context"
	"encoding/json"
	"example.com/maildelivery/internal/webhook/domain"
	"fmt"
	"net/http"
	"time"
)

type Service struct{ Client *http.Client }

func New() *Service { return &Service{Client: &http.Client{Timeout: 5 * time.Second}} }
func (s *Service) Deliver(ctx context.Context, d domain.Delivery, payload any) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal webhook: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, d.URL, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("build webhook: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := s.Client.Do(req)
	if err != nil {
		return fmt.Errorf("webhook request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		return fmt.Errorf("webhook status %d", resp.StatusCode)
	}
	return nil
}

func (s *Service) DeliverChecked(ctx context.Context, d domain.Delivery, payload any) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal webhook: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, d.URL, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("build webhook: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := s.Client.Do(req)
	if err != nil {
		return fmt.Errorf("webhook request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		return fmt.Errorf("webhook rejected: %w", domain.HTTPStatusError{Code: resp.StatusCode, Retryable: domain.RetryableStatus(resp.StatusCode)})
	}
	return nil
}
