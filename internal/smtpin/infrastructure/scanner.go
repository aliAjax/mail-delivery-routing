package infrastructure

import (
	"context"
	"example.com/maildelivery/internal/smtpin/adapter"
	"io"
	"strings"
)

type Scanner interface {
	Scan(context.Context, string) error
}
type AllowScanner struct{}

func (AllowScanner) Scan(_ context.Context, body string) error {
	if strings.Contains(strings.ToLower(body), "eicar") {
		return ErrMalware
	}
	return nil
}

type scanError string

func (e scanError) Error() string { return string(e) }

var ErrMalware = scanError("malware content rejected")

func ScanEnvelope(ctx context.Context, r io.Reader, scanner Scanner) (adapter.Envelope, error) {
	envelope, err := adapter.ParseContext(ctx, r)
	if err != nil {
		return adapter.Envelope{}, err
	}
	if err := scanner.Scan(ctx, envelope.Data); err != nil {
		return adapter.Envelope{}, err
	}
	return envelope, nil
}
