package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal("error establishing new socketio server")
	}

	server.On("connection", func(so socketio.Socket) {
		log.Println("On connection established")

		so.Join("mychat") //join the room with the above socketio connection

		so.On("chat message", func(msg string) {
			log.Println("messge received: " + msg)
			so.BroadcastTo("mychat", "chat message", msg) //broadcasting the message to the room so all the clients connected can get it
		})
	})

	fmt.Println("Server running on localhost: 5001")

	//create a folder with name static and within that index.html
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.Handle("/socket.io/", server)

	log.Fatal(http.ListenAndServe(":5001", nil))

}
