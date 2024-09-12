package sslmate

import (
	"context"

	"github.com/akakou/ctstream/core"
)

func SelectByDomain(
	domain string,
	streams *core.ConcurrentCTsStream[*core.CTStream[*SSLMateCTClient]],
) (*core.CTStream[*SSLMateCTClient], int, error) {
	return core.SelectByDomain(domain, streams)
}

func AddByDomain(
	domain string,
	ctx context.Context,
	streams *core.ConcurrentCTsStream[*core.CTStream[*SSLMateCTClient]],
) error {
	return core.AddByDomain(domain, ctx, DefaultCTStream, streams)
}

func DelByDomain(
	domain string,
	ctx context.Context,
	streams *core.ConcurrentCTsStream[*core.CTStream[*SSLMateCTClient]],
) error {
	return core.DelByDomain(domain, ctx, streams)
}
