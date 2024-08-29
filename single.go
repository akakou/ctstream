package ctstream

import (
	"context"
	"sync"
	"time"
)

var DefaultSleep = 500 * time.Millisecond

type CTStream struct {
	Client *CTClient
	Sleep  time.Duration
	Ctx    context.Context
	Wg     sync.WaitGroup
	Stop   context.CancelFunc
}

func NewCTStream(client *CTClient, sleep time.Duration, Ctx context.Context) (*CTStream, error) {
	ctx, cancel := context.WithCancel(Ctx)
	return &CTStream{
		Client: client,
		Sleep:  sleep,
		Ctx:    ctx,
		Stop:   cancel,
		Wg:     sync.WaitGroup{},
	}, nil
}

func DefaultCTStream(url string, ctx context.Context) (*CTStream, error) {
	client, err := DefaultCTClient(url)
	if err != nil {
		return nil, err
	}

	return NewCTStream(
		client,
		DefaultSleep,
		ctx,
	)
}

func (stream *CTStream) Init() error {
	return stream.Client.Init()
}

func (stream *CTStream) Next(callback Callback) {
	stream.Client.Next(callback)
}

func (stream *CTStream) start(callback Callback) {
	for {
		select {
		case <-stream.Ctx.Done():
			stream.Wg.Done()
			return
		default:
			stream.Next(callback)
			time.Sleep(stream.Sleep)
		}
	}
}

func (stream *CTStream) Start(callback Callback) {
	stream.Wg.Add(1)
	go stream.start(callback)
}

func (stream *CTStream) Run(callback Callback) {
	stream.Start(callback)
	stream.Await()
}

func (stream *CTStream) Await() {
	stream.Wg.Wait()
}
