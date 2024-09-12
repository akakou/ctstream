package core

import (
	"context"
	"errors"
)

const ERROR_NOT_FOUND = "not found"

type DefaultCTStream[C CtClient] func(string, context.Context) (C, error)

func SelectByDomain[C CtClient](
	domain string,
	clients *CTClients[C],
) (*C, int, error) {
	for i, c := range clients.Clients {
		if c.GetDomain() == domain {
			return &c, i, nil
		}
	}

	return nil, -1, errors.New(ERROR_NOT_FOUND)
}

func AddByDomain[C CtClient](
	domain string,
	ctx context.Context,
	def DefaultCTStream[C],
	clients *CTClients[C],
) (*C, int, error) {
	client, err := def(domain, ctx)

	if err != nil {
		return nil, 0, err
	}

	clients.Clients = append(clients.Clients, client)

	return &client, len(clients.Clients) - 1, nil
}

func DelByDomain[C CtClient](
	domain string,
	ctx context.Context,
	clients *CTClients[C],
) (*C, int, error) {
	client, i, err := SelectByDomain(domain, clients)
	if err != nil {
		return nil, 0, err
	}

	clients.Clients = append(clients.Clients[:i], clients.Clients[i+1:]...)

	return client, i, nil
}
