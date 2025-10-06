package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// WaitGroup là một cơ chế đồng bộ trong Go (package sync)
// Dùng để đợi một hoặc nhiều goroutine hoàn thành trước khi main goroutine hoặc goroutine khác tiếp tục
// Thay vì dùng time.Sleep() để “chờ” – đây là cách chuẩn và an toàn

// tạo goroutine để thêm vào wg
func handler(id int, works chan int, results chan int, wg *sync.WaitGroup) {
	defer wg.Done() // báo wg đã hoàn thành xong 1 goroutine // giảm đi 1 goroutine trong wg

	for v := range works {
		fmt.Println("handler:", id, " - work:", v)
		results <- v * v
	}
	fmt.Printf("---------------- handler - %d done! ----------------\n", id)
}

// Worker Pool = một tập hợp các worker (goroutine) được tạo sẵn để xử lý nhiều task cùng lúc.

// Cấu trúc cơ bản worker pool
// Job queue (channel): nơi gửi các task cần xử lý
// Workers (goroutine): lấy task từ job queue, xử lý, gửi kết quả
// Result channel (tuỳ chọn): nhận kết quả xử lý
// WaitGroup / sync: đợi tất cả worker xong

func main() {
	works := make(chan int, 5) // buffer chứa tối đa 5 phần tử
	results := make(chan int)

	var wg sync.WaitGroup

	// thêm các goroutine vào wg
	for i := 0; i < 5; i++ {
		wg.Add(1) // tăng số goroutine đang đợi trong wg lên 1
		go handler(i, works, results, &wg)
	}

	// gửi dữ liệu vào works
	for i := 0; i < 5; i++ {
		works <- rand.Intn(100)
	}
	// đóng channel work
	close(works) // dừng tất cả các hàm for v:= range works

	go func() {
		wg.Wait()      // đợi tất cả wg thực hiện xong
		close(results) // đóng result
	}()

	for v := range results {
		fmt.Println("Received", v, "from results channel!")
	}

}
