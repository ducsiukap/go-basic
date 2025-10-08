#

## Create HTTP client

- Initialize HTTP client with `net/http.Client`

  ```go
  // structure of http.Client
  type Client struct {
    Transport RoundTripper // default is http.DefaultRoundTripper
                           // send & receive request/response

    CheckRedirect func(req *Request, via []*Request) error // callback khi gặp redirect (status code 3xx)

    Jar CookieJar

    Timeout time.Duration // maximum waiting time per request
  }

  // init client
  client := http.Client {
    // options
  }
  ```

- TLS config for HTTPS Client.

  ```go
  import (
  	"crypto/tls"
  	"crypto/x509"
  	"fmt"
  	"os"
  )

    // create certificate pool
  	certPool := x509.NewCertPool()
  	// load certificate
  	serverCert, err := os.ReadFile("/link/to/server.cert")
  	if err != nil {
  		fmt.Printf("Error when read TLS Certificates: %v\n", err)
  		return
  	}
  	// add certificate to cert pool
  	certPool.AppendCertsFromPEM(serverCert)

  	// config TLS
  	tlsConf = &tls.Config{
  		RootCAs:    certPool, // danh sách chứng chỉ gốc (Certificate Authorities) mà client tin tưởng để xác minh server
  		MinVersion: tls.VersionTLS13,
  	}
  ```

- Config client to support `HTTP/1.1` / `HTTP/2`:

  ```go
  // http.Transport
  transport := &http.Transport{
  	TLSClientConfig: tlsConf, // add TLS config, optional

    // http.Transport supports HTTP/1.1 & HTTP/2 by default
  	// to disable http2 -> only support http/1.1
  	ForceAttemptHTTP2: false,
  }

  // create client
  client := &http.Client{
  	Transport: transport,
  }
  ```

- Config client to support `HTTP/3`:

  - first, run this command in `cmd`: `go get github.com/quic-go/quic-go/http3`
  - then, create http3 client:

    ```go
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
    ```

#

## Send & Receive Request/Response

- Send/Receive simple request/response:

  ```go
    // Request
    // simple request using : client.Get, client.Post, client.Do
    response, err := client.Get(api_url)
    if err != nil {
        // handle request failed
        //panic(err)
    }
    defer resp.Body.Close()


    // Response
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }

    fmt.Println("Status code:", resp.StatusCode)
    fmt.Println("Response body:", string(body))

  ```

- Advanced request using `http.NewRequest("METHOD", "api_url", body io.Reader)`

  ```go
  // Add header
  req, _ := http.NewRequest("GET", "https://localhost:4433/", nil)
  req.Header.Set("User-Agent", "MyHTTP3Client/1.0")
  resp, err := client.Do(req)

  // Post with body
  //formdata
  data := strings.NewReader(`{"name":"ducsjukap"}`)
  // json
  data := bytes.NewBuffer([]byte(`{"name":"Táo","price":12.5}`))
  // create request
  req, _ := http.NewRequest("POST", "https://localhost:4433/data", data)
  req.Header.Set("Content-Type", "application/json")
  resp, err := client.Do(req)
  ```
