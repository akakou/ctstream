package main

import (
	"fmt"
	"os"
	"time"

	ctstream "github.com/akakou/ctstream"
	ct "github.com/google/certificate-transparency-go"
	"github.com/google/certificate-transparency-go/client"
	ctx509 "github.com/google/certificate-transparency-go/x509"
)

func main() {
	m, err := ctstream.DefaultCTsStream([]string{
		"https://oak.ct.letsencrypt.org/2024h2/",
		"https://mammoth2024h2.ct.sectigo.com/",
	})

	if err != nil {
		fmt.Printf("Failed to create new ctstream")
		os.Exit(1)
	}

	err = m.Init()
	if err != nil {
		fmt.Printf("Failed to initialize ctstream")
		os.Exit(1)
	}

	m.Run(func(cert *ctx509.Certificate, i ctstream.LogID, c *client.LogClient, err error) {
		if err != nil {
			fmt.Printf("Failed to fetch %v: \n", err)
		}

		fmt.Printf("%d, %s\n", i, cert.DNSNames)
		fmt.Printf("%v%v?start=%v&end=%v\n\n", c.BaseURI(), ct.GetEntriesPath, i, i)
	}, 1000*time.Millisecond)
}
