package direct

import (
	"context"
	"time"

	"github.com/akakou/ctstream/client"
	"github.com/akakou/ctstream/core"
	"github.com/google/certificate-transparency-go/jsonclient"
)

var DefaultSleep = 1 * time.Second

func DefaultCTClient(url string) (*client.CTClient, error) {
	return client.NewCTClient(url, core.DefaultMaxEntries, jsonclient.Options{})
}

func DefaultCTStream(url string) (*core.CTStream[*client.CTClient], error) {
	c, err := DefaultCTClient(url)
	if err != nil {
		return nil, err
	}

	return core.NewCTStream(
		c,
		DefaultSleep,
		context.Background(),
	)
}

func DefaultCTsStream(urls []string) (*core.CTsStream[*core.CTStream[*client.CTClient]], error) {
	streams := []*core.CTStream[*client.CTClient]{}

	for _, url := range urls {
		stream, err := DefaultCTStream(url)
		if err != nil {
			return nil, err
		}

		streams = append(streams, stream)
	}

	return core.NewCTsStream(streams, DefaultSleep)
}
