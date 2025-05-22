package sslmate

import (
	"fmt"
	"strconv"
	"time"

	"github.com/akakou/ctstream/core"
	"github.com/akakou/ctstream/utils"
	api "github.com/akakou/sslmate-cert-search-api"
)

type SSLMateCTClient struct {
	Api    *api.SSLMateSearchAPI
	Domain string
	First  int
	Sleep  time.Duration
}

type SSLMateCTParams struct {
	Index  *api.Index
	Client *SSLMateCTClient
}

func NewCTClient(domain string, api *api.SSLMateSearchAPI, sleep time.Duration) (*SSLMateCTClient, error) {
	return &SSLMateCTClient{
		Domain: domain,
		Api:    api,
		First:  0,
		Sleep:  sleep,
	}, nil
}

func DefaultCTClient(domain string) (*SSLMateCTClient, error) {
	return &SSLMateCTClient{
		Domain: domain,
		Api:    api.Default(),
		Sleep:  core.DefaultPullingSleep,
	}, nil
}

func (client *SSLMateCTClient) fetch() ([]api.SSLMateCertEntry, error) {
	query := api.Query{
		Domain:            client.Domain,
		IncludeSubdomains: true,
		MatchWildcards:    true,
		After:             strconv.Itoa(client.First),
		Expand:            "",
	}

	entries, err := client.Api.Search(&query)
	if err != nil {
		return nil, fmt.Errorf("%v: %v", ErrrorFailedToSearch, err)
	}

	client.First = entries[len(entries)-1].Id
	return entries, nil
}

func (client *SSLMateCTClient) fetchAll(callback core.Callback) []api.SSLMateCertEntry {
	entries := []api.SSLMateCertEntry{}

	l := -1
	for l != 0 {
		time.Sleep(core.DefaultPullingSleep)

		es, err := client.fetch()
		if err != nil {
			callback(nil, 0, nil, err)
			continue
		}

		entries = append(entries, es...)
		l = len(es)
	}

	return entries
}

func (client *SSLMateCTClient) Init() error {
	client.First = 0
	return nil
}

func (client *SSLMateCTClient) Next(callback core.Callback) {
	entries := client.fetchAll(callback)

	for _, entry := range entries {
		ctcert, err := utils.ReformatCertificate(entry.Cert)
		callback(ctcert, entry.Id, nil, err)
	}
}

func (client *SSLMateCTClient) GetDomain() string {
	return client.Domain
}
