package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// 把 HTTP 连接升级为 WebSocket 连接
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// type Handler func(w http.ResponseWriter, r *http.Request)

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			return
		}

		log.Printf("Received message: %s", p)

		err = conn.WriteMessage(messageType, p)
		if err != nil {
			log.Println("Write error:", err)
			return
		}
	}
}
