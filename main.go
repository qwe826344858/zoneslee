package main

import (
	"fmt"
	"github.com/qwe826344858/zoneslee/pkg/websocket"
	"net/http"
)

func serverWS(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("websocket endpoint reached")

	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprint(w, "%+v\n", err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Regsiter <- client
	client.Read()
}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serverWS(pool, w, r)
	})
}

func main() {
	fmt.Println("多人聊天室开启")
	setupRoutes()
	http.ListenAndServe(":9000", nil)
}
