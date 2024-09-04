package ctstream

import (
	"context"
	"errors"
	"sync"
	"time"
)

var DefaultSleep = 1 * time.Second

// type CtStream[T any] interface {
// 	Init() error
// 	Await()
// 	Run(callback Callback[T])
// 	Stop()
// }

type CTStream[T any] struct {
	Client *CtClient
	Sleep  time.Duration
	Ctx    context.Context
	Wg     sync.WaitGroup
}

func NewCTStream[T any](client *CTClient, sleep time.Duration, Ctx context.Context) (*CTStream[T], error) {
	return &CTStream[T]{
		Client: client,
		Sleep:  sleep,
		Ctx:    Ctx,
		Wg:     sync.WaitGroup{},
	}, nil
}

func DefaultCTStream[T any](url string) (*CTStream[T], error) {
	client, err := DefaultCTClient(url)
	if err != nil {
		return nil, err
	}

	return NewCTStream[T](
		client,
		DefaultSleep,
		context.Background(),
	)
}

func (stream *CTStream[T]) Init() error {
	return stream.Client.Init()
}

func (stream *CTStream[T]) Next(callback Callback[CTClientParams]) {
	entries, err1 := stream.Client.Next()

	for _, entry := range entries {
		cert, err2 := extractCertFromEntry(&entry)
		err := errors.Join(err1, err2)

		callback(cert, CTClientParams{
			entry.Index,
			stream.Client.LogClient,
		}, err)
	}
}

func (stream *CTStream[T]) Run(callback Callback[CTClientParams]) {
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
