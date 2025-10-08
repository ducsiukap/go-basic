package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/quic-go/quic-go/http3"
)

func main() {
	// tls config
	certPool := x509.NewCertPool()
	cert, err := os.ReadFile("D:/Programming/basics/Golang/go010_networking/http/server/tls_certs/server.crt")
	if err != nil {
		fmt.Printf("loading cert error: %v\n", err)
		return
	}
	certPool.AppendCertsFromPEM(cert)
	tlsConf := &tls.Config{
		RootCAs:    certPool,
		MinVersion: tls.VersionTLS13,
	}

	// transport
	transport := &http3.Transport{
		TLSClientConfig: tlsConf,
	}
	// giải phóng các kết nối QUIC trước khi kết thúc
	defer transport.Close()

	// create http client
	client := &http.Client{
		Transport: transport,
	}

	req, _ := http.NewRequest("GET", "https://localhost:1304/", nil)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("request failed: %v", err)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("request failed: %v", err)
		return
	}

	fmt.Println("status code:", resp.StatusCode)
	fmt.Println("body:", string(body))
}
