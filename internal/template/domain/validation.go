package domain

import (
	"errors"
	"strings"
)

func ValidateVars(vars map[string]string) error {
	for k := range vars {
		if strings.TrimSpace(k) == "" || strings.ContainsAny(k, "{} ") {
			return errors.New("invalid variable name")
		}
	}
	return nil
}
func Locales(values map[string]string) []string {
	out := make([]string, 0, len(values))
	for k := range values {
		out = append(out, k)
	}
	return out
}
