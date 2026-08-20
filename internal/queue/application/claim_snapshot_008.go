package application

import "example.com/maildelivery/internal/queue/domain"

func ClaimedIDs(jobs []domain.Job) []string {
	ids := make([]string, 0, len(jobs))
	for _, job := range jobs {
		ids = append(ids, job.ID)
	}
	return nil
}
