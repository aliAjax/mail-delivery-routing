package application

func SafeUsage(used int) int {
	if used < 0 {
		return 0
	}
	return used
}
