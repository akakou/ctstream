package ctstream

import (
	"time"

	"github.com/google/certificate-transparency-go/jsonclient"
)

type CTStream struct {
	streams []*singleStream
	Sleep   time.Duration
}

func New(urls []string, sleep time.Duration) (*CTStream, error) {
	streams := []*singleStream{}

	for _, url := range urls {
		stream, err := newSingleStream(url, jsonclient.Options{})
		if err != nil {
			return nil, err
		}

		streams = append(streams, stream)
	}

	stream := CTStream{
		streams: streams,
		Sleep:   sleep,
	}

	return &stream, nil
}

func (stream *CTStream) Init() error {
	for _, stream := range stream.streams {
		err := stream.init()
		if err != nil {
			return err
		}
	}

	return nil
}

func (stream *CTStream) Next(callback Callback) {
	for _, m := range stream.streams {
		entries, err := m.next()
		first := m.first

		for _, entry := range entries {
			go callback(entry.Precert, first+entry.Index, err)
		}
	}

	time.Sleep(stream.Sleep)
}

func (stream *CTStream) Run(callback Callback) {
	for {
		stream.Next(callback)
	}
}
