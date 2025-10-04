package go003

import "fmt"

// hàm init() được gọi khi package chứa nó được import vào trong package khác
// hoặc main() chạy nếu nó là package main
func init() {
	fmt.Println("init() of go003 package is called!")
}

func init() {
	fmt.Println("Another init() of go003 package is called!")
}

// nhiều hàm init -> gọi theo thứ tự khai báo

func Sayhello() {
	fmt.Println("go003 hello!")
}
