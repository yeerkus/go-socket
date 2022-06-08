package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Message struct {
	Greeting string `json:"greeting"`
}

var (
	wsUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	wsConn *websocket.Conn
)

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	wsUpgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	wsConn, err := wsUpgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println("Coudn't upgrade Socket connection")
		fmt.Println(err)
		return
	}

	for {
		var msg Message

		err := wsConn.ReadJSON(&msg)

		if err != nil {
			fmt.Println("Couldn't Read message")
			fmt.Print(err)
			break
		}

		fmt.Printf("Message: %s ", msg.Greeting)
	}

	defer wsConn.Close()
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/socket", WsEndpoint)

	log.Fatal(http.ListenAndServe(":9090", r))
}
