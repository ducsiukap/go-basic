package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Trong go, channel (chan) thường là kênh giao tiếp giữa các goroutines
func writer(ch chan int) {

	for i := 1; i < 10; i++ {
		x := rand.Intn(100)
		ch <- x
		fmt.Println("send", x, "to channel ch!")
		time.Sleep(300 * time.Millisecond)
	}

	fmt.Println("Close the channel!")
	close(ch) // close để đóng channel
	// để đọc dữ liệu từ channel, nên dùng
	// v, ok := <-ch
	// trả về cả flag, để tránh panic khi ch đã close.
}

func reader(ch chan int) {
	for v := range ch {
		fmt.Println("receive", v, "from channel ch!")
	}
	// for tự dừng khi channel close

	fmt.Println("channel ch is closed!")
}
func main() {

	ch := make(chan int)
	// or make(chan type, buffer int)

	// bình thường, 1 channel là 2 chiều (gửi và nhận)
	// có thể tạo channel chỉ gửi:
	// 		make(chan<- type)
	// hoặc chỉ nhận:
	// 		make(<-chan type)

	go writer(ch) // thêm writer(ch) vào Scheduler
	go reader(ch)

	time.Sleep(5 * time.Second)
}
