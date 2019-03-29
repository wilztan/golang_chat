package main

import (
	"fmt"
	"net/http"

	"github.com/wilztan/golang_chat/pkg/websocket"
)

func serverWs(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Println(err)
		return
	}
	go websocket.Writer(ws)
	websocket.Reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})
	http.HandleFunc("/ws", serverWs)
}

func main() {
	setupRoutes()
	http.ListenAndServe(":8082", nil)
}
