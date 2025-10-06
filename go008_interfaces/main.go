package main

import "fmt"

// Interface:
// - tập các method không có body
// - không thể tạo instance trực tiếp của interface nhưng có thể tạo biến có type là interface đó :)

// declare an interface
type Shape interface {
	area() float64
	perimeter() float64
}

// định nghĩa struct
type Circle struct {
	radius float64
}
type Rectangle struct {
	width, height float64
}

// định nghĩa các method cho struct
func (c Circle) area() float64 {
	return c.radius * c.radius * 3.14
}
func (c Circle) perimeter() float64 {
	return 2 * c.radius * 3.14
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}
func (r Rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

// note: chỉ cần định nghĩa đầy đủ các method của interface thì struct tự động implements
// implement ngầm định!

// ===============================
// Một struct có thể implement nhiều interface
// -> yêu cầu: nếu các interface có method trùng tên thì phải có cùng danh sách tham số và kiểu trả về

// ===============================
// Một interface có thể implement/embedding interface khác theo 2 cách:
// 1. nhúng interface
/*
	type A interface {
		methods_of_A
	}

	type B inteface {
		A
		methods_of_B
	}
*/
// 2. Khai báo các method của interface A trong interface B
/*
	type A interface {
		A_method1()
		A_method2()
	}

	type B inteface() {
		A_method1()
		A_method2()
		B_methods()
	}
*/
// tương tự struct, interface có thể embedded nhiều interface khác.

// =========================================
// Polymorphis : đạt được đa hình bằng cách tạo nhiều struct
// - các struct cùng implement interface chứa method muốn overloading
// - mỗi struct tự định nghĩa body của method đó

func main() {
	// tạo biến kiểu Shape
	var s Shape

	// Dynamic value: an interface-typed variable can hold any value (bất kỳ type nào) :)
	// chỉ cần type đó implement các method của interface.
	// gán circle cho shape :)
	s = Circle{radius: 5}
	fmt.Println("Value of s:", s)
	fmt.Printf("Type of s: %T\n", s)
	fmt.Println("s.area():", s.area())
	fmt.Println("s.perimeter():", s.perimeter())
	fmt.Println("===")
	// gán Rectangle cho shape
	s = Rectangle{width: 3, height: 4}
	fmt.Println("Value of s:", s)
	fmt.Printf("Type of s: %T\n", s)
	fmt.Println("s.area():", s.area())
	fmt.Println("s.perimeter():", s.perimeter())

	fmt.Println()
	fmt.Println("============================")
	fmt.Println()

	// type assertion
	fmt.Println("getArea(0):", getArae(0))
	fmt.Println("getArea(Circle{radius: 3):", getArae(Circle{radius: 3}))
	fmt.Println("getArea(Rectangle{width: 1, height: 2}):", getArae(Rectangle{width: 1, height: 2}))

	fmt.Println()
	fmt.Println("============================")
	fmt.Println()
	// Type switch
	fmt.Println(typeSwitch(Circle{5}))
	fmt.Println(typeSwitch(Rectangle{2, 5}))
	fmt.Println(typeSwitch(30))
	fmt.Println(typeSwitch("Hello Go"))
}

// Type assertion: Ép kiểu từ interface về kiểu gốc
// interfaceVar.(Type) -> ép kiểu của biến interfaceVar về kiểu Type (interface, struct)
// trả về value , ok (flag)
func getArae(i interface{}) float64 { // i interface{} // chứa bất kì giá trị nào
	// fmt.Printf("values: %s, type: %T\n", i, i)

	// ép kiểu về Shape
	shape, ok := i.(Shape) // nếu không có flag, ép kiểu sai gây panic

	if ok { // ép thành công
		return shape.area()
	} else { // ép thất bại
		return -1
	}
}

func typeSwitch(i interface{}) string { // truyền vào bất cứ cái gì
	switch i.(type) { // v = i.(type) // trả về kèm value
	case Circle:
		return "Circle"
	case Rectangle:
		return "Rectangle"
	case int:
		return "int"
	default:
		return "Unknown type"
	}
}
