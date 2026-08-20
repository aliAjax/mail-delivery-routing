package application

import (
	"example.com/maildelivery/internal/template/domain"
	"testing"
)

func TestFallbackDeepClaims(t *testing.T) {
	if got := fallbackValue(map[string]string{"name": "A"}, map[string]string{"name": "B"}, "name"); got != "A" {
		t.Fatalf("value=%q", got)
	}
	values := map[string]string{"name": "A"}
	copy := domain.CopyVariables(values)
	copy["name"] = "B"
	if values["name"] != "A" {
		t.Fatal("copy failed")
	}
}
