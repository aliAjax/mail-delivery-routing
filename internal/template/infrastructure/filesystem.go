package infrastructure

import (
	"context"
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
	if l.Open == nil {
		return os.ReadFile(path)
	}
	reader, err := l.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := reader.Close(); err == nil {
			err = closeErr
		}
	}()
	return io.ReadAll(reader)
}
func (l Loader) Safe(name string) bool { return filepath.Base(name) == name }
