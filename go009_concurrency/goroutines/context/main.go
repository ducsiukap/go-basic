package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Context: là nhóm công việc liên quan đến nhau / cần hủy đồng thời
// 2 mục đích sử dụng:
// 		- Kiểm soát đồng thời (Concurrency Control):
// 			Cho phép các goroutine con nhận biết khi nào goroutine cha muốn dừng hoạt động.
// 			Điều này giúp tránh lãng phí tài nguyên và rò rỉ goroutine (goroutine leak)
// 		- Truyền giá trị (Value Propagation):
// 			Cho phép truyền dữ liệu theo phạm vi yêu cầu (request-scoped data) giữa các hàm hoặc goroutine
// 			mà không cần thay đổi chữ ký hàm (function signature) một cách rối rắm.
// 			Thường dùng để truyền các thông tin như ID yêu cầu, thông tin xác thực,
// 			hoặc các giá trị theo dõi (tracing values).

// Methods:
// 	- Deadline() -> trả về timeout + 1 giá trị boolean cho biết context này có được đặt timeout?
// 	- Done() -> trả về 1 kênh chỉ đọc <-chan struct{}
// 				kênh này sẽ bị đóng khi Context bị hủy bỏ hoặc hết hạn
// 	- Err() -> lý do Done()
// 	- Value(key interface{}) -> truy xuất value ứng với key

// định nghĩa các key cho context.WithValue
// thực tế nên tạo context_keys.go chỉ để lưu các biến khóa context
var k = contextKey{}

func main() {
	// Context gốc:
	// 		context.Background() // Context gốc, không bị hủy bỏ, không có giá trị hay thời hạn.
	// 								Thường được dùng ở hàm main, khởi tạo,
	// 								hoặc làm Context cấp cao nhất cho một yêu cầu đến.
	// 		context.TODO()

	// Sinh context
	// - Cách 1
	// 		context.WithCancel(parent) (context, cancelFunc)
	// 							gọi cancelFunc() sẽ hủy toàn bộ context được tạo từ context hiện tại.
	ctx, cancel := context.WithCancel(context.Background())
	// defer cancel() // thường sử dụng

	// kết hợp với WaitGroup
	var wg sync.WaitGroup

	wg.Add(1)
	go task_1(ctx, &wg)
	wg.Add(1)
	go task_2(ctx, &wg)
	// cancel sau 5s
	time.Sleep(5 * time.Second)
	fmt.Println("ctx is being cancel()!")
	cancel()
	// chờ wg
	wg.Wait()

	fmt.Println()
	fmt.Println("==========================")
	fmt.Println()

	// - Cách 2:
	// 		context.WithTimeout(parent, timeout) -: tự động cancel sau khoảng thời gian nhất định.

	// - Cách 3:
	// 		context.WithDeadline(parent, deadline) -> tương tự timeout nhưng hết hạn vào 1 thời điểm deadline

	// - Cách 4:
	// 		context.WithValue(parent, key any, value any) -> tạo context với cặp key, value
	// 														 để truyền tới các hàm, goroutines sử dụng context này.

	// value
	u := user{name: "vduczz", age: 21}
	cwl := context.WithValue(context.Background(), k, u)

	// truyền cho context cho các method/goroutine trong cùng package/ package cấp dưới.
	readCWL(cwl)
}

// ========================== context.WithValue ===========================//
// struct rỗng được khuyến nghị dùng làm key cho context
type contextKey struct{}

// data để truyền
type user struct {
	name string
	age  int
}

func readCWL(cwl context.Context) {
	value := cwl.Value(k) // phải dùng đúng
	fmt.Println("Context content:", value)
}

// =========================== Context.WithCancel() ========================= //
// Các task con hoạt động theo task cha
func task_1(ctx context.Context, wg *sync.WaitGroup) {
	// giảm wg
	defer wg.Done()

	cnt := 1

	// task 1 lặp vô hạn
	for {
		select {
		// chờ tín hiệu cancel() từ ctx
		case <-ctx.Done(): // nhận được tín hiệu cancel
			fmt.Println("task-1 nhận được tín hiệu cancel từ ctx!")
			return // dừng hàm

		// thực hiện việc khác
		default:
			fmt.Println("task-1 count:", cnt)
			time.Sleep(1 * time.Second)
		}
	}
}

func task_2(ctx context.Context, wg *sync.WaitGroup) {
	// giảm wg
	defer wg.Done()

	cnt := 1

	// tiếp tục tạo context con dựa trên ctx
	// để truyền hàm cấp thấp hơn
	// => lan truyền context
	ctx1, cancel := context.WithCancel(ctx)

	// thực hiện task3
	go task_3(ctx1)

	// thwucj hiện task2 song song
	for {
		select {
		// chờ tín hiệu cancel() từ ctx
		case <-ctx.Done(): // nhận được tín hiệu cancel
			fmt.Println("task-2 nhận được tín hiệu cancel từ ctx!")
			// lan truyền tới task3
			cancel()
			return // dừng hàm

		// thực hiện việc khác
		default:
			fmt.Println("task-2 count:", cnt)
			cnt++
			time.Sleep(1 * time.Second)
		}
	}
}

func task_3(ctx context.Context) {
	cnt := 1

	// task 1 lặp vô hạn
	for {
		select {
		// chờ tín hiệu cancel() từ ctx
		case <-ctx.Done(): // nhận được tín hiệu cancel
			fmt.Println("task-3 nhận được tín hiệu cancel từ ctx1!")
			return // dừng hàm

		// thực hiện việc khác
		default:
			fmt.Println("task-3 count:", cnt)
			cnt++
			time.Sleep(1 * time.Second)
		}
	}
}
