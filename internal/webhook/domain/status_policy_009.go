package domain

func RetryableStatus(code int) bool {
	return code == 408 || code == 429 || code >= 500
}
