package domain

import "strings"

type Kind string

const (
	Soft      Kind = "soft"
	Hard      Kind = "hard"
	Complaint Kind = "complaint"
)

type Event struct {
	MessageID, Address, Reason string
	Kind                       Kind
}

func Classify(code int, reason string) Kind {
	if code >= 500 {
		return Hard
	}
	if strings.Contains(strings.ToLower(reason), "complaint") {
		return Complaint
	}
	return Soft
}
