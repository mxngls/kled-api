package server

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
)

const (
	localCertFile = "bundle.cer"
)

func createClient() (client *http.Client, err error) {
	cert, err := ioutil.ReadFile(localCertFile)
	if err != nil {
		panic(err)
	}

	p := x509.NewCertPool()
	c, err := x509.ParseCertificates([]byte(cert))
	for _, n := range c {
		p.AddCert(n)
	}

	config := &tls.Config{
		RootCAs: p,
	}

	tr := &http.Transport{TLSClientConfig: config}
	client = &http.Client{Transport: tr}
	return client, err
}
