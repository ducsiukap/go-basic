package main

import "fmt"

func main() {
	// Arithmetic ops: +, -, *, /, %
	a := 100
	b := 36
	fmt.Printf("%d + %d=%d\n", a, b, a+b)
	fmt.Printf("%d - %d=%d\n", a, b, a-b)
	fmt.Printf("%d * %d=%d\n", a, b, a*b)
	fmt.Printf("%d / %d=%d\n", a, b, a/b)
	fmt.Printf("%d mod %d=%d\n", a, b, a%b)

	fmt.Println("===================================")

	// relational ops: ==, !=, >, >=, <, <=

	// logical ops: &&, ||, !

	// bitwise ops: &, |, ^ (xor / not), <<, >>
	x := 12                                 // 1100
	y := 5                                  // 0101
	fmt.Printf("^%d = %d\n", x, ^x)         // ^a -> not a  // ^n = -n - 1
	fmt.Printf("%d^%d = %d\n", x, y, x^y)   // a^b -> a^b // 1100^0101 = 1001 (9)
	fmt.Printf("%d&^%d = %d\n", x, y, x&^y) // a&^b -> a and (not b)  (đặt các vị trí ai thành công nếu bi=1)

	fmt.Println("===================================")

	// Assignment ops:
	// =, +=, -=, *=, /=, %=, &=, |=, ^=, <<=, >>=
	// := (init variable & assign value->  variable:=value)

	// Mics ops:
	// & : reference // (return the address of variable ): &a = address(a)
	// * : pointer // that can hold the address, not value of other variable
	// <- : receive // receive value from channel
	variable := 5
	fmt.Printf("address of variable: %d\n", &variable)
	// pointer
	var pointer *int = &variable
	// now:
	// pointer is hold the address of variable.
	// *pointer hold value that stored in the address that holded of pointer
	fmt.Printf("*pointer (*%d) : %d, variable = %d\n", pointer, *pointer, variable)
	// store new value
	*pointer = 100
	fmt.Printf("*pointer (*%d) : %d, variable = %d\n", pointer, *pointer, variable)

}
