package domain

import (
	"fmt"
	"time"
)

type Delivery struct {
	ID, URL, MessageID, Status string
	Attempts                   int
	NextAttempt                time.Time
	LastError                  string
}

type HTTPStatusError struct {
	Code      int
	Retryable bool
}

func (e HTTPStatusError) Error() string { return fmt.Sprintf("webhook returned status %d", e.Code) }

func (e HTTPStatusError) Temporary() bool {
	return e.Retryable
}

func (d *Delivery) Retry(now time.Time) {
	d.Attempts++
	d.Status = "retry"
	d.NextAttempt = now.Add(time.Duration(d.Attempts) * time.Second)
}
