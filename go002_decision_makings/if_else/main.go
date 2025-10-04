package main

import "fmt"

// if-else if-else syntax:
/*
if condition_1 {
	// execution_1
} else if condition_2 {
	// execution_2

} else if ...  {
	// execution ...
} else {
	// default_execution
}
*/

// if-else block have:
// - if statment (required, only one)
// - else if statment (optional, multiple is ok)
// - else (optional, only one)

func main() {
	// get number of day in a month with if-else
	var month string
	fmt.Print("enter a month (1-12): ")
	fmt.Scan(&month)

	if month == "4" || month == "6" || month == "9" || month == "11" {
		fmt.Printf("The %sth month in the year has 30 days", month)
	} else if month == "2" {
		fmt.Printf("The %sth month in the year has 28 or 29 days", month)
	} else if month == "1" || month == "3" || month == "5" ||
		month == "7" || month == "8" || month == "10" || month == "12" {
		fmt.Printf("The %sth month in the year has 31 days", month)
	} else {
		fmt.Println("Invalid month!")
	}

}
