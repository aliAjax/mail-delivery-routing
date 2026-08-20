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
