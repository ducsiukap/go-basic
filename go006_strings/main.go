package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	// String: represent one or more bytes using UTF8-Encoding
	// - String Literal with ""
	str_literal := "Hello guys.\nGo is not funny! is string-literal in Go."
	// str_literal := "Hello guys.\nGo is not funny!
	//  is string-literal in Go." // error

	// - Raw literal // coi ký tự đặc biệt như ký tự thường, format tự định nghĩa (xuống dòng ở đâu, ... )
	raw_literal := `Hello guys.\nGo is not funny!
	This is raw-string in Go.`

	fmt.Printf("String literal:\n%s\n", str_literal)
	fmt.Printf("Raw Literal:\n%s\n", raw_literal)

	fmt.Println()
	fmt.Println("============================")
	fmt.Println()

	// Iterating overstring
	str := "hello go!"
	// - using for in
	for _, chr := range str {
		fmt.Printf("%c ", chr)
	}
	fmt.Println()
	// - using forloop // duyệt từng byte
	// dùng nếu string dạng UTF8-Encoder
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c ", str[i])
	}
	fmt.Println()
	// nếu string chứa rune
	str1 := "hello Go👋 !!!"
	for i := 0; i < len(str1); i++ {
		fmt.Printf("%c ", str1[i]) //h e l l o   G o ð      ! ! !
	}
	fmt.Println()

	fmt.Println()
	fmt.Println("============================")
	fmt.Println()

	// Create string from slice
	s1 := []byte("vduczz")
	s2 := []rune("vduczz👋")
	// using string([]T)
	fmt.Println(string(s1), string(s2))

	fmt.Println()
	fmt.Println("============================")
	fmt.Println()

	// Compare
	cmp1 := "Hello go"
	cmp2 := "Hello go"
	cmp3 := "Hello Golang"
	// using ==, !=, <, <=, >, >=,
	fmt.Println("cmp1 == cmp2 :", cmp1 == cmp2)
	fmt.Println("cmp1 == cmp3 :", cmp1 == cmp3)
	// using strings package
	fmt.Println("strings.Compare(cmp1, cmp2) : ", strings.Compare(cmp1, cmp2))
	fmt.Println("strings.Compare(cmp1, cmp3) : ", strings.Compare(cmp1, cmp3))

	fmt.Println()
	fmt.Println("============================")
	fmt.Println()

	// Concatenning
	concat1 := "hello"
	concat2 := "world"
	// using + or +=
	fmt.Println(concat1 + concat2)
	// using bytes.Buffer
	var byte_buf bytes.Buffer     // create buffer
	byte_buf.WriteString(concat1) // Write
	byte_buf.WriteString(concat2)
	fmt.Println(byte_buf.String())
	// using fmt.Sprintf, Sprint, Sprintln
	fmt.Println(fmt.Sprintf("%s %s!", concat1, concat2))
	// using strings.Join
	fmt.Println(strings.Join([]string{concat1, concat2}, " ")) // strings.Join(s []string, delim string)
	// using strings.Builder & WriteString
	var builder strings.Builder
	builder.WriteString(concat1)
	builder.WriteString(" ")
	builder.WriteString(concat2)
	fmt.Println(builder.String())

	fmt.Println()
	fmt.Println("============================")
	fmt.Println()

	// strings package methods for string

	// Trim(str string, cutset []byte)
	// TrimLeft, TrimRight, TrimSuffix, TrimPrefix
	// TrimSpace

	// Split(str string, delim string) // không lấy delim
	// SplitAfter(str string, delim string) // lấy cả delim ở mỗi token
	// SplitN, SplitAfterN // tách tối đa n substring

	// Contains(str string, substr string)
	test_str := "hello go"
	fmt.Println(strings.Contains(test_str, "go")) // true

	// Index(str string, substr string) // find index of substring
	// IndexByte(str string, chr byte)
	fmt.Println(strings.Index(test_str, "go"))

	// Repeat(str string, times int)
	test_str = "loading...\n"
	fmt.Println(strings.Repeat(test_str, 3))

	// Count(str string, substr string) // count number of repeated substr in str
	test_str = "go go. hello go"
	fmt.Println(strings.Count(test_str, "go"))

	// bytes, strings có những API rất giống nhau, chỉ khác chỗ:
	// - bytes thao tác trên []byte
	// - strings thao tác trên string
}
