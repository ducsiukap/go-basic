package main

import "fmt"

func main() {
	/* To declare a variable:
		var var_name dtype = expression

	- either [dtype] or [= expression] can be omitted, but not both.
	- if [= expression] is omitted, the default value of dtype is assigned to variable
	*/
	var loveSport = false // int

	// we can declare multiple variables in one line
	var name, age, smile = "vduczz", 21, '😀'

	//  or using short variable declaration
	// varname:=value
	myFavoriteColor, green := "green", '🟢'

	fmt.Println("Hi, guys!")
	fmt.Printf("I'm %s, %dyo.\n", name, age)
	fmt.Printf("My favourite color is: %s%c\n", myFavoriteColor, green)
	fmt.Printf("Am i love sport : %t%c", loveSport, smile)
}
