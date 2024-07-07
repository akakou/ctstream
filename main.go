package ctstream

import (
	"time"

	"github.com/google/certificate-transparency-go/client"
	"github.com/google/certificate-transparency-go/jsonclient"
	ctx509 "github.com/google/certificate-transparency-go/x509"
)

type Callback func(*ctx509.Certificate, LogID, *client.LogClient, error)

type CTStream struct {
	streams []*singleStream
	Sleep   time.Duration
	stop    bool
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
		stop:    false,
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
	for _, s := range stream.streams {
		entries, err := s.next()
		first := s.first

		for _, entry := range entries {
			cert, err2 := ExtractCertFromEntry(&entry)
			if err2 != nil {
				panic(err2)
			}

			go callback(cert, first+entry.Index, s.LogClient, err)
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
