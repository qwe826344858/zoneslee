package websocket
import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

// 根据gorilla使用文档
// 我们需要定义一个 Upgrader
// 它需要定义 ReadBufferSize 和 WriteBufferSize
var upgrader = websocket.Upgrader{
	ReadBufferSize : 1024,
	WriteBufferSize : 1024,
}

func Upgrade(w http.ResponseWriter,r *http.Request)(*websocket.Conn,error){
	upgrader.CheckOrigin = func(r *http.Request) bool {return true}
	conn,err :=upgrader.Upgrade(w,r,nil)
	if err != nil{
		log.Println(err)
		return nil,err
	}
	return conn,nil
}