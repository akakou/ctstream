package ctstream

import (
	"time"

	"github.com/google/certificate-transparency-go/client"
	"github.com/google/certificate-transparency-go/jsonclient"
	ctx509 "github.com/google/certificate-transparency-go/x509"
)

const DEFAULT_MAX_ENTRIES = 256

type Callback func(*ctx509.Certificate, LogID, *client.LogClient, error)

type CTStream struct {
	streams      []*singleStream
	Sleep        time.Duration
	stop         bool
	MaxEntrySize int
}

func New(urls []string, maxEntrySize int64, sleep time.Duration) (*CTStream, error) {
	streams := []*singleStream{}

	for _, url := range urls {
		stream, err := newSingleStream(url, maxEntrySize, jsonclient.Options{})
		if err != nil {
			return nil, err
		}

		streams = append(streams, stream)
	}

	stream := CTStream{
		streams: streams,
		Sleep:   sleep,
		stop:    false,
	}

	return &stream, nil
}

func Default(urls []string, sleep time.Duration) (*CTStream, error) {
	return New(urls, DEFAULT_MAX_ENTRIES, sleep)
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
	for _, s := range stream.streams {
		entries, err1 := s.next()

		for _, entry := range entries {
			cert, err2 := ExtractCertFromEntry(&entry)
			if err2 != nil {
				panic(err2)
			}

			go callback(cert, entry.Index, s.LogClient, err1)
		}
	}

	time.Sleep(stream.Sleep)
}

func (stream *CTStream) Run(callback Callback) {
	for {
		stream.Next(callback)

		if stream.stop {
			break
		}
	}
}

func (stream *CTStream) Stop() {
	stream.stop = true
}
