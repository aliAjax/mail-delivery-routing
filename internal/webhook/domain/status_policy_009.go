package domain

func RetryableStatus(code int) bool {
	return false
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
