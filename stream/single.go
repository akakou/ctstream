package stream

import (
	"context"
	"sync"
	"time"

	"github.com/akakou/ctstream/client"
	"github.com/akakou/ctstream/core"
)

var DefaultSleep = 1 * time.Second

type CTStream[T core.CtClient] struct {
	Client T
	Sleep  time.Duration
	Ctx    context.Context
	Wg     sync.WaitGroup
}

func NewCTStream[T core.CtClient](
	client T,
	sleep time.Duration,
	Ctx context.Context,
) (*CTStream[T], error) {
	return &CTStream[T]{
		Client: client,
		Sleep:  sleep,
		Ctx:    Ctx,
		Wg:     sync.WaitGroup{},
	}, nil
}

func DefaultCTStream(url string) (*CTStream[*client.CTClient], error) {
	c, err := client.DefaultCTClient(url)
	if err != nil {
		return nil, err
	}

	return NewCTStream[*client.CTClient](
		c,
		DefaultSleep,
		context.Background(),
	)
}

func (stream *CTStream[T]) Init() error {
	return stream.Client.Init()
}

func (stream *CTStream[T]) Next(callback core.Callback) {
	stream.Client.Next(callback)
}

func (stream *CTStream[T]) Run(callback core.Callback) {
	stream.Wg.Add(1)
	defer stream.Wg.Done()

	for {
		select {
		case <-stream.Ctx.Done():
			return
		default:
			stream.Next(callback)
			time.Sleep(stream.Sleep)
		}
	}
}

func (stream *CTStream[T]) Await() {
	stream.Wg.Wait()
}

func (stream *CTStream[T]) Stop() {
	stream.Ctx.Done()
}
