package main

import (
	"fmt"
	"net"
	"os"
)

// handler xử lý request
func handlerTCPRequest(conn net.Conn) {
	defer conn.Close() // đóng kết nối sau khi request thực hiện xong

	// xử lý byte
	buffer := make([]byte, 2048)
	n, e := conn.Read(buffer)

	// đọc lỗi
	if e != nil {
		fmt.Println("Lỗi khi đọc request:", e)
		return
	}

	// xử lý request
	msg := string(buffer[:n])
	client := conn.RemoteAddr()
	fmt.Println("Receive:\"", msg, "\"from client: ", client)

	// gửi lại cho client
	res := "Hello " + client.String()
	buffer = []byte(res)
	conn.Write(buffer)
	fmt.Println("Send:", res)
}

func main() {
	// Tạo tcpServer
	// net.Listen(network, "host:port")
	// network canbe: tcp, tcp4, tcp6, unix, unixpacket
	address := ":1303" // lắng nghe trên cổng 1303 của mọi network interface
	server, err := net.Listen("tcp", address)

	// lỗi khi lắng nghe
	if err != nil {
		fmt.Println("Error when trying listen on", address, ":", err)
		os.Exit(1)
	}

	// tạo server thành công
	fmt.Println("Server is running on", address)
	// tạo infinite loop để lắng nghe liên tục
	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("Error when accept the connection:", err)
		}

		// handle connection
		go handlerTCPRequest(conn)
	}

}
