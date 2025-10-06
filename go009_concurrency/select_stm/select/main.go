package main

import (
	"fmt"
	"time"
)

func task1(ch chan string) {
	// do smth
	time.Sleep(2 * time.Second)
	// done
	ch <- "Task-1 done!"
}

func task2(ch chan string) {
	// do smth
	time.Sleep(2 * time.Second)
	// done
	ch <- "Task-2 done!"
}

func main() {

	// var ch1 chan string
	// var ch2 chan string
	ch1 := make(chan string)
	ch2 := make(chan string)

	go task1(ch1)
	go task2(ch2)

	// select nhận từ channel
	// nhận được message từ channel nào
	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	default: // thêm default để tránh deadlock
		fmt.Println("No task is finished!")
	}

	// select gửi
	// case ch <- data: // do smth

}
