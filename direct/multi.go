package direct

import (
	"github.com/akakou/ctstream/client"
	"github.com/akakou/ctstream/core"
)

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
