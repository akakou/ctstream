package sslmate

import (
	"context"

	"github.com/akakou/ctstream/core"
)

func AddByDomain(
	domain string,
	ctx context.Context,
	streams *core.ConcurrentCTsStream[*core.CTStream[*SSLMateCTClient]],
) error {
	stream, err := DefaultCTStream(domain, ctx)

	if err != nil {
		return err
	}

	streams.Streams = append(streams.Streams, stream)

	return nil
}

func DelByDomain(
	domain string,
	ctx context.Context,
	streams *core.ConcurrentCTsStream[*core.CTStream[*SSLMateCTClient]],
) error {
	for i, stream := range streams.Streams {
		if stream.Client.Domain == domain {
			streams.Streams = append(streams.Streams[:i], streams.Streams[i+1:]...)
		}
	}

	return nil
}
