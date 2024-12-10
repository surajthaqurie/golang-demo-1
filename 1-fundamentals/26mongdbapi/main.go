package main

import (
	"fmt"
	"log"
	"mongodbapi/router"
	"net/http"
)

func main() {
	fmt.Println("Welcome to db integration in golang")

	r := router.Router()

	fmt.Printf("Serving is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Printf("Listing at port 4000...")

}
