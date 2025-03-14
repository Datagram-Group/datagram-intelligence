package main

import (
	"datagram-intelligence/config"
	"datagram-intelligence/internal/handler"
	"datagram-intelligence/internal/router"
	"datagram-intelligence/internal/service"
	"datagram-intelligence/pkg/grpc"
	"datagram-intelligence/pkg/validator"
	"flag"
	"log"
)

func main() {
	configPath := flag.String("config", "../config.yaml", "Path to the config file")

	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	// Initialize the data validator
	validator.InitValidator()

	chatService := service.NewChatService(cfg)

	chatHandler := handler.NewChatHandler(chatService)

	go grpc.StartGRPCServer(cfg)

	// Setup routes
	router := router.SetupRouter(chatHandler, cfg)
	router.Run(cfg.HTTPServer.Port)
}
