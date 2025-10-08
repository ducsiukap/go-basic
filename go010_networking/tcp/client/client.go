package main

import (
	"fmt"
	"net"
)

func main() {

	// tao ket noi toi server
	// net.Dial(network, "host:port")
	conn, err := net.Dial("tcp", "localhost:1303")

	// error
	if err != nil {
		fmt.Println("Connection error:", err)
		return
	}

	// no error
	defer conn.Close()

	// send msg to server
	_, err = conn.Write([]byte("hello server!"))
	if err != nil {
		fmt.Println("Write error:", err)
		return
	}

	// receive from server
	buf := make([]byte, 2048)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}
	fmt.Println("Receive:", string(buf[:n]))
}
