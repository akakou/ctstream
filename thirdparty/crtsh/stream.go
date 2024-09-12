package crtsh

import (
	"context"

	"github.com/akakou/ctstream/core"
)

func DefaultCTStream(domain string, context context.Context) (*core.CTStream[*CrtshCTClient], error) {
	client, err := NewCTClient(domain)
	if err != nil {
		return nil, err
	}

	stream, err := core.NewCTStream(client, core.DefaultEpochSleep, context)
	if err != nil {
		return nil, err
	}

	return stream, nil
}

func DefaultCTsStream(domains []string, context context.Context) (*core.ConcurrentCTsStream[*core.CTStream[*CrtshCTClient]], error) {
	var streams []*core.CTStream[*CrtshCTClient]

	for _, domain := range domains {
		stream, err := DefaultCTStream(domain, context)
		if err != nil {
			return nil, err
		}

		streams = append(streams, stream)
	}

	res, err := core.NewConcurrentCTsStream(streams, core.DefaultEpochSleep)

	return res, err
}
