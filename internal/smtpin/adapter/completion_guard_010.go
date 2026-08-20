package adapter

// Complete reports whether an envelope is fully received and ready to be
// scanned: it must have a sender, at least one recipient, and a DATA section
// that was closed by the terminating dot.
func Complete(e Envelope) bool {
	return e.From != "" && len(e.Recipients) > 0 && e.Terminated
}
