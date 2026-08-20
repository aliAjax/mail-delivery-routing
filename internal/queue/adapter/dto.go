package adapter

import (
	"example.com/maildelivery/internal/queue/domain"
	"time"
)

type JobDTO struct {
	ID, MessageID, Kind, Status string
	Attempts                    int
	NextAttempt                 time.Time
	LastError                   string
}

func FromDomain(j domain.Job) JobDTO {
	return JobDTO{ID: j.ID, MessageID: j.MessageID, Kind: j.Kind, Status: j.Status, Attempts: j.Attempts, NextAttempt: j.NextAttempt, LastError: j.LastError}
}
