package adapter

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"strings"
)

func EscapeHTML(s string) string { return template.HTMLEscapeString(s) }
func NormalizeLocale(v string) string {
	return strings.ToLower(strings.ReplaceAll(strings.TrimSpace(v), "_", "-"))
}

type Loader interface {
	Load(context.Context, string) ([]byte, error)
}

func RenderLoaded(ctx context.Context, loader Loader, name string) (string, error) {
	data, err := loader.Load(ctx, name)
	if err != nil {
		return "", fmt.Errorf("load template: %w", err)
	}
	return EscapeHTML(string(data)), nil
}

func ReadAndClose(read func() (io.ReadCloser, error)) (string, error) {
	closer, err := read()
	if err != nil {
		return "", err
	}
	defer closer.Close()
	data, err := io.ReadAll(closer)
	if err != nil {
		return "", fmt.Errorf("read template: %w", err)
	}
	return string(data), nil
}
