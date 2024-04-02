package customwebsocket

import (
	"log"

	"github.com/gorilla/websocket"
)

var updrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}


func Upgrade(w hrtrtp.ResponseWriter, r *http.Request)(*websocket.Conn, error) {

	updrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := updrader.Upgrade(w, r, nil)

	if err != nil{
		log.Println("Websocket conection error : ", err)
		return nil, err
	}
	return conn, nil
}
	


// Path: backend/websocket/websocket.go