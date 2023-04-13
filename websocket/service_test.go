package websocket

import (
	"log"
	"net/http"
	"testing"
)

func TestWebSocket(t *testing.T) {
	http.HandleFunc("/websocket", websocketHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe error:", err)
	}
}
