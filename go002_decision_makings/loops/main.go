package main

import "fmt"

func main() {
	// Normal `for` loop
	// for initializer; condition; post { execution }
	fmt.Println("Count from 0 to 5: ")
	for i := 0; i < 6; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	fmt.Println("===================================")

	// Infinite loop:
	// for {execution}

	// for loop as while loop
	// for condition { execution }
	fmt.Println("Count from 0 to 5: ")
	cur := 0
	for cur < 6 {
		fmt.Printf("%d ", cur)
		cur++
	}
	fmt.Println()
	fmt.Println("===================================")

	// loop in iterable
	// for i, v := range rvariable { execution}
	rvariable := []int{0, 1, 2, 3, 4, 5}
	for index, value := range rvariable {
		fmt.Printf("At index %d, value is %d\n", index, value)
	}
	fmt.Println("===================================")

	// for in string
	// for index, chr := range stringvar
	myname := "vduczz"
	for index, chr := range myname {
		fmt.Printf("character at index %d is %c\n", index, chr)
	}
	fmt.Println("===================================")

	// for in map...
	// for in channel...

	// với for in range
	// - lấy 2 tham số : i, j := range iterable -> trả về index + value
	// - lấy 1 tham số: i := range interable -> trả về index :)
	// => nếu chỉ muốn lấy value: for _, v := range smth
	hello := "こんにちは"
	for _, chr := range hello {
		// note: chr is rune (unicode)
		fmt.Printf("%c", chr)
	}
	fmt.Println("\n===================================")

	// =====================================

	// Loop control
	// - break : break out the nearest outer loop
	// - continue : skip this loop, go to next step
	// - goto : go back to a label

	// count from 0 to 5 with goto
	cur = 0

Counting: // label, can go back to here with goto Counting
	fmt.Printf("%d ", cur)
	cur++

	if cur < 6 {
		goto Counting
	}

}
