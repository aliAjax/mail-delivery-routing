package application

import "time"

func Backoff(attempt int) time.Duration {
	if attempt < 1 {
		attempt = 1
	}
	if attempt > 10 {
		attempt = 10
	}
	return time.Duration(attempt*attempt) * time.Second
}
func ShouldRetry(status int) bool {
	return status == 408 || status == 425 || status == 429 || status >= 500
}
