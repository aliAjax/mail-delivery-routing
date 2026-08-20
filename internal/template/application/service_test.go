package application

import (
	"context"
	"example.com/maildelivery/internal/template/domain"
	"testing"
)

func TestRenderWithFallbackAcceptsNilVariables(t *testing.T) {
	s := New()
	s.Put(context.Background(), domain.Template{ID: "welcome", Body: "Hi {{name}} from {{region}}"})
	got, err := s.RenderWithFallback(context.Background(), "welcome", nil, map[string]string{"name": "Mina", "region": "Tokyo"})
	if err != nil {
		t.Fatal(err)
	}
	if got != "Hi Mina from Tokyo" {
		t.Fatalf("got %q", got)
	}
}
