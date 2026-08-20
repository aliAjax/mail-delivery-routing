package domain

func RetryBudget(attempts int) int {
	if attempts < 1 {
		return 1
	}
	return attempts
}
