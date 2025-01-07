package main

import (
	"fmt"
	"log"
	"net/http"
	"youtubeMonitor/websocket"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

// our new stats function which will expose any YouTube
// stats via a websocket connection
func stats(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	// to this websocket connection
	go websocket.Writer(ws)
}
func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/stats", stats)

	fmt.Println("Server started on :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	fmt.Println("Welcome to youtube monitor")
	setupRoutes()
}
