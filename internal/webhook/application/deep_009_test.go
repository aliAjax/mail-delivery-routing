package application

import (
	"errors"
	"example.com/maildelivery/internal/webhook/domain"
	"testing"
)

func TestStatusDeepClaims(t *testing.T) {
	if !domain.RetryableStatus(503) || domain.RetryableStatus(400) {
		t.Fatal("status policy wrong")
	}
	var target domain.HTTPStatusError
	if !errors.As(StatusFailure(503), &target) || !target.Temporary() {
		t.Fatalf("status=%+v", target)
	}
}

func TestStatusPolicyClaims(t *testing.T) {
	if !domain.RetryableStatus(503) || domain.RetryableStatus(400) {
		t.Fatal("status policy wrong")
	}
}

func TestStatusBridgeClaims(t *testing.T) {
	var target domain.HTTPStatusError
	if !errors.As(StatusFailure(503), &target) || !target.Temporary() {
		t.Fatalf("status=%+v", target)
	}
}

func TestStatusClassClaims(t *testing.T) {
	if got := domain.StatusClassForDelivery(503); got != "server" {
		t.Fatalf("class=%q", got)
	}
}
