package main

import (
	"fmt"
	"reflect"
	"slices"
	"sort"
)

// slice: flexible & efficient way to represent array
//  often used in place of arrays because of their dynamic size and added features

// Component of slice:
// - Pointer
// 		note: nếu slice tạo từ array -> slice trỏ tới đúng array đó
// 		mọi thay đổi đều được thực hiện cả trên array
// - Length
// - Capacity // maximum sizable of slice

func main() {
	// ===========================
	// Create slice
	// 1. similar to create an array without size specified
	var slice1 []int // or slice1 := []int{} // empty slice
	// or init value
	slice2 := []int{1, 2, 3, 4, 5}
	// or chỉ định chỉ số
	slice3 := []int{1: 99, 10: 17}
	// 2. Create from an existing array -> point to array likely pointer
	arr := [...]int{1, 2, 3, 4, 5}
	slice4 := arr[1:3]
	// array[startindex:endindex] -> return a view on array from startindex to (endindex-1)
	// :end -> 0 to end
	// from: -> from to len-1
	// : -> 0 to len-1
	fmt.Println("Create slice1: ", slice1)
	fmt.Println("Create slice2: ", slice2)
	fmt.Println("Create slice3: ", slice3)
	fmt.Println("Create slice4: ", slice4)
	fmt.Println()
	// 3. copy an slice
	var slice5 []int // null
	copy(slice5, slice4)
	// tuonwg tự: slice5 = append(slice5, slice4...)
	// 4. Using make
	// slice4 := make([]int, 10, 100) // make([]type, len, cap)

	// ===========================
	// Accessing
	// - using index
	for i := 0; i < len(slice2); i++ {
		fmt.Printf("index: %d, value: %d\n", i, slice2[i])
	}
	// - using for in range
	for _, v := range slice3 {
		fmt.Print(v, " ")
	}
	fmt.Println()
	fmt.Println()

	// ===========================
	// Sorting
	// sort.Types(slice)
	slc := []string{"c++", "python", "go", "java", "c"}
	sort.Strings(slc) // sort.Type([]type)
	fmt.Println(slc)
	//  checking slice is sorted
	fmt.Println("slc is sorted:", sort.StringsAreSorted(slc))
	// sort with user defined compare function
	sort.Slice(slc, func(i, j int) bool { return len(slc[i]) < len(slc[j]) }) // sort.Slice([]T, func(i, j int))// i, j là index :)
	fmt.Println("slc sorted by element length: ", slc)
	fmt.Println()

	// ===========================
	// Copying a slice
	// 1. copy from an array
	src := [...]string{"Hello", "i'm", "vduczz"}
	fmt.Println("the copied: ", src[:]) // array[from:to]
	// 2. copy a slice
	dst := make([]string, 4)
	// - using copy
	copy(dst, slc) // copy(dst, src) copy tối đa len(dst) phần tử từ src
	fmt.Println("the copied: ", dst)
	// - using append
	dst = []string{}
	dst = append(dst, slc...)
	fmt.Println("the copied: ", dst)
	// using forloop
	fmt.Println()

	// ===========================
	// Slice comparing
	// 1. using slices package
	// slices.Equal(), slices.Compare()...
	fmt.Println("slices.Equal(dst, slc): ", slices.Equal(dst, slc))        // boolean
	fmt.Println("slices.Compare(dst, src): ", slices.Compare(dst, src[:])) // int
	// slice.EqualFunc() // tự định nghĩa hàm so sánh
	fmt.Println("slices.EqualFunc: ", slices.EqualFunc([]int{1, 2, 3}, []int{1, 4, 9}, func(s1 int, s2 int) bool { return s1*s1 == s2 }))
	// 2. reflect.DeepEqual()
	fmt.Println("reflect.DeepEqual(dst, slc): ", reflect.DeepEqual(dst, slc)) // boolean
	// so sanh voi null
	fmt.Println("dst == nil : ", dst == nil)

}
