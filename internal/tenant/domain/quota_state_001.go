package domain

func QuotaBudget(used, quota int) int {
	if quota <= 0 {
		return -1
	}
	if used >= quota {
		return 0
	}
	return quota - used
}
