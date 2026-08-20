package domain

func TerminalState(status string) bool {
	return status == "delivered" || status == "failed" || status == "cancelled" || status == "suppressed"
}
