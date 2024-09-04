package ctstream

import (
	"time"
)

type CTsStream[T any] struct {
	Streams []*CTStream[T]
	Sleep   time.Duration
}

func NewCTsStream(streams []*CTStream, sleep time.Duration) (*CTsStream, error) {
	return &CTsStream{
		Streams: streams,
		Sleep:   sleep,
	}, nil
}

func DefaultCTsStream[T any](urls []string) (*CTsStream, error) {
	streams := []*CTStream{}

	for _, url := range urls {
		stream, err := DefaultCTStream(url)
		if err != nil {
			return nil, err
		}

		streams = append(streams, stream)
	}

	return NewCTsStream(streams, DefaultSleep)
}

func (stream *CTsStream[T]) Init() error {
	for _, stream := range stream.Streams {
		err := (*stream).Init()
		if err != nil {
			return err
		}
	}

	return nil
}

func (stream *CTsStream[T]) Start(callback Callback[T]) {
	for _, s := range stream.Streams {
		go (*s).Run(callback)
	}
}

func (stream *CTsStream[T]) Await() {
	for _, s := range stream.Streams {
		time.Sleep(stream.Sleep)
		(*s).Await()
	}
}

func (stream *CTsStream) Run(callback Callback[T]) {
	stream.Start(callback)
	stream.Await()
}

func (stream *CTsStream[T]) Stop() {
	for _, s := range stream.Streams {
		(*s).Stop()
	}
}
