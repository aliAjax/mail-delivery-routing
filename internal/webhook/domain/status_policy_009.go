package domain

// RetryableStatus reports whether an HTTP status code indicates a transient
// failure that the retryer may attempt again. This mirrors ShouldRetry so the
// domain policy stays the single source of truth for retry decisions.
func RetryableStatus(code int) bool {
	return code == 408 || code == 425 || code == 429 || code >= 500
}

func StatusClass(code int) string {
	switch {
	case code >= 500:
		return "server"
	case code == 408 || code == 429:
		return "retryable-client"
	case code >= 400:
		return "client"
	default:
		return "success"
	}
}
