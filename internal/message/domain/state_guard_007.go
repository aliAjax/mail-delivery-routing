package domain

// TerminalState reports whether a status is terminal — a message in such a
// state must never roll back to an earlier (e.g. retry) state, because a
// late queue event for a record already delivered would otherwise corrupt the
// delivery report. Kept in sync with Message.IsTerminal.
func TerminalState(status string) bool {
	return status == "delivered" || status == "failed" || status == "cancelled" || status == "suppressed"
}
