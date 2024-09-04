package stream

import (
	"time"

	"github.com/akakou/ctstream/client"
	"github.com/akakou/ctstream/core"
)

type CTsStream[T core.CtStream] struct {
	Streams []T
	Sleep   time.Duration
}

func NewCTsStream[T core.CtStream](streams []T, sleep time.Duration) (*CTsStream[T], error) {
	return &CTsStream[T]{
		Streams: streams,
		Sleep:   sleep,
	}, nil
}

func DefaultCTsStream(urls []string) (*CTsStream[*CTStream[*client.CTClient]], error) {
	streams := []*CTStream[*client.CTClient]{}

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
	for _, s := range stream.Streams {
		err := s.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

func (stream *CTsStream[T]) Start(callback core.Callback) {
	for _, s := range stream.Streams {
		go s.Run(callback)
	}
}

func (stream CTsStream[T]) Await() {
	for _, s := range stream.Streams {
		time.Sleep(stream.Sleep)
		s.Await()
	}
}

func (stream *CTsStream[T]) Run(callback core.Callback) {
	stream.Start(callback)
	stream.Await()
}

func (stream *CTsStream[T]) Stop() {
	for _, s := range stream.Streams {
		s.Stop()
	}
}
