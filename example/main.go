package main

import (
	"fmt"
	"os"
	"time"

	ctstream "github.com/akakou/ctstream"
	"github.com/google/certificate-transparency-go/client"
	ctx509 "github.com/google/certificate-transparency-go/x509"
)

func main() {
	m, err := ctstream.New([]string{
		"https://ct.googleapis.com/logs/us1/argon2024/",
		"https://ct.googleapis.com/logs/us1/argon2025h1/",
		"https://ct.googleapis.com/logs/us1/argon2025h2/",
		"https://ct.googleapis.com/logs/eu1/xenon2024/",
		"https://ct.googleapis.com/logs/eu1/xenon2025h1/",
		"https://ct.googleapis.com/logs/eu1/xenon2025h2/",
		"https://ct.cloudflare.com/logs/nimbus2024/",
		"https://ct.cloudflare.com/logs/nimbus2025/",
		"https://yeti2024.ct.digicert.com/log/",
		"https://yeti2025.ct.digicert.com/log/",
		"https://nessie2024.ct.digicert.com/log/",
		"https://nessie2025.ct.digicert.com/log/",
		"https://sabre.ct.comodo.com/",
		"https://sabre2024h1.ct.sectigo.com/",
		"https://sabre2024h2.ct.sectigo.com/",
		"https://sabre2025h1.ct.sectigo.com/",
		"https://sabre2025h2.ct.sectigo.com/",
		"https://mammoth2024h1.ct.sectigo.com/",
		"https://mammoth2024h1b.ct.sectigo.com/",
		"https://mammoth2024h2.ct.sectigo.com/",
		"https://mammoth2025h1.ct.sectigo.com/",
		"https://mammoth2025h2.ct.sectigo.com/",
		"https://oak.ct.letsencrypt.org/2024h1/",
		"https://oak.ct.letsencrypt.org/2024h2/",
		"https://oak.ct.letsencrypt.org/2025h1/",
		"https://oak.ct.letsencrypt.org/2025h2/",
		"https://ct2024.trustasia.com/log2024/",
		"https://ct2025-a.trustasia.com/log2025a/",
		"https://ct2025-b.trustasia.com/log2025b/",
	}, 10*time.Second)

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

		fmt.Printf("%v(%v): %v\n", c.BaseURI(), i, cert.DNSNames)
	})
}
