package adapter

import "example.com/maildelivery/internal/template/domain"

type Request struct {
	ID, TenantID, Locale, Subject, Body string
	Version                             int
}

func ToDomain(r Request) domain.Template {
	return domain.Template{ID: r.ID, TenantID: r.TenantID, Locale: r.Locale, Subject: r.Subject, Body: r.Body, Version: r.Version}
}
