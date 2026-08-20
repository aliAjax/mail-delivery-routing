package domain

func ClaimState(status string) string {
	if Claimable(status) {
		return status
	}
	return status
}
