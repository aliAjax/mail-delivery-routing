package infrastructure

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Loader struct {
	Root string
	Open func(string) (io.ReadCloser, error)
}

func (l Loader) Load(_ context.Context, name string) (data []byte, err error) {
	path := filepath.Join(l.Root, filepath.Clean(name))
	if !l.Safe(name) {
		return nil, fmt.Errorf("unsafe template path: %s", name)
	}
	if l.Open == nil {
		return os.ReadFile(path)
	}
	reader, err := l.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open template: %w", err)
	}
	if reader == nil {
		return nil, fmt.Errorf("open template: nil reader")
	}
	defer func() { _ = reader.Close() }()
	data, err = io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("read template: %w", err)
	}
	return data, nil
}
func (l Loader) Safe(name string) bool { return filepath.Base(name) == name }
