package utils

import (
	core "github.com/akakou/ctstream/core"
	ctx509 "github.com/google/certificate-transparency-go/x509"
)

func Callbacks(certs []*ctx509.Certificate, params any, callback core.Callback, err error) {
	for _, c := range certs {
		callback(c, params, err)
	}
}
