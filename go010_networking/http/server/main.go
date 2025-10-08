package main

import (
	"fmt"
	http2server "httpserverdemo/http1.1_http2"
	http3server "httpserverdemo/http3"
	"sync"
)

func main() {
	http2server := http2server.Server()
	http3server := http3server.Server()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("http2server running on", http2server.Addr)
		err := http2server.ListenAndServeTLS("", "")
		if err != nil {
			fmt.Printf("http2server error: %v\n", err)
			return
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("http3server running on", http3server.Addr)
		err := http3server.ListenAndServe()
		if err != nil {
			fmt.Printf("http3server error: %v\n", err)
			return
		}
	}()

	wg.Wait()

}
