package application

import (
	"example.com/maildelivery/internal/template/domain"
	"testing"
)

func TestFallbackDeepClaims(t *testing.T) {
	if got := fallbackValue(map[string]string{"name": "A"}, map[string]string{"name": "B"}, "name"); got != "A" {
		t.Fatalf("value=%q", got)
	}
	copy := domain.CopyVariables(map[string]string{"name": "A"})
	copy["name"] = "B"
	if copy["name"] != "B" {
		t.Fatal("copy failed")
	}
}
