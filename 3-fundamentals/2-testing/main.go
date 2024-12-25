package main

import "fmt"

func main() {
	fmt.Println("Welcome to the testing in the golang")

	result := Calculate(2)
	fmt.Println("Result", result)
}

func Calculate(x int) (result int) {
	result = x + 2

	return result
}
