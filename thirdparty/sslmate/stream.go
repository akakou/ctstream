package sslmate

import (
	"context"
	"time"

	"github.com/akakou/ctstream/core"
)

var SSLMateDefaultSleep = time.Minute * 20

func NewSSLMateCTStream(client *SSLMateCTClient, sleep time.Duration, ctx context.Context) (*core.CTStream[*SSLMateCTClient], error) {
	return core.NewCTStream(client, sleep, ctx)
}

func DefaultSSLMateCTStream(domain string, sleep time.Duration, ctx context.Context) (*core.CTStream[*SSLMateCTClient], error) {
	client, err := DefaultSSLMateCTClient(domain)
	if err != nil {
		return nil, err
	}

	stream, err := core.NewCTStream(client, sleep, ctx)
	if err != nil {
		return nil, err
	}

	stream.Sleep = SSLMateDefaultSleep

	return stream, nil
}

func NewSSLMateCTsStream(streams []*core.CTStream[*SSLMateCTClient], sleep time.Duration) (*core.CTsStream[*core.CTStream[*SSLMateCTClient]], error) {
	return core.NewCTsStream(streams, sleep)
}

func DefaultSSLMateCTsStream(domains []string) (*core.CTsStream[*core.CTStream[*SSLMateCTClient]], error) {
	streams := []*core.CTStream[*SSLMateCTClient]{}
	for _, domain := range domains {
		stream, err := DefaultSSLMateCTStream(domain, SSLMateDefaultSleep, context.Background())

		if err != nil {
			return nil, err
		}

		streams = append(streams, stream)
	}

	return core.NewCTsStream(streams, SSLMateDefaultSleep)
}
