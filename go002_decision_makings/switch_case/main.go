package main

import "fmt"

// Syntax
/*
switch expression {
	case case_1a, case1b, case1c:
		// execution 1
	case case_2:
		// execution 2
	...
	case case_na, case_nb:
		// execution n
	default:
		// default execution
}
*/

// Mặc định, sau khi thực hiện xong matched case, go sẽ thoát khỏi vòng lặp (không cần break)

// Muốn sau khi thực hiện xong case đó mà vẫn tiếp tục xét các case bên dưới, thêm keyword 'fallthough'
/*
	case case_x:
		// execution
		fallthough
*/

func main() {

	fmt.Print("enter a month (1-12): ")
	var month string
	fmt.Scan(&month)

	var month_name string
	// get month name
	switch month {
	case "1":
		month_name = "January"
	case "2":
		month_name = "February"
	case "3":
		month_name = "March"
	case "4":
		month_name = "April"
	case "5":
		month_name = "May"
	case "6":
		month_name = "June"
	case "7":
		month_name = "July"
	case "8":
		month_name = "August"
	case "9":
		month_name = "September"
	case "10":
		month_name = "October"
	case "11":
		month_name = "November"
	case "12":
		month_name = "December"
	default:
		month_name = "Not a month"
	}

	// get number of day in month with switch case
	switch month {
	case "1", "3", "5", "7", "8", "10", "12":
		fmt.Printf("The %s has 31 days!", month_name)
	case "2":
		fmt.Printf("The %s has 28 or 29 days!", month_name)
	case "4", "6", "9", "11":
		fmt.Printf("The %s has 30 days!", month_name)
	default:
		fmt.Print(month_name)
	}

	// =====================================
	// Switch-case cho channel
	ch := make(chan int)

	// gửi vào channel
	select {
	case ch <- 3: // gui thanh cong
		fmt.Println("Send 3 to channel ch")

	default: // gui that bai
		fmt.Println("Can not send to channel ch")
	}

	// nhận từ channel
	select {
	// nếu có người gửi vào channel ch
	case x := <-ch:
		// or:
		// case <- ch
		fmt.Printf("Read %d from channel", x)

	// nếu channel đang không có gì
	default:
		fmt.Println("There nothing in the channel")
	}

	// note: nếu select không có default, nó sẽ chờ cho tới khi
	// có người gửi vào channel (trường hợp select nhận)
	// gửi thành công (trường hợp select gui)
	// => deadlock
	ch2 := make(chan string)

	// ex: no one send anything to channel ch
	select {
	// ch has data
	case x := <-ch:
		fmt.Printf("Read %d from channel ch", x)

	// ch2 has data
	case x := <-ch2:
		fmt.Printf("Read %s from channel ch", x)

	}
	// dòng select trên chờ tới khi ch / ch2 có ai đó gửi vào
	// error:
	// 	fatal error: all goroutines are asleep - deadlock!

	fmt.Println("This line will never be reached!")

}
