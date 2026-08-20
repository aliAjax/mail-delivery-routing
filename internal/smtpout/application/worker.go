package application

import (
	"context"
	messageapp "example.com/maildelivery/internal/message/application"
	queueapp "example.com/maildelivery/internal/queue/application"
	"example.com/maildelivery/internal/smtpout/domain"
	"log/slog"
	"time"
)

type Worker struct {
	Queue     *queueapp.Service
	Messages  *messageapp.Service
	Transport domain.Transport
	Log       *slog.Logger
}

func (w *Worker) Run(ctx context.Context) {
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			w.once(ctx)
		}
	}
}
func (w *Worker) once(ctx context.Context) {
	j, ok := w.Queue.Claim(ctx, time.Now())
	if !ok {
		return
	}
	m, err := w.Messages.Get(ctx, j.MessageID)
	if err == nil {
		err = w.Transport.Send(ctx, m.From, m.To, m.Subject, m.EffectiveBody())
	}
	w.Queue.Finish(ctx, j.ID, err)
	if err == nil {
		_ = w.Messages.UpdateStatus(ctx, m.ID, "delivered")
	} else {
		_ = w.Messages.UpdateStatus(ctx, m.ID, "retry")
	}
	if w.Log != nil {
		w.Log.Info("delivery attempt", "job_id", j.ID, "message_id", j.MessageID, "error", err)
	}
}
