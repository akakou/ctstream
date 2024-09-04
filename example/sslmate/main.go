package main

import (
	"fmt"
	"os"
	"time"

	"github.com/akakou/ctstream/direct"
	"github.com/akakou/ctstream/thirdparty/sslmate"

	ctx509 "github.com/google/certificate-transparency-go/x509"
)

func main() {
	direct.DefaultSleep = time.Second * 10

	m, err := sslmate.DefaultCTsStream([]string{
		"test2.ochano.co",
	})

	if err != nil {
		fmt.Printf("Failed to create new ctstream: ")
		os.Exit(1)
	}

	err = m.Init()
	if err != nil {
		fmt.Printf("Failed to initialize ctstream: ")
		os.Exit(1)
	}

	m.Start(func(cert *ctx509.Certificate, option any, err error) {
		if err != nil {
			fmt.Printf("Failed to fetch %v: \n", err)
		}

		fmt.Printf("%v\n", cert.DNSNames)
	})

	m.Await()
}
