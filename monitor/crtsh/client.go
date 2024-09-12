package crtsh

import (
	"github.com/akakou/crtsh"
	"github.com/akakou/ctstream/core"
	"github.com/akakou/ctstream/utils"
)

type CrtshCTClient struct {
	Index  int
	Domain string
}

type CrtshCTParams struct {
	Index  int
	Client *CrtshCTClient
}

func NewCTClient(domain string) (*CrtshCTClient, error) {
	return &CrtshCTClient{
		Index:  0,
		Domain: domain,
	}, nil
}

func (client *CrtshCTClient) Init() error {
	_, err := crtsh.Fetch(client.Domain, crtsh.EXCLUDE_EXPIRED)
	return err
}

func (client *CrtshCTClient) Next(callback core.Callback) {
	entries, err := crtsh.Fetch(client.Domain, crtsh.EXCLUDE_EXPIRED)
	if err != nil {
		callback(nil, 0, &CrtshCTParams{}, err)
		return
	}

	for i, entry := range entries {
		c, err := utils.ReformatCertificate(entry.Certificate)
		callback(c, i, &CrtshCTParams{
			Index:  i,
			Client: client,
		}, err)
	}
}

func (client *CrtshCTClient) GetDomain() string {
	return client.Domain
}
