package application

import (
	"errors"
	"strings"
)

func ValidAddress(v string) bool {
	return strings.Count(v, "@") == 1 && !strings.ContainsAny(v, "\r\n")
}
func ValidateAttachments(names []string) error {
	for _, n := range names {
		if n == "" || strings.ContainsAny(n, "/\\") || strings.HasPrefix(n, ".") {
			return errors.New("unsafe attachment name")
		}
	}
	return nil
}
func StatusTerminal(s string) bool { return s == "delivered" || s == "failed" || s == "suppressed" }
