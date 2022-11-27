package main

import (
	"fmt"
	"net/http"
	"github.com/qwe826344858/zoneslee/pkg/websocket"
)

func serverWS(pool *websocket.Pool,w http.ResponseWriter,r *http.Request){
	fmt.Println("websocket endpoint reached")

	conn,err := websocket.Upgrade(w,r)
	if err!=nil{
		fmt.Fprint(w,"%+V\n",err)
	}

	client := &websocket.Client{
		Conn : conn,
		Pool : pool,
	}

	pool.Regsiter <- client
	client.Read()
}

func setupRoutes(){
	pool:=websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws",func(w http.ResponseWriter,r *http.Request){
		serverWS(pool,w,r)
	})
}

func main(){
	fmt.Println("多人聊天室")
	setupRoutes()
	http.ListenAndServe(":3000",nil)
}