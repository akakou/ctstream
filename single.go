package ctstream

import (
	"errors"
	"time"
)

type CTStream struct {
	Client *CTClient
	Sleep  time.Duration
	stop   bool
}

func NewCTStream(client *CTClient) (*CTStream, error) {
	return &CTStream{
		Client: client,
	}, nil
}

func DefaultCTStream(url string) (*CTStream, error) {
	client, err := DefaultCTClient(url)
	if err != nil {
		return nil, err
	}

	return &CTStream{
		Client: client,
		stop:   false,
	}, nil
}

func (stream *CTStream) Init() error {
	return stream.Client.Init()
}

func (stream *CTStream) Next(callback Callback) {
	entries, err1 := stream.Client.Next()

	for _, entry := range entries {
		cert, err2 := extractCertFromEntry(&entry)
		err := errors.Join(err1, err2)

		go callback(cert, entry.Index, stream.Client.LogClient, err)
	}
}

func (stream *CTStream) Run(callback Callback, sleep time.Duration) {
	stream.stop = false
	stream.Sleep = sleep

	for {
		stream.Next(callback)

		if stream.stop {
			break
		}

		time.Sleep(stream.Sleep)
	}
}

func (stream *CTStream) Stop() {
	stream.stop = true
}
