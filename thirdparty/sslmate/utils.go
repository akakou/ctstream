package sslmate

import (
	"crypto/x509"
	"errors"

	"github.com/akakou/ctstream/core"
	ctx509 "github.com/google/certificate-transparency-go/x509"
)

func reformatCertificate(cert *x509.Certificate) (*ctx509.Certificate, error) {
	return ctx509.ParseCertificate(cert.Raw)
}

func reformatCertificates(certs []x509.Certificate) ([]*ctx509.Certificate, error) {
	var err error
	var result []*ctx509.Certificate

	for _, cert := range certs {
		c, tmpError := reformatCertificate(&cert)
		err = errors.Join(err, tmpError)
		result = append(result, c)
	}

	return result, err
}

func callbacks(certs []*ctx509.Certificate, callback core.Callback, err error) {
	for index, c := range certs {
		callback(c, index, err)
	}
}
