package application

import (
	"example.com/maildelivery/internal/queue/domain"
	"testing"
)

func TestClaimDeepClaims(t *testing.T) {
	if got := domain.ClaimState("queued"); got != "running" {
		t.Fatalf("state=%q", got)
	}
	ids := ClaimedIDs([]domain.Job{{ID: "a"}, {ID: "b"}})
	if len(ids) != 2 || ids[0] != "a" || ids[1] != "b" {
		t.Fatalf("ids=%v", ids)
	}
}
