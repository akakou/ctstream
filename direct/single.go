package direct

import (
	"context"
	"time"

	"github.com/akakou/ctstream/client"
	"github.com/akakou/ctstream/core"
)

var DefaultSleep = 1 * time.Second

func DefaultCTStream(url string) (*core.CTStream[*client.CTClient], error) {
	c, err := client.DefaultCTClient(url)
	if err != nil {
		return nil, err
	}

	return core.NewCTStream[*client.CTClient](
		c,
		DefaultSleep,
		context.Background(),
	)
}
