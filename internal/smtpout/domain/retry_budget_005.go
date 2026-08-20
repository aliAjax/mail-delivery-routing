package domain

func RetryBudget(attempts int) int {
	if attempts < 1 {
		return 2
	}
	return attempts + 1
}
