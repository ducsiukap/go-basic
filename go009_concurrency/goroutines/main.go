package main

import (
	"fmt"
	"time"
)

// Goroutines trong go cho phép nhiều hàm chạy đồng thời.
// ~ Coroutines trong kotlin
// Là thread nhẹ (lightweight thread) do Go runtime quản lý, chứ không phải thread hệ điều hành.
// Có thể tạo hàng nghìn goroutine cùng lúc mà không tốn nhiều tài nguyên.

// Async task: chạy công việc tốn thời gian mà không chặn main
// Concurrent server: mỗi request xử lý trong goroutine riêng
// Pipeline: chia dữ liệu thành nhiều bước, goroutine + channel
// Worker pool: xử lý nhiều task nhỏ, giới hạn số worker
// Timer / định kỳ: task chạy lặp lại mà không chặn main

// định nghĩa hàm:
func log(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s, i)
		time.Sleep(200 * time.Millisecond) // sleep 0.5s
		// Sleep -> tạm dừng thread hiện tại
	}
}

func main() {
	fmt.Println("---------------- main-start ----------------")
	// chaỵ tuần tự
	log("mess-1")
	log("mess-2")

	fmt.Println()
	fmt.Println("=======================")
	fmt.Println()

	// chạy song song
	go log("goroutines-1") // tạo 1 thread
	go log("goroutines-2") // tạo 1 thread
	log("main-1")          // chạy trongg main thread
	log("main-2")          // chạy trong main thread
	// main-1, main-2 vẫn in tuần tự, đồng thời chạy song song với goroutines-1 và goroutines-2

	// Note: mọi thứ trong hàm main là chạy đồng bộ, tức là từ trên xuống.
	// lệnh "go funcname()" thực ra ghi lời gọi hàm vào Scheduler và chờ được chạy
	// vì vậy:
	// 		log(main-1)
	// 		go log(goroutines-1)
	// không chạy song song, chạy đúng thứ tự thực thi
	// vì log main-1 thực hiện xong mới tới goroutines-1 được thêm vào Scheduler.
	// còn
	// 		go log(goroutines-1)
	// 		log(main-1)
	// là chạy song song, goroutine-1 được thêm vào Scheduler (main luôn có trong scheduler)

	fmt.Println()
	fmt.Println("=======================")
	fmt.Println()

	// Annonoymous goroutines
	go func(s string) {
		for i := 0; i < 3; i++ {
			fmt.Println(s, i, "--- from annonymous goroutines")
			time.Sleep(100 * time.Millisecond)
		}
	}("hello")
	// time.Sleep(2 * time.Second) // chờ annonymous goroutines thực thi
	log("world!")

	fmt.Println("---------------- main-end ----------------")
	// Note: Goroutine kết thúc khi chương trình chính kết thúc
	// main() kết thúc -> toàn bộ goroutines đều bị kill
}
