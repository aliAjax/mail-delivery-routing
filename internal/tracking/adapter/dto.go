package adapter

import (
	"example.com/maildelivery/internal/tracking/domain"
	"time"
)

type DTO struct {
	MessageID, TenantID, Status, Reason string
	At                                  time.Time
	Attempt                             int
}

func From(e domain.Event) DTO {
	return DTO{MessageID: e.MessageID, TenantID: e.TenantID, Status: e.Status, Reason: e.Reason, At: e.At, Attempt: e.Attempt}
}
