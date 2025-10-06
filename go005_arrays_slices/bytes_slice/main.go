package main

import (
	"bytes"
	"fmt"
)

func main() {
	//
	slc1 := []byte{'v', 'd', 'u', 'c', 'z', 'z'}
	slc2 := []byte{'h', 'e', 'l', 'l', 'o', 'G', 'o'}

	// Go cung cấp thư viện bytes để làm việc với byte

	// compare
	fmt.Println("bytes.Compare(slc1, slc2) :", bytes.Compare(slc1, slc2))
	fmt.Println("bytes.Equal(slc1, slc2) :", bytes.Equal(slc1, slc2))

	// sort
	// sort.Slice(slc1, func(i, j int) bool) //

	//trim
	// byte.Trim(b []byte, cutset string)
	// TrimSpace, TrimLeft, TrimRight
	// cutset: quét đồng thời 2 đầu (đầu & cuối chuỗi)
	// 			1 bên dừng khi gặp ký tự đầu tiên không nằm trong cutset
	slc := []byte("!Hello G!o")
	fmt.Printf("%s\n", bytes.Trim(slc, "!")) // Hello G!o

	// Split
	slc = []byte("Java C++ C,Go Python")
	parts := bytes.Split(slc, []byte(" ")) // Split(b []byte, delim []byte)
	for i, v := range parts {
		fmt.Printf("part %d, string: %s\n", i, v)
	}

}
