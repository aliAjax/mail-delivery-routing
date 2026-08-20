package adapter

import (
	"context"
	"errors"
	"io"
	"strings"
	"testing"
)

type closeReader struct {
	io.Reader
	closed bool
}

func (r *closeReader) Close() error { r.closed = true; return nil }

func TestReadAndCloseClosesReader(t *testing.T) {
	reader := &closeReader{Reader: strings.NewReader("body")}
	got, err := ReadAndClose(func() (io.ReadCloser, error) { return reader, nil })
	if err != nil || got != "body" {
		t.Fatalf("got %q, err=%v", got, err)
	}
	if !reader.closed {
		t.Fatal("reader was not closed")
	}
}

func TestRenderLoadedWrapsLoaderError(t *testing.T) {
	want := errors.New("disk unavailable")
	_, err := RenderLoaded(nil, failingLoader{err: want}, "x")
	if !errors.Is(err, want) {
		t.Fatalf("err=%v, want wrapped loader error", err)
	}
}

type failingLoader struct{ err error }

func (f failingLoader) Load(_ context.Context, _ string) ([]byte, error) {
	return nil, f.err
}
