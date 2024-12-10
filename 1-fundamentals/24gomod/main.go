package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in golang")

	greeter()

	//Initialize the route
	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("GET")

	//Server listing with port and router
	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hey there mod user")
}

func serverHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1>Welcome to golang series on YT</h1>"))
}
