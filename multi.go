package ctstream

import (
	"context"
	"time"
)

type CTsStream struct {
	Streams []*CTStream
	Sleep   time.Duration
}

func NewCTsStream(streams []*CTStream, sleep time.Duration) (*CTsStream, error) {
	return &CTsStream{
		Streams: streams,
		Sleep:   sleep,
	}, nil
}

func SemiDefaultCTsStream(streams []*CTStream, sleep time.Duration) (*CTsStream, error) {
	return &CTsStream{
		Streams: streams,
		Sleep:   sleep,
	}, nil
}

func DefaultCTsStream(urls []string, ctx context.Context) (*CTsStream, error) {
	streams := []*CTStream{}

	for _, url := range urls {
		stream, err := DefaultCTStream(url, ctx)
		if err != nil {
			return nil, err
		}

		streams = append(streams, stream)
	}

	return NewCTsStream(streams, DefaultSleep)
}

func (stream *CTsStream) Init() error {
	for _, stream := range stream.Streams {
		err := stream.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

func (stream *CTsStream) Start(callback Callback) {
	for _, s := range stream.Streams {
		s.Start(callback)
	}
}

func (stream *CTsStream) Await() {
	for _, s := range stream.Streams {
		s.Await()
	}
}

func (stream *CTsStream) Run(callback Callback) {
	stream.Start(callback)
	stream.Await()
}

func (stream *CTsStream) Stop() {
	for _, s := range stream.Streams {
		s.Stop()
	}
}
