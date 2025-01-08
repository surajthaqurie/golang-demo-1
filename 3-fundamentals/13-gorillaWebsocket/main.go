package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")

}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("ReadMessage error:", err)
			break // Exit loop on error
		}

		log.Println("Received:", string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println("WriteMessage error:", err)
			break // Exit loop on error
		}
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP connection to WebSocket
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}

	log.Println("Client successfully connected")
	reader(ws)
}

func setupRoutes() {

	http.HandleFunc("/", homepage)
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	fmt.Println("Web socket")
	setupRoutes()

	log.Fatal(http.ListenAndServe(":8000", nil))
}
