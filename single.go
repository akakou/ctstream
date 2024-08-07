package ctstream

import (
	"context"
	"errors"
	"sync"
	"time"
)

var DefaultSleep = 1 * time.Second

type CTStream struct {
	Client *CTClient
	Sleep  time.Duration
	Ctx    context.Context
	Wg     sync.WaitGroup
}

func NewCTStream(client *CTClient, sleep time.Duration, Ctx context.Context) (*CTStream, error) {
	return &CTStream{
		Client: client,
		Sleep:  sleep,
		Ctx:    Ctx,
		Wg:     sync.WaitGroup{},
	}, nil
}

func DefaultCTStream(url string) (*CTStream, error) {
	client, err := DefaultCTClient(url)
	if err != nil {
		return nil, err
	}

	return &CTStream{
		Client: client,
		Sleep:  DefaultSleep,
		Ctx:    context.Background(),
	}, nil
}

func (stream *CTStream) Init() error {
	return stream.Client.Init()
}

func (stream *CTStream) Next(callback Callback) {
	entries, err1 := stream.Client.Next()

	for _, entry := range entries {
		cert, err2 := extractCertFromEntry(&entry)
		err := errors.Join(err1, err2)

		callback(cert, entry.Index, stream.Client.LogClient, err)
	}
}

func (stream *CTStream) Run(callback Callback) {
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

func (stream *CTStream) Await() {
	stream.Wg.Wait()
}

func (stream *CTStream) Stop() {
	stream.Ctx.Done()
}
