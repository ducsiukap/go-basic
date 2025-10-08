package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// create certificate pool
	certPool := x509.NewCertPool()
	// load certificate
	serverCert, err := os.ReadFile("D:/Programming/basics/Golang/go010_networking/http/server/tls_certs/server.crt")

	if err != nil {
		fmt.Printf("Error when read TLS Certificates: %v\n", err)
		return
	}
	// add certificate to cert pool
	certPool.AppendCertsFromPEM(serverCert)

	// config TLS
	tlsConf := &tls.Config{
		RootCAs:    certPool, // danh sách chứng chỉ gốc (Certificate Authorities) mà client tin tưởng để xác minh server
		MinVersion: tls.VersionTLS13,
	}

	// http.Transport
	transport := &http.Transport{
		TLSClientConfig: tlsConf, // add TLS config, optional

		// disable http2
		// ForceAttemptHTTP2: false,
	}

	// create client
	client := &http.Client{
		Transport: transport,
	}

	// send request
	resp, err := client.Get("https://localhost:1303/")
	if err != nil {
		fmt.Printf("Request failed: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Read response error: %v\n", err)
		return
	}
	fmt.Println("Status code:", resp.StatusCode)
	fmt.Println("body:", string(body))
}
