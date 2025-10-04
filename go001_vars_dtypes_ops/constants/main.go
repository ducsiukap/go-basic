package main

import "fmt"

func main() {
	/* Declare a constant : using const keyword
		const varname = value

	or declare multiple constants
		const (
			var1 = value1
			var2 = value2
			var3 = valu3
		)

	//  CANNOT be declared with syntax: varname:=value
	*/

	// Numberic constants: Integer, Floating point, Complex
	// Integer constant:
	// 		- prefix specifies the base: 0x/0X for hexadecimal, 0 for octal & nothing for decimal
	// 			-> 88 -> decimal, 088 -> octal, 0x88 -> hexa
	// 		- suffix : u/U -> unsigned, l/L -> long, nothing -> integer
	// Floating Point:
	// 		- has integer part, decimal point, fractional part & an exponent part
	// Complext

	// String constant
	// Boolean constant

	fmt.Print("enter your name: ")
	var guy string
	fmt.Scan(&guy)

	const (
		smile       = '😀'
		waving_hand = '👋'
	)

	hello := fmt.Sprintf("hi %s%c", guy, waving_hand)
	const introduction = "im vduczz. nice to meet you "

	fmt.Printf("%s\n%s%c", hello, introduction, smile)
}
