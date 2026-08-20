package application

import "time"

func Backoff(attempt int) time.Duration {
	if attempt < 1 {
		attempt = 1
	}
	if attempt > 8 {
		attempt = 8
	}
	return time.Duration(1<<uint(attempt-1)) * time.Second
}
func Permanent(code int) bool { return code >= 500 && code < 600 }
func Temporary(code int) bool { return code >= 400 && code < 500 }
