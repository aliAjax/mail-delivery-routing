package domain

func ClaimState(status string) string {
	if Claimable(status) {
		return "running"
	}
	return status
}
