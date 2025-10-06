package main

import "fmt"

type user struct {
	name string
	age  int
}

func (u user) log() {
	fmt.Printf("user {\n\tname: %s,\n\tage: %d\n}\n", u.name, u.age)
}

// passing pointer to function
func swap(a *int, b *int) {
	*a += *b
	*b = *a - *b
	*a -= *b
}

// return pointer from function
// func funcname(args types) *T {

// 	return &variable // variable of type T
// }

func main() {
	// Create pointer
	// 1. create a null pointer
	// var intPtr *int // nil

	// 2. point to existing variable
	var intVar int = 5
	intPtr := &intVar

	// 3. pointer to pointer
	ptrPtr := &intPtr

	// note: *T point to &varibale of type T
	// ptrPtr -> intPtr -> intVar

	fmt.Println("address of intVar:", &intVar)
	fmt.Println("intPtr:", intPtr)   // pointer holds address of intVar
	fmt.Println("*intPtr:", *intPtr) // *pointer holds value
	fmt.Println("===")
	fmt.Println("address of intPtr:", &intPtr)
	fmt.Println("ptrPtr:", ptrPtr)     // address of intPtr
	fmt.Println("*ptrPtr:", *ptrPtr)   // address of intVar
	fmt.Println("**ptrPtr:", **ptrPtr) // intVar

	fmt.Println("===")
	// array of pointer
	var arrPtr [5]*int
	// access elem
	arrPtr[0] = &intVar
	fmt.Println(*arrPtr[0])
	fmt.Println("cap:", cap(arrPtr)) // maximum element storable in array
	fmt.Println("len:", len(arrPtr)) // số phần tử hiện truy cập được
	// với array, cap & len luôn bằng nhau

	// với slice có thể khác
	slc := make([]int, 3, 5)
	fmt.Println("cap(slc):", cap(slc))
	fmt.Println("len(slc):", len(slc))

	fmt.Println()
	fmt.Println("============================")
	fmt.Println()

	// passing pointer to func
	a, b := 3, 5
	fmt.Printf("before swap: a = %d, b = %d\n", a, b)
	swap(&a, &b)
	fmt.Printf("after swap: a = %d, b = %d\n", a, b)

	fmt.Println()
	fmt.Println("============================")
	fmt.Println()
	// point to struct
	// 1. using &
	u1 := user{name: "vduczz", age: 21}
	userPtr1 := &u1
	// truy cập thành phần field, method qua pointer trực tiếp
	// không cần qua *pointer
	userPtr1.log() // u1.log() // (*userPtr1).log()
	// 2. using new()
	userPtr2 := new(user)
	userPtr2.name = "vduczz"
	userPtr2.age = 21
	userPtr2.log()

	fmt.Println()
	fmt.Println("============================")
	fmt.Println()
}
