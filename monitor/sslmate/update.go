package sslmate

import (
	"context"

	"github.com/akakou/ctstream/core"
	"github.com/akakou/ctstream/monitor"
)

func AddByDomain(
	domain string,
	ctx context.Context,
	streams *core.ConcurrentCTsStream[*core.CTStream[*SSLMateCTClient]],
) error {
	return monitor.AddByDomain(domain, ctx, DefaultCTStream, streams)
}

func DelByDomain(
	domain string,
	ctx context.Context,
	streams *core.ConcurrentCTsStream[*core.CTStream[*SSLMateCTClient]],
) error {
	return monitor.DelByDomain(domain, ctx, streams)
}
