package ctstream

import (
	"context"
	"net/http"

	"errors"

	ct "github.com/google/certificate-transparency-go"
	"github.com/google/certificate-transparency-go/client"
	"github.com/google/certificate-transparency-go/jsonclient"
	ctx509 "github.com/google/certificate-transparency-go/x509"
)

type LogID = int64

var DefaultMaxEntries int64 = 256

type Callback func(*ctx509.Certificate, LogID, *client.LogClient, error)

type CTClient struct {
	Url string
	*client.LogClient
	context.Context
	first        LogID
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
	return NewCTClient(url, DefaultMaxEntries, jsonclient.Options{})
}

func (stream *CTClient) Init() error {
	sct, err := stream.GetSTH(stream.Context)
	if err != nil {
		return errors.New(ERROR_FAILED_TO_FETCH_STH)
	}

	stream.first = int64(sct.TreeSize)
	return nil
}

func (stream *CTClient) Next() ([]ct.LogEntry, error) {
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

	logEntries, err := stream.LogClient.GetEntries(stream.Context, stream.first, LogID(last))
	if err != nil {
		return []ct.LogEntry{}, errors.New(ERROR_FAILED_TO_FETCH_STH)
	}

	nextFirst := logEntries[len(logEntries)-1]
	stream.first = int64(nextFirst.Index + 1)

	return logEntries, nil
}
