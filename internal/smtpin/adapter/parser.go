package adapter

import (
	"bufio"
	"context"
	"io"
	"strings"
)

type Envelope struct {
	From       string
	Recipients []string
	Data       string
}

func Parse(r io.Reader) (Envelope, error) {
	s := bufio.NewScanner(r)
	var e Envelope
	var data strings.Builder
	inData := false
	for s.Scan() {
		line := s.Text()
		u := strings.ToUpper(line)
		if strings.HasPrefix(u, "MAIL FROM:") {
			e.From = strings.Trim(strings.TrimSpace(line[10:]), "<>")
		}
		if strings.HasPrefix(u, "RCPT TO:") {
			e.Recipients = append(e.Recipients, strings.Trim(strings.TrimSpace(line[8:]), "<>"))
		}
		if u == "DATA" {
			inData = true
			continue
		}
		if inData && line == "." {
			inData = false
			continue
		}
		if inData {
			data.WriteString(line)
			data.WriteByte('\n')
		}
	}
	e.Data = data.String()
	return e, s.Err()
}

func ParseContext(ctx context.Context, r io.Reader) (Envelope, error) {
	s := bufio.NewScanner(r)
	var e Envelope
	var data strings.Builder
	inData := false
	for s.Scan() {
		select {
		case <-ctx.Done():
			return Envelope{}, ctx.Err()
		default:
		}
		line := s.Text()
		u := strings.ToUpper(line)
		if strings.HasPrefix(u, "MAIL FROM:") {
			e.From = strings.Trim(strings.TrimSpace(line[10:]), "<>")
		}
		if strings.HasPrefix(u, "RCPT TO:") {
			e.Recipients = append(e.Recipients, strings.Trim(strings.TrimSpace(line[8:]), "<>"))
		}
		if u == "DATA" {
			inData = true
			continue
		}
		if inData && line == "." {
			inData = false
			continue
		}
		if inData {
			data.WriteString(line)
			data.WriteByte('\n')
		}
	}
	if err := s.Err(); err != nil {
		return Envelope{}, err
	}
	select {
	case <-ctx.Done():
		return Envelope{}, ctx.Err()
	default:
	}
	e.Data = data.String()
	return e, nil
}
