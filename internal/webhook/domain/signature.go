package domain

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func Sign(secret, payload string) string {
	h := hmac.New(sha256.New, []byte(secret))
	_, _ = h.Write([]byte(payload))
	return hex.EncodeToString(h.Sum(nil))
}
func Verify(secret, payload, sig string) bool {
	return hmac.Equal([]byte(Sign(secret, payload)), []byte(sig))
}
