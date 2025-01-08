package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "My awesome go app")
}

func setupRoutes() {
	http.HandleFunc("/", homePage)

}

func main() {

	fmt.Println("Go app started on port 8000")
	setupRoutes()

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
