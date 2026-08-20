package application

func ClampConcurrent(v int) int {
	if v < 1 {
		return 1
	}
	if v > 100 {
		return 100
	}
	return v
}
func ClampQuota(v int) int {
	if v < 0 {
		return 0
	}
	if v > 1000000 {
		return 1000000
	}
	return v
}
