package domain

import "strings"

func ReasonCode(text string) string {
	for _, code := range []string{"5.1.1", "5.2.2", "4.2.0", "4.4.1"} {
		if strings.Contains(text, code) {
			return code
		}
	}
	return "unknown"
}
func IsComplaint(text string) bool {
	return strings.Contains(strings.ToLower(text), "complaint") || strings.Contains(strings.ToLower(text), "feedback-type")
}
