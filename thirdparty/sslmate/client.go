package sslmate

import (
	"crypto/x509"
	"errors"
	"fmt"
	"time"

	"github.com/akakou/ctstream/core"
	"github.com/akakou/ctstream/utils"
	"github.com/akakou/sslmate-cert-search-api/api"
)

type SSLMateCTClient struct {
	Api    *api.SSLMateSearchAPI
	Domain string
	First  string
	Sleep  time.Duration
}

type SSLMateCTParams *api.Index

func NewCTClient(domain string, api *api.SSLMateSearchAPI, sleep time.Duration) (*SSLMateCTClient, error) {
	return &SSLMateCTClient{
		Domain: domain,
		Api:    api,
		First:  "",
		Sleep:  sleep,
	}, nil
}

func DefaultCTClient(domain string) (*SSLMateCTClient, error) {
	return &SSLMateCTClient{
		Domain: domain,
		Api:    api.Default(),
		Sleep:  DefaultPullingSleep,
	}, nil
}

func (client *SSLMateCTClient) next() ([]x509.Certificate, SSLMateCTParams, error) {
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
	return certs, last, nil
}

func (client *SSLMateCTClient) Init() error {
	return nil
}

func (client *SSLMateCTClient) Next(callback core.Callback) {
	l := 1

	for l != 0 {
		certs, index, err1 := client.next()
		formated, err2 := reformatCertificates(certs)

		err := errors.Join(err1, err2)
		utils.Callbacks(formated, index, callback, err)

		time.Sleep(DefaultPullingSleep)
		l = len(certs)
	}
}
