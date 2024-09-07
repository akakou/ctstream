package utils

import (
	core "github.com/akakou/ctstream/core"
	ctx509 "github.com/google/certificate-transparency-go/x509"
)

func Callbacks(certs []*ctx509.Certificate, params any, callback core.Callback, err error) {
	for i, c := range certs {
		callback(c, i, params, err)
	}
}
