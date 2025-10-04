package main

import "fmt"

type user struct {
	name string
	age  int
}

type product struct {
	name  string
	price float32
}

// method = gắn function vào type nào đó
// syntax:
// 		func (r receive_type) method_name() {}

// chỉ gắn được vào r là user-defined type
// không gắn được vào built-in type
// ex:
// gắn log vào user
func (u user) log() {
	fmt.Printf("user {\n\tname: %s, \n\tage: %d\n}\n", u.name, u.age)
}

// go cho phép 2 method có cùng tên nhưng buộc phải khác receive_type
func (p product) log() {
	fmt.Printf("product {\n\tname: %s, \n\tprice: %.2f\n}\n", p.name, p.price)
}
func main() {
	u := user{name: "vduczz", age: 21}
	p := product{name: "Laptop XXX", price: 799}

	// call method
	u.log()
	p.log()
}
