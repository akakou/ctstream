package crtsh

import (
	"context"

	"github.com/akakou/ctstream/core"
)

func SelectByDomain(
	domain string,
	clients *core.CTClients[*CrtshCTClient],
) (*CrtshCTClient, int, error) {
	client, i, err := core.SelectByDomain(domain, clients)
	return *client, i, err
}

func AddByDomain(
	domain string,
	ctx context.Context,
	clients *core.CTClients[*CrtshCTClient],
) (*CrtshCTClient, int, error) {
	client, i, err := core.AddByDomain(domain, ctx, nil, clients)
	return *client, i, err
}

func DelByDomain(
	domain string,
	ctx context.Context,
	clients *core.CTClients[*CrtshCTClient],
) (*CrtshCTClient, int, error) {
	client, i, err := core.DelByDomain(domain, ctx, clients)
	return *client, i, err
}
