package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

/*************** Example 1 ***************/

// func myFun(wg *sync.WaitGroup) {
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("Finished executing Goroutine")
// 	wg.Done() // done the wait group
// }
// func main() {
// 	fmt.Println("Go Wait groups")

// 	// Initialized wait groups
// 	var wg sync.WaitGroup
// 	wg.Add(1) // add the wait groups how many time to waits (time --)
// 	go myFun(&wg)
// 	wg.Wait() // blocks until 0
// 	fmt.Println("Finished Executing my go program")
// }

/*********** Example 2 ***********/
var urls = []string{
	"https://google.com",
	"https://tutorialedge.net",
	"https://twitter.com",
}

func fetchStatus(w http.ResponseWriter, r *http.Request) {

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)

		go func(url string) {
			resp, err := http.Get(url)

			if err != nil {
				fmt.Fprintf(w, "%+v\n", err)
			}

			fmt.Fprintf(w, "%+v\n", resp.Status)
			wg.Done()
		}(url)
	}

	wg.Wait()
}

func main() {
	fmt.Println("Go Wait groups")

	http.HandleFunc("/", fetchStatus)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
