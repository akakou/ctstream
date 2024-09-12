package core

import (
	"context"
	"errors"
)

const ERROR_NOT_FOUND = "not found"

type DefaultCTStream[T CtClient] func(string, context.Context) (*CTStream[T], error)

func SelectByDomain[C CtClient](
	domain string,
	streams *ConcurrentCTsStream[*CTStream[C]],
) (*CTStream[C], int, error) {
	for i, stream := range streams.Streams {
		if stream.Client.GetDomain() == domain {
			return stream, i, nil
		}
	}

	return nil, -1, errors.New(ERROR_NOT_FOUND)
}

func AddByDomain[C CtClient](
	domain string,
	ctx context.Context,
	def DefaultCTStream[C],
	streams *ConcurrentCTsStream[*CTStream[C]],
) (*CTStream[C], int, error) {
	stream, err := def(domain, ctx)

	if err != nil {
		return nil, 0, err
	}

	streams.Streams = append(streams.Streams, stream)

	return stream, len(streams.Streams) - 1, nil
}

func DelByDomain[C CtClient](
	domain string,
	ctx context.Context,
	streams *ConcurrentCTsStream[*CTStream[C]],
) (*CTStream[C], int, error) {
	stream, i, err := SelectByDomain(domain, streams)
	if err != nil {
		return nil, 0, err
	}

	streams.Streams = append(streams.Streams[:i], streams.Streams[i+1:]...)

	return stream, i, nil
}
