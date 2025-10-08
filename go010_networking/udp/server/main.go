package main

import (
	"fmt"
	"net"
)

func main() {
	host, port := "", "1303"
	serverAdd, err := net.ResolveUDPAddr("udp", host+":"+port)
	// Phân rã lỗi
	if err != nil {
		fmt.Println("Lỗi khi phân rã địa chỉ: ", host, ":", port)
		return
	}

	// chờ kết nối
	conn, err := net.ListenUDP("udp", serverAdd)
	if err != nil {
		fmt.Println("Lỗi khi khởi tạo server tại ", serverAdd)
	}
	// close the connection when the main end
	defer conn.Close()

	fmt.Println("Server is running on ", serverAdd)

	// server loop
	buf := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Lỗi khi đọc datagram:", err)
			continue
		}
		go reqHandler(conn, addr, string(buf[:n]))
	}
}

// handler
func reqHandler(conn *net.UDPConn, addr *net.UDPAddr, req string) {
	fmt.Println("Receive:", req, "from client", addr)

	_, err := conn.WriteToUDP([]byte("Hello "+addr.String()), addr)
	if err != nil {
		fmt.Println("Lỗi khi gửi datagram tới", addr.String())
	}
}
