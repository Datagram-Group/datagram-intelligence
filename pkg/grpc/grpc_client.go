package grpc

import (
	chat "datagram-intelligence/proto"
	"fmt"
	"log"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// CreateGRPCConnection creates a gRPC connection
func CreateGRPCConnection(port string) (*grpc.ClientConn, chat.ChatServiceClient, error) {
	llamaCityPort, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("Error converting port: %v", err)
	}

	conn, err := grpc.NewClient(fmt.Sprintf("llama-city:%d", llamaCityPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to gRPC server: %v", err)
	}
	client := chat.NewChatServiceClient(conn)
	return conn, client, nil
}

// CreateGRPCMessages creates gRPC messages from received data
func CreateGRPCMessages(messages []struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}) []*chat.Message {
	var grpcMessages []*chat.Message
	for _, msg := range messages {
		grpcMessages = append(grpcMessages, &chat.Message{Role: msg.Role, Content: msg.Content})
	}
	return grpcMessages
}
