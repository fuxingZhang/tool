package tool

import (
	"context"
	"net"
)

// LookupHost looks up the given host using the local resolver.
// It returns a slice of that host's addresses.
func LookupHost(dns, host string) (addrs []string, err error) {
	var dialer net.Dialer
	var r = &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			return dialer.DialContext(ctx, network, dns)
		},
	}
	return r.LookupHost(context.Background(), host)
}
