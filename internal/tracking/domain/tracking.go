package domain

import "time"

type Event struct {
	MessageID, TenantID, Status, Reason string
	At                                  time.Time
	Attempt                             int
}
type Filter struct {
	TenantID, Status string
	From, To         time.Time
}

func (e Event) Valid() bool { return e.MessageID != "" && e.Status != "" && !e.At.IsZero() }
