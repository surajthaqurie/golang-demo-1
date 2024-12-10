package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Welcome to channels")

	myCh := make(chan int, 2)

	// fmt.Println(<-myCh)
	// myCh <- 5 // push value

	wg := &sync.WaitGroup{}

	wg.Add(2)
	// Receive ONLY
	go func(ch <-chan int, wg *sync.WaitGroup) {

		// close(myCh)

		val, isChannelOpen := <-myCh

		fmt.Println(isChannelOpen)
		fmt.Println(val)

		// fmt.Println(<-myCh)
		// fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)

	//Send ONLY
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 0
		close(myCh) // close the channel
		// myCh <- 5
		// myCh <- 6

		wg.Done()
	}(myCh, wg)

	wg.Wait()

}
