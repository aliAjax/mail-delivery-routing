package domain

var Terminal = map[string]bool{"delivered": true, "failed": true, "suppressed": true}

func CanTransition(from, to string) bool {
	if Terminal[from] {
		return false
	}
	switch from {
	case "queued":
		return to == "running" || to == "cancelled"
	case "running":
		return to == "delivered" || to == "retry" || to == "failed"
	case "retry":
		return to == "running" || to == "failed"
	}
	return false
}
