package config

import (
	"errors"
	"fmt"
)

var ErrInvalidConfig = errors.New("invalid configuration")

func (c Config) Validate() error {
	if c.HTTPAddr == "" || c.SMTPAddr == "" || c.GRPCAddr == "" {
		return fmt.Errorf("all listener addresses are required")
	}
	if c.MaxBody < 1024 {
		return fmt.Errorf("max body must be at least 1024")
	}
	if c.QueueWorkers < 1 {
		return fmt.Errorf("queue workers must be positive")
	}
	return nil
}
func (c Config) Public() map[string]any {
	return map[string]any{"http_addr": c.HTTPAddr, "smtp_addr": c.SMTPAddr, "grpc_addr": c.GRPCAddr, "max_body": c.MaxBody, "queue_workers": c.QueueWorkers}
}

func (c Config) ValidateDetailed() error {
	if err := c.Validate(); err != nil {
		return fmt.Errorf("%w: %w", ErrInvalidConfig, err)
	}
	return nil
}
