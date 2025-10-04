package main

import (
	"fmt"
	go003 "function/init_function" // init() of go003 will be called automatically
)

// declare a function
// func function_name(args type) return_type { // executions }
func add(a int64, b int64) int64 {
	return a + b
}

// variadic function (varargs)
// can mix variadict & regular params but variadic must at the end of arg list.
func add_multi(nums ...int64) int64 {
	var total int64 = 0
	for _, num := range nums {
		total += num
	}
	return total
}

// reference arg
func swap(a *int, b *int) {
	c := *a
	*a = *b
	*b = c

}

// return multiple value
//
//	func funcname(args types) (type1, type2, type3, ...) {}
//
// or
//
//	func funcname(args types) (var1 type1, ...) {}
//
// or
//
//	func funcname(args types) (var ...T)
func add_sub(a, b int) (add int, sub int) {
	add = a + b
	sub = a - b
	return add, sub
}

func main() {

	// ==============================
	fmt.Println("==============================")
	go003.Sayhello()
	fmt.Println()

	// ==============================
	// call
	fmt.Println("==============================")
	fmt.Printf("sum of 3 and 5 is %d\n", add(3, 5))
	fmt.Println()

	// ==============================
	// call variadic function
	fmt.Println("==============================")
	fmt.Printf("sum of 1, 2 and 3 is %d\n", add_multi(1, 2, 3))
	fmt.Printf("sum of 13, 99, -1, 32, 57 is %d\n", add_multi(13, 99, -1, 32, 57))
	fmt.Println()

	// ==============================
	// reference arg
	fmt.Println("==============================")
	a, b := 3, 5
	fmt.Printf("before: a = %d, b = %d\n", a, b)
	swap(&a, &b)
	fmt.Printf("after: a = %d, b = %d\n", a, b)
	fmt.Println()

	// ==============================
	// Annonymous function
	fmt.Println("==============================")
	add_var := func(varargs ...int) int {
		total := 0
		for _, n := range varargs {
			total += n
		}
		return total
	}
	// call
	fmt.Printf("Sum of 9, 10, 11 is %d\n", add_var(9, 10, 11))
	fmt.Printf("sum of 1, 2, 3, 4, 5 is %d\n", add_var(1, 2, 3, 4, 5))
	fmt.Println()
	// can return annonymous function from another function
	// func f1(args types) func(args2 types2) type_of_annonymous_func { execution }

	// ==============================
	// multiple return
	fmt.Println("==============================")
	add, sub := add_sub(5, 3) // destructure
	fmt.Printf("add_sub(5, 3) -> (%d, %d)\n", add, sub)
	fmt.Println()

	// ==============================
	// defer hoãn thực hàm được gọi sau defer cho tới khi hàm bao gần nhất kết thúc
	//
	fmt.Println("==============================")
	test_defer := func() {
		fmt.Println("start....")

		// defer
		defer fmt.Println("Defer1")
		defer fmt.Println("Defer2")
		// nhiều defer -> sử dụng cơ chế viết sau thực hiện trước (LIFO)

		fmt.Println("end...")
	}
	test_defer()
}
