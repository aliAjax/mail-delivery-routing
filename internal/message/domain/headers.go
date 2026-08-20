package domain

import (
	"errors"
	"strings"
)

func ValidateHeaders(headers map[string]string) error {
	for k, v := range headers {
		if strings.ContainsAny(k+v, "\r\n") {
			return errors.New("header contains newline")
		}
		if k == "Bcc" || strings.EqualFold(k, "X-Internal-Token") {
			return errors.New("restricted header")
		}
	}
	return nil
}
func CanonicalHeaders(headers map[string]string) map[string]string {
	out := map[string]string{}
	for k, v := range headers {
		out[strings.ToLower(strings.TrimSpace(k))] = strings.TrimSpace(v)
	}
	return out
}
