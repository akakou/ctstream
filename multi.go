package ctstream

import (
	"time"
)

type CTsStream struct {
	Streams []*CTStream
	Sleep   time.Duration
}

func NewCTsStream(streams []*CTStream) (*CTsStream, error) {
	return &CTsStream{
		Streams: streams,
	}, nil
}

func DefaultCTsStream(urls []string) (*CTsStream, error) {
	streams := []*CTStream{}

	for _, url := range urls {
		stream, err := DefaultCTStream(url)
		if err != nil {
			return nil, err
		}

		streams = append(streams, stream)
	}

	stream := CTsStream{
		Streams: streams,
	}

	return &stream, nil
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

func (stream *CTsStream) Start(callback Callback, sleep time.Duration) {
	stream.Sleep = sleep

	for _, s := range stream.Streams {
		go s.Run(callback)
	}
}

func (stream *CTsStream) Await() {
	for _, s := range stream.Streams {
		s.Await()
	}
}

func (stream *CTsStream) Run(callback Callback, sleep time.Duration) {
	stream.Start(callback, sleep)
	stream.Await()
}

func (stream *CTsStream) Stop() {
	for _, s := range stream.Streams {
		s.Stop()
	}
}

func (stream *CTsStream) SetTimeout(t time.Duration) {
	for _, s := range stream.Streams {
		s.SetTimeout(t)
	}
}
