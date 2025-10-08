package main

import (
	"fmt"
	"net"
)

func main() {
	client, err := net.Dial("udp", "localhost:1303")
	if err != nil {
		fmt.Println("Lỗi Dial:", err)
		return
	}
	defer client.Close()

	msg := "hello server!"
	fmt.Println("send:", msg)

	// gửi tới server
	_, err = client.Write([]byte(msg))
	if err != nil {
		fmt.Println("Lỗi gửi datagram packet tới server:", err)
		return
	}

	// nhận từ server
	buf := make([]byte, 1024)
	n, err := client.Read(buf)
	if err != nil {
		fmt.Println("Lỗi khi đọc phản hồi từ server!")
		return
	}
	fmt.Println("Receive:", string(buf[:n]))
}
