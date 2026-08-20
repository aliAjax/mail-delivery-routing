package infrastructure

import (
	"context"
	"net"
)

type Resolver struct{}

func (Resolver) MX(ctx context.Context, domain string) ([]*net.MX, error) {
	return net.DefaultResolver.LookupMX(ctx, domain)
}
func (Resolver) Valid(domain string) bool { _, err := net.LookupHost(domain); return err == nil }
