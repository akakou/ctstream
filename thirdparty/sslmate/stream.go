package sslmate

import (
	"context"
	"time"

	"github.com/akakou/ctstream/core"
)

var DefaultEpochSleep = time.Minute * 20
var DefaultPullingSleep = time.Second * 10

func NewCTStream(client *SSLMateCTClient, sleep time.Duration, context context.Context) (*core.CTStream[*SSLMateCTClient], error) {
	return core.NewCTStream(client, sleep, context)
}

func DefaultCTStream(domain string, context context.Context) (*core.CTStream[*SSLMateCTClient], error) {
	client, err := DefaultCTClient(domain)
	if err != nil {
		return nil, err
	}

	stream, err := core.NewCTStream(client, DefaultEpochSleep, context)
	if err != nil {
		return nil, err
	}

	stream.Sleep = DefaultEpochSleep

	return stream, nil
}

func NewCTsStream(streams []*core.CTStream[*SSLMateCTClient], sleep time.Duration) (*core.SyncCTsStream[*core.CTStream[*SSLMateCTClient]], error) {
	return core.NewSyncCTsStream(streams, sleep)
}

func DefaultCTsStream(domains []string, context context.Context) (*core.SyncCTsStream[*core.CTStream[*SSLMateCTClient]], error) {
	var streams []*core.CTStream[*SSLMateCTClient]

	for _, domain := range domains {
		stream, err := DefaultCTStream(domain, context)
		if err != nil {
			return nil, err
		}

		streams = append(streams, stream)
	}

	res, err := core.NewSyncCTsStream(streams, DefaultEpochSleep)

	return res, err
}
