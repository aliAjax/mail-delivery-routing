package adapter

import "example.com/maildelivery/internal/domain/domain"

type DTO struct {
	ID, TenantID, Name, Verification string
	Paused                           bool
}

func FromDomain(d domain.Domain) DTO {
	return DTO{ID: d.ID, TenantID: d.TenantID, Name: d.Name, Verification: d.Verification, Paused: d.Paused}
}
func ToDomain(v DTO) domain.Domain {
	return domain.Domain{ID: v.ID, TenantID: v.TenantID, Name: v.Name, Verification: v.Verification, Paused: v.Paused}
}
