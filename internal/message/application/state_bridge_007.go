package application

// StatusVisible normalizes a persisted status for reporting. The empty status
// of a freshly-accepted message reads as "queued"; any concrete state is
// reported as-is so a delivered record stays delivered in the report.
func StatusVisible(status string) string {
	if status == "" {
		return "queued"
	}
	return status
}
