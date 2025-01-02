package main

import (
	"fmt"
	"time"
)

func SensValue(c chan string) {
	fmt.Println("Executing the go routine")

	time.Sleep(1 * time.Second)
	c <- "Hello world" //assign the value

	fmt.Println("Finished executing Goroutine")

}

func main() {
	fmt.Println("Go Channels")

	// values := make(chan string) //un-buffer channel
	values := make(chan string, 2) //buffer channel- capacity
	defer close(values)

	go SensValue(values)
	go SensValue(values)

	value := <-values //Get the value
	fmt.Println("Channel value", value)

	time.Sleep(1 * time.Second)

}
