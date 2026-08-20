package config

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func LoadSimpleYAML(path string, c *Config) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		parts := strings.SplitN(strings.TrimSpace(s.Text()), ":", 2)
		if len(parts) != 2 {
			continue
		}
		k, v := strings.TrimSpace(parts[0]), strings.Trim(strings.TrimSpace(parts[1]), "\"'")
		switch k {
		case "http_addr":
			c.HTTPAddr = v
		case "smtp_addr":
			c.SMTPAddr = v
		case "grpc_addr":
			c.GRPCAddr = v
		case "log_level":
			c.LogLevel = v
		case "max_body":
			if n, e := strconv.ParseInt(v, 10, 64); e == nil {
				c.MaxBody = n
			}
		case "queue_workers":
			if n, e := strconv.Atoi(v); e == nil {
				c.QueueWorkers = n
			}
		}
	}
	return s.Err()
}
