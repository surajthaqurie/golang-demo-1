package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Rance condition in golang")

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	var score = []int{0}

	wg.Add(4)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("One Routine")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	// wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {

		fmt.Println("Two Routine")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()

	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Three Routine")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()

	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Four Routine")
		mut.RLock()
		fmt.Println(score)
		mut.RUnlock()

		wg.Done()

	}(wg, mut)
	wg.Wait() // wait until all wait group finishes jobs

	fmt.Println("Score", score)

}
