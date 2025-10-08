package http3server

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"github.com/quic-go/quic-go/http3"
)

func request_handler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	switch req.Method {
	case http.MethodGet:
		// res.WriteHeader(200) // default
		fmt.Fprintf(res, "Hello client!")
	default:
		res.WriteHeader(http.StatusMethodNotAllowed) // 405
		fmt.Fprintf(res, "Phương thức HTTP '%s' không được hỗ trợ tại endpoint này.", req.Method)
	}
}

func Server() *http3.Server {
	route := http.NewServeMux()
	route.HandleFunc("/", request_handler)

	// TLSConfig
	tlsConf := &tls.Config{
		// NextProtos: []string{"h2", "http/1.1"},
		MinVersion: tls.VersionTLS13,
	}

	cert, err := tls.LoadX509KeyPair("./tls_certs/server.crt", "./tls_certs/server.key")
	if err != nil {
		log.Fatalf("failed to load cert: %v", err)
		// log.Print()
		// log.Fatal(), log.Fatalf... -> log + os.Exit(1)
		// log.Panic(...) -> ném exception
	}
	tlsConf.Certificates = []tls.Certificate{cert}

	server := &http3.Server{
		Addr:      ":1304",
		Handler:   route,
		TLSConfig: tlsConf,
	}

	return server
}
