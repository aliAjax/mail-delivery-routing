package application

import (
	"context"
	"example.com/maildelivery/internal/queue/domain"
	"fmt"
	"sync"
	"time"
)

type Service struct {
	mu   sync.Mutex
	jobs map[string]domain.Job
	wake chan struct{}
}

func New() *Service { return &Service{jobs: map[string]domain.Job{}, wake: make(chan struct{}, 1)} }
func (s *Service) Enqueue(_ context.Context, j domain.Job) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if j.ID == "" {
		return fmt.Errorf("job id required")
	}
	s.jobs[j.ID] = j
	select {
	case s.wake <- struct{}{}:
	default:
	}
	return nil
}
func (s *Service) Claim(_ context.Context, now time.Time) (domain.Job, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for id, j := range s.jobs {
		if j.CanRun(now) {
			j.Status = "running"
			j.Attempts++
			s.jobs[id] = j
			return j, true
		}
	}
	return domain.Job{}, false
}
func (s *Service) Finish(_ context.Context, id string, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	j, ok := s.jobs[id]
	if !ok {
		return
	}
	if err == nil {
		j.Status = "done"
	} else {
		j.Status = "retry"
		j.LastError = err.Error()
		j.NextAttempt = time.Now().Add(time.Duration(j.Attempts) * time.Second)
	}
	s.jobs[id] = j
}
func (s *Service) List(_ context.Context) []domain.Job {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]domain.Job, 0, len(s.jobs))
	for _, j := range s.jobs {
		out = append(out, j)
	}
	return out
}

func (s *Service) ClaimBatch(_ context.Context, now time.Time, limit int) []domain.Job {
	if limit < 1 {
		return nil
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	claimed := make([]domain.Job, 0, limit)
	for id, job := range s.jobs {
		if len(claimed) == limit {
			break
		}
		if !domain.Claimable(job.Status) || !job.CanRun(now) {
			continue
		}
		job.Status = "running"
		job.Attempts++
		s.jobs[id] = job
		claimed = append(claimed, job)
	}
	return claimed
}
func (s *Service) Wake() <-chan struct{} { return s.wake }
