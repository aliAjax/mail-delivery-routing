package domain

// RetryBudget turns a configured attempt count into the total number of
// delivery attempts that SendWithRetry may perform. A zero/unset value means
// "exactly one attempt, no retries" rather than an implicit retry budget.
func RetryBudget(attempts int) int {
	if attempts < 1 {
		return 1
	}
	return attempts
}
