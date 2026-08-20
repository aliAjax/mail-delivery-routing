package domain

import (
	"net/mail"
	"strings"
)

type Policy struct {
	TenantID        string
	AllowedDomains  []string
	MaxMessageBytes int64
	RequireVerified bool
}

func (p Policy) Allows(address string) bool {
	a, err := mail.ParseAddress(address)
	if err != nil {
		return false
	}
	if len(p.AllowedDomains) == 0 {
		return true
	}
	for _, d := range p.AllowedDomains {
		if strings.EqualFold(strings.Split(a.Address, "@")[1], d) {
			return true
		}
	}
	return false
}
