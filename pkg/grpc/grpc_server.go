package grpc

import (
	"context"
	"datagram-intelligence/config"
	"datagram-intelligence/pkg/constant"
	chat "datagram-intelligence/proto"
	"fmt"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

// server defines the server struct for the gRPC server
type server struct {
	chat.UnimplementedChatServiceServer
}

// ReceiveOllamaResponse processes the gRPC request from llama_citizen
func (s *server) ReceiveOllamaResponse(ctx context.Context,
	req *chat.OllamaResponseRequest) (*chat.OllamaResponse, error) {
	// Send the result to the channel
	constant.ResultChan <- req.Response

	return &chat.OllamaResponse{Response: "received successfully"}, nil
}

// StartGRPCServer initializes the gRPC server
func StartGRPCServer(config *config.Config) {
	serverPort, err := strconv.Atoi(config.GRPC.ServerPort)
	if err != nil {
		log.Fatalf("Error converting port: %v", err)
	}

	grpcServer := grpc.NewServer()
	chat.RegisterChatServiceServer(grpcServer, &server{})

	// Listen on the port from the configuration
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", serverPort))
	if err != nil {
		log.Fatalf("Could not listen on port %d: %v", serverPort, err)
	}
	fmt.Printf("Server is running on port %d...\n", serverPort)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Could not serve gRPC: %v", err)
	}
}
