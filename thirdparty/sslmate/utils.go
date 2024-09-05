package sslmate

import (
	"crypto/x509"

	ctx509 "github.com/google/certificate-transparency-go/x509"
)

func reformatCertificate(cert *x509.Certificate) (*ctx509.Certificate, error) {
	return ctx509.ParseCertificate(cert.Raw)
}
