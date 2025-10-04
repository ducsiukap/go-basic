package main

import (
	exported "export/exported"
	"fmt"
)

// định nghĩa struct
/*
type struct_name struct {
	field1 dtype
	field2 dtype
	...
	fieldn dtype
}
*/

// add method to struct
// func (s struct_name) method_name(params types) { execution }

// READ IMPLEMENT IN ./exported/main.go

// Annonymous struct
// can be assigned to variable: variable := struct{}
type Song struct {
	author struct { // annonymous struct
		name string
		age  int
	}
	name string
}

// Nested struct, annonoymous field
type Author struct {
	name string
	age  int
}

type Book struct {
	// annonymous field
	// 1 type chỉ được dùng cho 1 annonymous field
	string // annonymous field // book's name
	Author // annonymous field // author
	// khi này, các field của Author sẽ thăng cấp lên thành field của Book :)
	// string // error
}

// Field/Method Promoted : Field/Method của nested struct sẽ được gắn vào lớp bao bên ngoài
// -> có thể trực tiếp truy cập thông qua outer.field

// Function field: -> field có kiểu là function
// func_field function(args)
// => call: struct_name.func_field(args)

func main() {
	// ==================================
	song := Song{
		author: struct {
			name string
			age  int
		}{name: "vduczz", age: 21},
		name: "Love",
	}
	fmt.Println("song's name: ", song.name)
	fmt.Println("song's author: ", song.author.name)
	fmt.Println()

	// ==================================
	book := Book{
		"Love",                          // book's name
		Author{name: "vduczz", age: 20}, // author
	}
	fmt.Println("book's name: ", book.string)
	fmt.Println("book's author: ", book.Author.name)
	fmt.Println()
	// book.Author.name = book.name :)
	// vì field của Author thăng cấp thành field của book, clm :)
	// nếu lớp ngoài đã có field trùng với lớp trong thì field lớp ngoài phủ lên field của lớp trong
	// => truy cập vào field của lớp trong theo cách truyền thống: outer.inner.field

	// ==================================
	// không thể tạo trực tiếp user từ exported.user{}
	// do khác package

	// tạo qua hàm tạo
	user := exported.CreateUser("xxxxxx", "vduczz", 21)

	// call smth public
	user.Introduction()
	fmt.Println("user's card number: ", user.GetIDCard())

	// cannot call setIDCard(id_card string) vì nó là private
	// user.setIDCard("123") // error

	// ==================================
	fmt.Println()
	// Struct equality
	// một struct có thể so sánh nếu các field của nó đều có dtype có thể so sánh
	// int, float, string, bool
	user2 := exported.CreateUser("xxxxxx", "vduczz", 21)
	user3 := exported.CreateUser("xxxxxx", "vduczz", 22)
	// sử dung == -> so sánh từng thuộc tính // field by field
	fmt.Println(user == user2)   // true
	fmt.Println((user == user3)) // false
}
