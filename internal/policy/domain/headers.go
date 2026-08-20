package domain

import "strings"

func SafeHeader(name, value string) bool {
	if strings.ContainsAny(name+value, "\r\n") {
		return false
	}
	switch strings.ToLower(name) {
	case "received", "return-path":
		return false
	}
	return true
}
