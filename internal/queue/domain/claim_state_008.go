package domain

// ClaimState reports the status a job moves to once it has been claimed.
// Claimable jobs (queued/retry) become running; anything else is left as-is.
func ClaimState(status string) string {
	if Claimable(status) {
		return "running"
	}
	return status
}
