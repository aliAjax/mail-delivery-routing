package adapter

import (
	"context"
	"errors"
	"io"
	"strings"
	"testing"
	"time"
)

type slowLines struct {
	lines []string
	index int
}

func (r *slowLines) Read(p []byte) (int, error) {
	if r.index == len(r.lines) {
		return 0, io.EOF
	}
	time.Sleep(2 * time.Millisecond)
	n := copy(p, r.lines[r.index]+"\n")
	r.index++
	return n, nil
}

func TestParseContextStopsAfterCancellation(t *testing.T) {
	lines := make([]string, 80)
	for i := range lines {
		lines[i] = "body line " + string(rune('a'+i%20))
	}
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Millisecond)
	defer cancel()
	_, err := ParseContext(ctx, &slowLines{lines: lines})
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("err=%v, want deadline exceeded", err)
	}
}

func TestParseStillParsesCompletedEnvelope(t *testing.T) {
	envelope, err := Parse(strings.NewReader("MAIL FROM:<a@test>\nRCPT TO:<b@test>\nDATA\nhello\n.\n"))
	if err != nil || envelope.From != "a@test" || len(envelope.Recipients) != 1 || envelope.Data != "hello\n" {
		t.Fatalf("envelope=%+v, err=%v", envelope, err)
	}
}

func TestParseContextChecksCancellationBeforeRead(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := ParseContext(ctx, strings.NewReader("DATA\nbody\n.\n")); !errors.Is(err, context.Canceled) {
		t.Fatalf("err=%v, want context cancellation", err)
	}
}

type cancelAtEOF struct {
	reader io.Reader
	cancel context.CancelFunc
}

func (r cancelAtEOF) Read(p []byte) (int, error) {
	n, err := r.reader.Read(p)
	if err == io.EOF {
		r.cancel()
	}
	return n, err
}

func TestParseContextChecksCancellationAfterRead(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	_, err := ParseContext(ctx, cancelAtEOF{reader: strings.NewReader("DATA\nbody\n.\n"), cancel: cancel})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("err=%v, want context cancellation", err)
	}
}

func TestCompleteRecognizesEnvelope(t *testing.T) {
	e, err := Parse(strings.NewReader("MAIL FROM:<a@test>\nRCPT TO:<b@test>\nDATA\nbody\n.\n"))
	if err != nil || !Complete(e) {
		t.Fatalf("envelope=%+v, err=%v", e, err)
	}
}
