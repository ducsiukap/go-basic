// Go or Golang, open-source programming language developed by Google in 2007 and officially released in 2009
// Goal ->  combining :
// 	+ performance & security of a low-level language like C
// 	+ simplicity and productivity of modern languages like Python

// mọi chương trình .go phải có dòng khai báo package đầu tiên
// note: trong cùng 1 folder, tất cả file .go phải có cùng package nhưng chỉ cho phép 1 file có hàm main :) (package phải là main)
// main là package đặc biệt, chứa hàm main()
package main

// nếu viết thư viện thì: package libName

//  include neccessary external package
// fmt: format package -> for scan input, print output and processing formatted string
import "fmt"

// or multiple package
// import (
// 	"fmt"
// 	"pack2"
// 	"pack3"
// 	"pack4"
// )

// entry point of program
func main() {
	var name string

	// standart output to screen with fmt.Println()
	fmt.Print("Enter your name: ")

	// take input with fmt.Scan
	fmt.Scan(&name)

	fmt.Printf("Hi, %s\n", name)

	// single line comment
	/*
		multi-lines
		comment
	*/
}

// to run program:
// go run programName.go
