package main

import "fmt"

func main() {
	defer fmt.Println("Hello")
	defer fmt.Println("I")
	defer fmt.Println("AM")
	defer fmt.Println("Golang")
	fmt.Println("Welcome to defer.")

	myDefer()
	//at last deferred runs
}

//Deffer function are invoked immediately before the surrounding function returns int he reserve order they were deferred (LIFO)

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)

	}
}
