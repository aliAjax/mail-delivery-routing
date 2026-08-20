package adapter

type DTO struct {
	TenantID        string
	AllowedDomains  []string
	MaxMessageBytes int64
	RequireVerified bool
}
