package domain

import "time"

type Job struct {
	ID, MessageID, Kind, Status string
	Attempts                    int
	NextAttempt                 time.Time
	LastError                   string
}

func (j *Job) CanRun(now time.Time) bool {
	return j.Status == "queued" || j.Status == "retry" && !j.NextAttempt.After(now)
}
