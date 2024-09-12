package crtsh

import (
	"context"

	"github.com/akakou/ctstream/core"
	"github.com/akakou/ctstream/thirdparty"
)

func AddByDomain(
	domain string,
	ctx context.Context,
	streams *core.ConcurrentCTsStream[*core.CTStream[*CrtshCTClient]],
) error {
	return thirdparty.AddByDomain(domain, ctx, DefaultCTStream, streams)
}

func DelByDomain(
	domain string,
	ctx context.Context,
	streams *core.ConcurrentCTsStream[*core.CTStream[*CrtshCTClient]],
) error {
	return thirdparty.DelByDomain(domain, ctx, streams)
}
