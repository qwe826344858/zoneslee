package websocket

import "fmt"

//连接池
type Pool struct{
	Regsiter chan *Client
	Unregsiter chan *Client
	Clients map[*Client] bool
	Broadcast chan Message
}

//当启动服务时会给连接池分配空间
func NewPool() *Pool{
	return &Pool{
		Regsiter:   make(chan *Client),
		Unregsiter: make(chan *Client),
		Clients:    make(map [*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start(){
	for{
		select {
		case client:= <-pool.Regsiter:
			pool.Clients[client] = true
			fmt.Println("连接池大小为：",len(pool.Clients))
			for client,_:=range pool.Clients{
				fmt.Println(client)
				client.Conn.WriteJSON(Message{Type:1,Body:"新用户正在加入..."})
			}
			break
		case client:= <-pool.Unregsiter:
			delete(pool.Clients,client)
			fmt.Println("连接池大小为：",len(pool.Clients))
			for client,_:=range pool.Clients{
				client.Conn.WriteJSON(Message{Type:1,Body:"用户退出..."})
			}
			break
		case message:= <-pool.Broadcast:
			fmt.Println("正在向全部用户发送信息")
			for client,_:=range pool.Clients{
				if err:=client.Conn.WriteJSON(message);
					err!=nil{
					fmt.Println(err)
					return
				}
			}
		}

	}

}