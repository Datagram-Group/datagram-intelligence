package main

import (
	"log"

	"github.com/Datagram-Group/datagram-intelligence/llama-citizen/config"
	"github.com/Datagram-Group/datagram-intelligence/llama-citizen/internal/ollama"
	"github.com/Datagram-Group/datagram-intelligence/llama-citizen/internal/server"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	if err := ollama.StartOllamaServer(cfg.OllamaDir); err != nil {
		log.Fatalf("Error starting Ollama server: %v", err)
	}

	server.StartServer(cfg)
}
