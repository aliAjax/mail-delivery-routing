package infrastructure

import (
	"context"
	"errors"
	"example.com/maildelivery/internal/smtpin/adapter"
	"strings"
	"testing"
)

type testScanner struct{}

func (testScanner) Scan(_ context.Context, body string) error {
	if !strings.Contains(body, "safe") {
		return errors.New("unexpected body")
	}
	return nil
}

func TestScanEnvelopePropagatesParserContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := ScanEnvelope(ctx, strings.NewReader("DATA\nsafe\n.\n"), testScanner{})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("err=%v, want context canceled", err)
	}
}

func TestScanEnvelopeRunsScannerAfterParsing(t *testing.T) {
	envelope, err := ScanEnvelope(context.Background(), strings.NewReader("MAIL FROM:<a@test>\nDATA\nsafe\n.\n"), testScanner{})
	if err != nil || envelope.From != "a@test" {
		t.Fatalf("envelope=%+v, err=%v", envelope, err)
	}
}

var _ adapter.Envelope
