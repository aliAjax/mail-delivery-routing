package application

import (
	"context"
	"errors"
	"example.com/maildelivery/internal/template/domain"
	"testing"
)

func TestRenderWithFallbackHonorsCanceledContext(t *testing.T) {
	s := New()
	s.Put(context.Background(), domain.Template{ID: "welcome", Body: "Hi {{name}}"})
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := s.RenderWithFallback(ctx, "welcome", nil, map[string]string{"name": "Mina"})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("err=%v, want context.Canceled", err)
	}
}
