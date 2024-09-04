package core

import (
	"context"
	"sync"
	"time"
)

var DefaultSleep = 1 * time.Second

type CTStream[T CtClient] struct {
	Client T
	Sleep  time.Duration
	Ctx    context.Context
	Wg     sync.WaitGroup
}

func NewCTStream[T CtClient](
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

func (stream *CTStream[T]) Init() error {
	return stream.Client.Init()
}

func (stream *CTStream[T]) Next(callback Callback) {
	stream.Client.Next(callback)
}

func (stream *CTStream[T]) Run(callback Callback) {
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
