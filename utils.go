package ctstream

import (
	"errors"

	ct "github.com/google/certificate-transparency-go"
	ctx509 "github.com/google/certificate-transparency-go/x509"
)

func extractCertFromEntry(entry *ct.LogEntry) (*ctx509.Certificate, error) {
	var ctCert *ctx509.Certificate
	var err error

	switch entry.Leaf.TimestampedEntry.EntryType {
	case ct.X509LogEntryType:
		ctCert, err = entry.Leaf.X509Certificate()
	case ct.PrecertLogEntryType:
		ctCert, err = entry.Leaf.Precertificate()
	default:
		return nil, errors.New(NO_CERT_AND_PRECERT_LEAF)
	}

	if err != nil {
		return nil, errors.New(ERROR_FAILED_TO_GET_CERT)
	}

	return ctCert, err
}
