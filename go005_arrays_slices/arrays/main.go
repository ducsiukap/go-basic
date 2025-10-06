package main

import "fmt"

// Array: same type, immuable

func main() {
	// ===========================
	// create
	var arr [5]int
	// or
	arr2 := [...]int{1, 2, 3, 4, 5} // tự tính độ dài, arr2 sẽ có độ dài cố định sau khi tạo
	// arr2 = append(arr2, 5) //error
	fmt.Println("Create array: arr = ", arr)
	fmt.Println("Create array: arr2 = ", arr2)
	// khác với
	arr3 := []int{1, 2, 3, 4, 5} // slice
	// slice có kích thước có thể thay đổi
	arr3 = append(arr3, 100)
	fmt.Println("arr3 = ", arr3)

	// multi dimensions arr:
	// var arr [length1][length2].....[lengthn]type

	// get length
	fmt.Printf("len(arr) = %d, len(arr2) = %d\n", len(arr), len(arr2))
	fmt.Println()

	// ===========================
	// Accessing : using index (0 -> len(arr - 1)), for range
	// - update elems
	arr[0] = 99
	for i := 1; i < 5; i++ {
		arr[i] = i
	}
	// - read elems
	for index, value := range arr {
		fmt.Printf("index: %d, value: %d\n", index, value)
	}
	fmt.Println()

	// ===========================
	// copy arr
	// - assignment (only work with arr, not slice)
	otherArr := arr
	fmt.Println("Copied arr: ", otherArr)
	fmt.Println("other == source : ", otherArr == arr)
	fmt.Println("&other == &source : ", &otherArr == &arr) // not point to the same memory with source arr
	// - using pointer
	fmt.Println("====")
	pointerArr := &arr                                   // pointer -> *[]int
	fmt.Println("copied: ", *pointerArr)                 // pointerArr hold address, *pointerArr hold value
	fmt.Println("(*pointerArr)[1] = ", (*pointerArr)[1]) // access specified elem
	// - using for loop to copy elems
	fmt.Println()

	// ===========================
	// passsing arr to function
	fmt.Println("sum of elems in arr: ", arr_sum(arr))
	// pass slice // slice is a view that reference to array  (point to array)
	fill(arr[:], -1)
	fmt.Println("arr filled with value -1: ", arr)
}

// passing arr to function
// 1. -> array cố định kích thước : [N]type -> value, *[N]type -> reference
func arr_sum(arr [5]int) int {
	total := 0
	for _, v := range arr {
		total += v
	}
	return total
}

// 2. -> slice => không cố định kích thước // note: slice trỏ tới dữ liệu thât (reference)
func fill(arr []int, value int) {
	for i := range arr {
		arr[i] = value
	}
}
