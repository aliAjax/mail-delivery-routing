package main

import (
	"context"
	"errors"
	"example.com/maildelivery/api/grpc"
	httpapi "example.com/maildelivery/api/http"
	"example.com/maildelivery/internal/message/application"
	"example.com/maildelivery/internal/platform/config"
	"example.com/maildelivery/internal/platform/logger"
	"example.com/maildelivery/internal/platform/metrics"
	"example.com/maildelivery/internal/platform/store"
	queueapp "example.com/maildelivery/internal/queue/application"
	smtpinapp "example.com/maildelivery/internal/smtpin/application"
	smtpoutapp "example.com/maildelivery/internal/smtpout/application"
	smtpoutinfra "example.com/maildelivery/internal/smtpout/infrastructure"
	"log/slog"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.Load("configs/config.yaml")
	log := logger.New(cfg.LogLevel)
	reg := &metrics.Registry{}
	st := store.NewMemory()
	messages := application.NewService(st, reg)
	queue := queueapp.New()
	transport := &smtpoutinfra.MemoryTransport{}
	worker := &smtpoutapp.Worker{Queue: queue, Messages: messages, Transport: transport, Log: log}
	inbound := &smtpinapp.Server{Addr: cfg.SMTPAddr, Messages: messages}
	api := httpapi.New(cfg.HTTPAddr, messages, queue, reg, log)
	g := &grpcapi.Server{Addr: cfg.GRPCAddr, Log: log}
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	go worker.Run(ctx)
	go func() {
		if err := inbound.Run(ctx); err != nil && !errors.Is(err, context.Canceled) {
			log.Error("smtp stopped", "error", err)
		}
	}()
	go func() {
		if err := g.Listen(ctx); err != nil && !errors.Is(err, context.Canceled) {
			log.Error("grpc stopped", "error", err)
		}
	}()
	go func() {
		if err := api.HTTP.ListenAndServe(); err != nil && !errors.Is(err, context.Canceled) {
			log.Error("http stopped", "error", err)
		}
	}()
	<-ctx.Done()
	shutdown(api, log)
}
func shutdown(api *httpapi.Server, log *slog.Logger) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := api.HTTP.Shutdown(ctx); err != nil {
		log.Error("shutdown", "error", err)
	}
	log.Info("mail delivery stopped")
}
