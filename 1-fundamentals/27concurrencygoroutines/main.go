package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup //pointer
var mut sync.Mutex    //pointer

var signals = []string{"test"}

func main() {
	// go greeter("Hello")
	// greeter("World")

	websitelist := []string{
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)

		wg.Add(1)
	}

	wg.Wait() // always end into the main

	fmt.Println("The signal is", signals)

}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()

	result, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS is endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", result.Status, endpoint)
	}
}
