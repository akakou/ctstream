package ctstream

import (
	"time"
)

type CTsStream struct {
	streams []*CTStream
	Sleep   time.Duration
	stop    bool
}

func NewCTsStream(streams []*CTStream) (*CTsStream, error) {
	return &CTsStream{
		streams: streams,
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
		streams: streams,
	}

	return &stream, nil
}

func (stream *CTsStream) Init() error {
	for _, stream := range stream.streams {
		err := stream.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

func (stream *CTsStream) Start(callback Callback, sleep time.Duration) {
	stream.stop = false
	stream.Sleep = sleep

	for _, s := range stream.streams {
		go s.Run(callback, sleep)
	}
}

func (stream *CTsStream) Await() {
	for !stream.stop {
		time.Sleep(stream.Sleep)
	}
}

func (stream *CTsStream) Run(callback Callback, sleep time.Duration) {
	stream.Start(callback, sleep)
	stream.Await()
}

func (stream *CTsStream) Stop() {
	stream.stop = true

	for _, s := range stream.streams {
		s.Stop()
	}
}
