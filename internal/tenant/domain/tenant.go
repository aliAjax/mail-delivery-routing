package domain

import "errors"

type Tenant struct {
	ID, Name         string
	DailyQuota, Used int
	Paused           bool
}

var ErrQuota = errors.New("tenant quota exceeded")

func (t *Tenant) Consume() error {
	if t.Paused {
		return errors.New("tenant paused")
	}
	if t.DailyQuota > 0 && t.Used >= t.DailyQuota {
		return ErrQuota
	}
	t.Used++
	return nil
}

func (t Tenant) Remaining() int {
	if t.DailyQuota <= 0 {
		return -1
	}
	remaining := t.DailyQuota - t.Used
	if remaining < 0 {
		return 0
	}
	return remaining
}

func (t Tenant) CanSend() bool {
	return !t.Paused && (t.DailyQuota <= 0 || t.Used < t.DailyQuota)
}
