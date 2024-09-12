package crtsh

import (
	"context"

	"github.com/akakou/ctstream/core"
)

func SelectByDomain(
	domain string,
	ctx context.Context,
	streams *core.ConcurrentCTsStream[*core.CTStream[*CrtshCTClient]],
) (*core.CTStream[*CrtshCTClient], int, error) {
	return core.SelectByDomain(domain, streams)
}

func AddByDomain(
	domain string,
	ctx context.Context,
	streams *core.ConcurrentCTsStream[*core.CTStream[*CrtshCTClient]],
) error {
	return core.AddByDomain(domain, ctx, DefaultCTStream, streams)
}

func DelByDomain(
	domain string,
	ctx context.Context,
	streams *core.ConcurrentCTsStream[*core.CTStream[*CrtshCTClient]],
) error {
	return core.DelByDomain(domain, ctx, streams)
}
