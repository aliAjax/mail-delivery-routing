package domain

var Allowed = map[string]map[string]bool{"queued": {"running": true, "cancelled": true}, "running": {"done": true, "retry": true, "dead": true}, "retry": {"running": true, "dead": true}, "done": {}, "dead": {}, "cancelled": {}}

func Transition(from, to string) bool { next, ok := Allowed[from]; return ok && next[to] }

func Claimable(status string) bool { return status == "queued" || status == "retry" }

func Terminal(status string) bool {
	return status == "done" || status == "dead" || status == "cancelled"
}
