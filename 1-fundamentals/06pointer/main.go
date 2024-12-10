package main

import "fmt"

func main() {
	fmt.Println("Welcome to the pointer")

	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr)

	myNumber := 23
	var ptr = &myNumber

	fmt.Println("Memory address of then pointer is", ptr)
	fmt.Println("Actual value of pointer is", *ptr)

	*ptr = *ptr + 2

	fmt.Println("New value is: ", myNumber)

}
