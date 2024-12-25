package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func allArticles(w http.ResponseWriter, r *http.Request) {

	articles := Article{Title: "articleOne", Desc: "This is first article", Content: "Hello world"}

	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func testPostArticles(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint Hit: TEST POST Articles Endpoint")
}

func main() {
	handleRequest()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint hit")
}

func handleRequest() {

	//@Normal router
	// http.HandleFunc("/", homePage)
	// http.HandleFunc("/articles", allArticles)
	// log.Fatal(http.ListenAndServe(":8081", nil))

	//@MUX router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage).Methods("GET")
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
