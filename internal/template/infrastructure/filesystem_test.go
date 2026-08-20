package infrastructure

import (
	"context"
	"errors"
	"io"
	"strings"
	"testing"
)

type trackedReader struct {
	io.Reader
	closed bool
}

func (r *trackedReader) Close() error { r.closed = true; return nil }

func TestLoaderClosesInjectedReader(t *testing.T) {
	reader := &trackedReader{Reader: strings.NewReader("hello")}
	loader := Loader{Root: "/templates", Open: func(string) (io.ReadCloser, error) { return reader, nil }}
	data, err := loader.Load(context.Background(), "welcome.txt")
	if err != nil || string(data) != "hello" {
		t.Fatalf("data=%q, err=%v", data, err)
	}
	if !reader.closed {
		t.Fatal("loader left reader open")
	}
}

func TestLoaderReturnsOpenError(t *testing.T) {
	want := errors.New("open failed")
	loader := Loader{Open: func(string) (io.ReadCloser, error) { return nil, want }}
	_, err := loader.Load(context.Background(), "missing")
	if !errors.Is(err, want) {
		t.Fatalf("err=%v, want open error", err)
	}
}
