package application

import (
	"example.com/maildelivery/internal/message/domain"
	"testing"
)

func TestStateDeepClaims(t *testing.T) {
	if !domain.TerminalState("delivered") {
		t.Fatal("delivered not terminal")
	}
	if got := StatusVisible(""); got != "queued" {
		t.Fatalf("status=%q", got)
	}
}
