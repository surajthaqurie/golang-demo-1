package main

import "fmt"

func main() {
	fmt.Println("Welcome in the function in golang")

	greeter()
	greeterTwo()

	result := adder(3, 5)
	fmt.Println("Result is:", result)

	proResult, myMessage := proAdder(2, 3, 4, 5)
	fmt.Println("pro Result is:", proResult)
	fmt.Println("Pro message is:", myMessage)

}

func greeter() {
	fmt.Println("Hello from golang")
}

func greeterTwo() {
	fmt.Println("Welcome the function")

}

func adder(valueOne int, valueTwo int) int {
	return valueOne + valueTwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, value := range values {
		total += value
	}

	return total, "Hi pro result function"

}

// Define function inside function Not allowed
