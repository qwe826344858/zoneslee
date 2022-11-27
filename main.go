package main

import (
	"fmt"
	"net/http"
	//"github.com/akhil/golang-chat"
	"backend/pkg/websocket"
)

func serverWS(pool *websocket.Pool,w http.ResponseWriter,r *http.Request){
	fmt.Println("websocket endpoint reached")

	conn,err := websocket.Updrade(w,r)
	if err!=nil{
		fmt.Fprint(w,"%+V\n",err)
	}

	client := &websocket.Clinent{
		Conn : conn,
		Pool : pool,
	}

	pool.Register <- client
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
	http.ListenAndServe(":9000",nil)
}