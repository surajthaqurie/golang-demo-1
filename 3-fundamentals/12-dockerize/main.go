package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Dockerize golang")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "Hello world")

	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
