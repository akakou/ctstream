package sslmate

import (
	"crypto/x509"
	"errors"
	"fmt"

	"github.com/akakou/ctstream/core"
	"github.com/akakou/sslmate-cert-search-api/api"
)

type SSLMateCTClient struct {
	Api    *api.SSLMateSearchAPI
	Domain string
	First  string
}

func NewCTClient(domain string, api *api.SSLMateSearchAPI, first string) (*SSLMateCTClient, error) {
	return &SSLMateCTClient{
		Domain: domain,
		Api:    api,
		First:  first,
	}, nil
}

func DefaultCTClient(domain string) (*SSLMateCTClient, error) {
	return &SSLMateCTClient{
		Domain: domain,
		Api:    api.Default(),
	}, nil
}

func (client *SSLMateCTClient) next() ([]x509.Certificate, *api.Index, error) {
	query := api.Query{
		Domain:            client.Domain,
		IncludeSubdomains: true,
		MatchWildcards:    true,
		After:             client.First,
		Expand:            "",
	}

	certs, last, err := client.Api.Search(&query)
	if err != nil {
		return nil, nil, fmt.Errorf("%v: %v", ErrrorFailedToSearch, err)
	}

	client.First = last.Last
	return certs, nil, nil
}

func (client *SSLMateCTClient) Init() error {
	_, _, err := client.next()
	return err
}

func (client *SSLMateCTClient) Next(callback core.Callback) {
	certs, index, err1 := client.next()

	for _, cert := range certs {
		cert, err2 := reformatCertificate(&cert)
		err := errors.Join(err1, err2)
		callback(cert, index, err)
	}
}