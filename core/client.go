package core

import (
	ctx509 "github.com/google/certificate-transparency-go/x509"
)

type LogID = int64

var DefaultMaxEntries int64 = 256

type Callback func(*ctx509.Certificate, int, any, error)

type CtClient interface {
	Init() error
	Next(callback Callback)
}
