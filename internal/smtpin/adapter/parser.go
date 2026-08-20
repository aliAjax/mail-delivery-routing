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
	// Terminated reports whether the DATA section was closed by a line
	// containing a single dot, i.e. the envelope was fully received.
	Terminated bool
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
			e.Terminated = true
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
	for {
		// Honour cancellation before each read so a cancelled or timed-out
		// inbound stream stops promptly instead of draining the body.
		if err := ctx.Err(); err != nil {
			return Envelope{}, err
		}
		if !s.Scan() {
			break
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
			e.Terminated = true
			continue
		}
		if inData {
			data.WriteString(line)
			data.WriteByte('\n')
		}
	}
	// A reader may surface the cancellation via EOF (or an error) on the
	// final read, so re-check the context before treating the scan as done.
	if err := ctx.Err(); err != nil {
		return Envelope{}, err
	}
	if err := s.Err(); err != nil {
		return Envelope{}, err
	}
	e.Data = data.String()
	return e, nil
}
