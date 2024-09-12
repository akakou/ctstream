package monitor

import (
	"context"

	"github.com/akakou/ctstream/core"
)

type CTMonitorClient interface {
	GetDomain() string
	Init() error
	Next(core.Callback)
}

type DefaultCTStream[T CTMonitorClient] func(string, context.Context) (*core.CTStream[T], error)

func AddByDomain[C CTMonitorClient](
	domain string,
	ctx context.Context,
	def DefaultCTStream[C],
	streams *core.ConcurrentCTsStream[*core.CTStream[C]],
) error {
	stream, err := def(domain, ctx)

	if err != nil {
		return err
	}

	streams.Streams = append(streams.Streams, stream)

	return nil
}

func DelByDomain[C CTMonitorClient](
	domain string,
	ctx context.Context,
	streams *core.ConcurrentCTsStream[*core.CTStream[C]],
) error {
	for i, stream := range streams.Streams {
		if stream.Client.GetDomain() == domain {
			streams.Streams = append(streams.Streams[:i], streams.Streams[i+1:]...)
		}
	}

	return nil
}
