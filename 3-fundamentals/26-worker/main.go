package main

import (
	"fmt"
	"log"
	"net/http"
)

// Site represents a website with a URL to be crawled.
type Site struct {
	URL string
}

// Result represents the result of crawling a site, including its URL and HTTP status code.
type Result struct {
	URL    string
	Status int
}

// crawl is a worker function that takes a worker ID, reads sites from the jobs channel,
// performs HTTP GET requests, and sends the results to the results channel.
func crawl(wId int, jobs <-chan Site, results chan<- Result) {
	for site := range jobs { // Loop until jobs channel is closed.
		log.Printf("Worker Id: %d\n", wId) // Log which worker is processing the site.

		resp, err := http.Get(site.URL) // Perform an HTTP GET request.
		if err != nil {                 // Handle errors during the request.
			log.Println(err.Error())
			continue
		}

		// Send the result to the results channel.
		results <- Result{Status: resp.StatusCode, URL: site.URL}
	}
}

func main() {
	fmt.Println("Worker pools in GO") // Start message.

	// Create buffered channels for jobs and results.
	jobs := make(chan Site, 3)
	results := make(chan Result, 3)

	// Start 3 worker goroutines.
	for w := 1; w <= 3; w++ {
		go crawl(w, jobs, results)
	}

	// List of URLs to crawl.
	urls := []string{
		"https://google.com",
		"https://facebook.com",
		"https://youtube.com",
		"https://go.dev/ref/mod",
	}

	// Send jobs (URLs) to the jobs channel.
	for _, url := range urls {
		jobs <- Site{URL: url}
	}

	// Close the jobs channel to signal workers that there are no more jobs.
	close(jobs)

	// Collect results from the workers. The number of results matches the number of URLs.
	for a := 1; a <= len(urls); a++ {
		result := <-results // Receive results from the results channel.
		log.Println(result) // Log the result (URL and status code).
	}
}
