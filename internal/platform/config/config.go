package config

import (
	"os"
	"strconv"
)

type Config struct {
	HTTPAddr     string
	SMTPAddr     string
	GRPCAddr     string
	LogLevel     string
	MaxBody      int64
	QueueWorkers int
}

func Load(path string) Config {
	c := Config{HTTPAddr: ":18089", SMTPAddr: ":2525", GRPCAddr: ":19089", LogLevel: "info", MaxBody: 1 << 20, QueueWorkers: 2}
	if v := os.Getenv("MAIL_HTTP_ADDR"); v != "" {
		c.HTTPAddr = v
	}
	if v := os.Getenv("MAIL_SMTP_ADDR"); v != "" {
		c.SMTPAddr = v
	}
	if v := os.Getenv("MAIL_GRPC_ADDR"); v != "" {
		c.GRPCAddr = v
	}
	if v := os.Getenv("MAIL_LOG_LEVEL"); v != "" {
		c.LogLevel = v
	}
	if v := os.Getenv("MAIL_MAX_BODY"); v != "" {
		if n, err := strconv.ParseInt(v, 10, 64); err == nil {
			c.MaxBody = n
		}
	}
	if v := os.Getenv("MAIL_QUEUE_WORKERS"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			c.QueueWorkers = n
		}
	}
	return c
}
