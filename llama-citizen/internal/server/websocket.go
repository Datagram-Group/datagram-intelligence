package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Datagram-Group/datagram-intelligence/llama-citizen/config"
	"github.com/Datagram-Group/datagram-intelligence/llama-citizen/internal/ollama"
	"github.com/Datagram-Group/datagram-intelligence/llama-citizen/pkg/structure"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}
	defer conn.Close()

	for {
		var msg structure.Message
		if err := conn.ReadJSON(&msg); err != nil {
			log.Println("Error reading message:", err)
			break
		}

		// Gửi yêu cầu đến Ollama server và nhận kết quả
		response, err := ollama.GetResponse(msg)
		if err != nil {
			log.Println("Error getting response from Ollama:", err)
			break
		}

		// Gửi phản hồi cho client
		if err := conn.WriteJSON(response); err != nil {
			log.Println("Error sending response:", err)
			break
		}
	}
}

func StartServer(config config.Config) {
	http.HandleFunc("/ws", HandleWebSocket)
	log.Printf("WebSocket server started on :%d", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil))
}
