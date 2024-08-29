package ctstream

import (
	"context"
	"fmt"
	"net/http"

	"errors"

	"github.com/google/certificate-transparency-go/client"
	"github.com/google/certificate-transparency-go/jsonclient"
	ctx509 "github.com/google/certificate-transparency-go/x509"
)

type LogID = int64

var DefaultMaxEntries int64 = 32
var FetchingThread = 0

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

func (stream *CTClient) Next(callback Callback) error {
	sct, err := stream.LogClient.GetSTH(stream.Context)
	if err != nil {
		return errors.New(ERROR_FAILED_TO_FETCH_STH)
	}

	treeSize := int64(sct.TreeSize)

	if treeSize == stream.first {
		return errors.New(ERROR_NEW_LOGS_NOT_FOUND)
	}

	var last int64
	var first int64

	for first = stream.first; last < treeSize; first += stream.maxEntrySize {
		last = first + stream.maxEntrySize

		if treeSize < last {
			last = treeSize
		}

		FetchingThread++

		go stream.fetchEntries(first, last, callback)
	}

	stream.first = treeSize

	return nil
}

func (stream *CTClient) fetchEntries(first, last int64, callback Callback) error {
	entries, err1 := stream.LogClient.GetEntries(stream.Context, first, last)
	if err1 != nil {
		FetchingThread--
		return errors.New(ERROR_FAILED_TO_FETCH_STH)
	}

	for _, entry := range entries {
		cert, err2 := extractCertFromEntry(&entry)
		err := errors.Join(err1, err2)

		callback(cert, entry.Index, stream.LogClient, err)
	}

	FetchingThread--
	fmt.Printf("Thread: %d\n", FetchingThread)
	return nil
}
