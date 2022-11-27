package websocket
import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrade{
	ReadBUfferSize : 1024,
	WriteBUfferSize : 1024,
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