package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"youtubeMonitor/youtube"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  10240,
	WriteBufferSize: 1024,
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	// this line allows other origin hosts to connect to our
	// websocket server
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// creates our websocket connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return ws, err
	}
	// returns our new websocket connection
	return ws, nil

}

func Writer(conn *websocket.Conn) {
	// we want to kick off a for loop that runs for the
	// duration of our websockets connection
	for {
		// we create a new ticker that ticks every 5 seconds
		ticker := time.NewTicker(5 * time.Second)

		// every time our ticker ticks
		for t := range ticker.C {
			// print out that we are updating the stats
			fmt.Printf("Updating Stats: %+v\n", t)
			// and retrieve the subscribers
			items, err := youtube.GetSubscriber()
			if err != nil {
				fmt.Println(err)
			}
			// next we marshal our response into a JSON string
			jsonString, err := json.Marshal(items)
			if err != nil {
				fmt.Println(err)
			}
			// and finally we write this JSON string to our WebSocket
			// connection and record any errors if there has been any
			if err := conn.WriteMessage(websocket.TextMessage, []byte(jsonString)); err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}
