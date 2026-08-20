package domain

import "time"

type Domain struct {
	ID, TenantID, Name, Verification string
	MaxConcurrent                    int
	Paused                           bool
	CreatedAt                        time.Time
}

func New(id, tenant, name string) Domain {
	return Domain{ID: id, TenantID: tenant, Name: name, Verification: "pending", MaxConcurrent: 4, CreatedAt: time.Now().UTC()}
}
func (d Domain) Verified() bool { return d.Verification == "verified" }
func (d *Domain) Verify()       { d.Verification = "verified" }
func (d *Domain) Pause()        { d.Paused = true }
func (d *Domain) Resume()       { d.Paused = false }
