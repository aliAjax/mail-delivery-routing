package application

import (
	"context"
	"example.com/maildelivery/internal/domain/application"
	domaininfra "example.com/maildelivery/internal/domain/infrastructure"
	"example.com/maildelivery/internal/tenant/domain"
	"testing"
)

func TestQuotaDeepClaims(t *testing.T) {
	if got := domain.QuotaBudget(3, 5); got != 2 {
		t.Fatalf("budget=%d", got)
	}
	if got := application.SafeUsage(-2); got != 0 {
		t.Fatalf("usage=%d", got)
	}
	values := []int{1, 2}
	clone := domaininfra.CloneUsage(values)
	clone[0] = 9
	if values[0] != 1 {
		t.Fatal("usage snapshot aliased")
	}
	s := New()
	s.Put(context.Background(), domain.Tenant{ID: "deep", DailyQuota: 2})
	got, err := ReservationSnapshot(s, context.Background(), "deep")
	if err != nil || got.Used != 1 {
		t.Fatalf("tenant=%+v err=%v", got, err)
	}
}
