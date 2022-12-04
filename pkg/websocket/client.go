package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
	mu   sync.Mutex
}

type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

//持续监听前端返回的信息
func (c *Client) Read() {
	defer func() {
		c.Pool.Unregsiter <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		message := Message{Type: messageType, Body: string(p)}
		c.Pool.Broadcast <- message
		fmt.Println("message received: %+v\n", message)
	}
}
