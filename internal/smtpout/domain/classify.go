package domain

import "strings"

func NormalizeAddress(v string) string {
	return strings.Trim(strings.ToLower(strings.TrimSpace(v)), "<>")
}
func SameDomain(a, b string) bool {
	pa := strings.Split(NormalizeAddress(a), "@")
	pb := strings.Split(NormalizeAddress(b), "@")
	return len(pa) == 2 && len(pb) == 2 && pa[1] == pb[1]
}
