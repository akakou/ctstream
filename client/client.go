package client

import (
	"context"
	"net/http"

	"errors"

	"github.com/akakou/ctstream/core"

	ct "github.com/google/certificate-transparency-go"
	"github.com/google/certificate-transparency-go/client"
	"github.com/google/certificate-transparency-go/jsonclient"
)

type CTClientParams struct {
	Index     core.LogID
	LogClient *client.LogClient
}

type CTClient struct {
	Url string
	*client.LogClient
	context.Context
	first        core.LogID
	opts         jsonclient.Options
	maxEntrySize int64
}

func NewCTClient(url string, maxEntrySize int64, ops jsonclient.Options) (*CTClient, error) {
	hc := http.Client{}
	ctx := context.Background()

	c, err := client.New(url, &hc, ops)
	if err != nil {
		return nil, errors.New(ERROR_FAILED_TO_NEW)
	}

	return &CTClient{
		Url:          url,
		LogClient:    c,
		Context:      ctx,
		maxEntrySize: maxEntrySize,
		opts:         ops,
	}, nil
}

func DefaultCTClient(url string) (*CTClient, error) {
	return NewCTClient(url, core.DefaultMaxEntries, jsonclient.Options{})
}

func (stream *CTClient) Init() error {
	sct, err := stream.GetSTH(stream.Context)
	if err != nil {
		return errors.New(ERROR_FAILED_TO_FETCH_STH)
	}

	stream.first = int64(sct.TreeSize)
	return nil
}

func (stream *CTClient) next() ([]ct.LogEntry, error) {
	sct, err := stream.LogClient.GetSTH(stream.Context)
	if err != nil {
		return []ct.LogEntry{}, errors.New(ERROR_FAILED_TO_FETCH_STH)
	}

	if sct.TreeSize == uint64(stream.first) {
		return []ct.LogEntry{}, errors.New(ERROR_NEW_LOGS_NOT_FOUND)
	}

	last := sct.TreeSize
	if sct.TreeSize > uint64(stream.first+stream.maxEntrySize) {
		last = uint64(stream.first + stream.maxEntrySize)
	}

	logEntries, err := stream.LogClient.GetEntries(stream.Context, stream.first, core.LogID(last))
	if err != nil {
		return []ct.LogEntry{}, errors.New(ERROR_FAILED_TO_FETCH_STH)
	}

	nextFirst := logEntries[len(logEntries)-1]
	stream.first = int64(nextFirst.Index + 1)

	return logEntries, nil
}

func (stream *CTClient) Next(callback core.Callback) {
	logEntries, err1 := stream.next()

	for _, entry := range logEntries {
		cert, err2 := extractCertFromEntry(&entry)
		err := errors.Join(err1, err2)

		callback(cert, CTClientParams{
			Index:     entry.Index,
			LogClient: stream.LogClient,
		}, err)
	}
}
