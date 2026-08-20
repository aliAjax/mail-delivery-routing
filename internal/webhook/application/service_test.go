package application

import (
	"context"
	"errors"
	"example.com/maildelivery/internal/webhook/domain"
	"net/http"
	"net/http/httptest"
	"runtime/debug"
	"testing"
)

func TestDeliverCheckedPreservesHTTPStatusError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusServiceUnavailable) }))
	defer server.Close()
	err := New().DeliverChecked(context.Background(), domain.Delivery{URL: server.URL}, map[string]string{"event": "bounce"})
	var statusErr domain.HTTPStatusError
	if !errors.As(err, &statusErr) {
		debug.PrintStack()
		t.Fatalf("err=%v is not HTTPStatusError", err)
	}
	if statusErr.Code != http.StatusServiceUnavailable || !statusErr.Temporary() {
		t.Fatalf("status error=%+v", statusErr)
	}
}
