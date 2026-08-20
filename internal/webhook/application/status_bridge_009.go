package application

import "example.com/maildelivery/internal/webhook/domain"

func StatusFailure(code int) error {
	return domain.HTTPStatusError{Code: code, Retryable: domain.RetryableStatus(code)}
}
