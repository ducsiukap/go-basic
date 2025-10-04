package main

import "fmt"

func main() {
	// =================================
	// Basic dtypes: Numbers, Booleans, Strings
	// =================================

	// String: UTF-8 string
	var name string
	fmt.Print("enter your name: ")
	fmt.Scan(&name)

	// Number -> Integer, Floating-point
	// Integer
	// 		int8 -> int64 : 1, 2, 4, 8 byte integer
	// 		uint8 -> uint64: usigned integer
	// 		int, uint -> 32 or 64 bit int/uint
	var age int
	fmt.Print("enter your age: ")
	fmt.Scan(&age)
	// 		rune %c : tương tự char nhưng dùng 32 bit (int32) biểu diễn Unicode code points
	// 		có thể biểu diễn MỌI ký tự trong unicode (Ascii, tiếng Việt, emoji)
	var smile rune = '😀'
	// 		byte : 8 bit (1 byte) integer (uint8)
	// 		uintptr : số nguyên đủ lớn để biểu diễn địa chỉ của con trỏ.
	// Floating-Point
	// 		float32 : float
	// 		float64 : double
	// Complex: -> real(), imag()
	// 		complex64 -> float32 for real & imag
	// 		complex128 -> float64 for real & imag

	// Boolean : 1bit for true/false
	fmt.Printf("hi, %s %d. let's keep your smile %c!", name, age, smile)

	// =================================
	// Aggregate type: Array, struct
	// =================================

	// =================================
	// Reference type: Pointers, slices, maps, functions & channels
	// =================================

	// =================================
	// Interface type
	// =================================
}

// doesn't support the Automatic Type Conversion or Implicit Type Conversion:P
// ex: var a int = 1
// var b float32 = a -> error
// var b float32 = 1.0*a -> error
// :)))))))))

// type casting using:
// newtype(value)
// -> cast value to new type
