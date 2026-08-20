package domain

import (
	"errors"
	"testing"
)

func TestHTTPStatusErrorSupportsAsAndTemporary(t *testing.T) {
	var target HTTPStatusError
	err := error(HTTPStatusError{Code: 429})
	if !errors.As(err, &target) || target.Code != 429 || !target.Temporary() {
		t.Fatalf("status=%+v", target)
	}
}
