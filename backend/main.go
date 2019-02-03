package main

import (
	"fmt"
	"net/http"

	"github.com/yuchou87/golang-react-chat/pkg/websocket"
)

func serverWs(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}
	go websocket.Writer(ws)

	websocket.Reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/ws", serverWs)
}

func main() {
	fmt.Println("Distributed Chat App v0.01")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
