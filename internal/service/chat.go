package service

import (
	"context"
	"datagram-intelligence/config"
	"datagram-intelligence/pkg/constant"
	"datagram-intelligence/pkg/grpc"
	chat "datagram-intelligence/proto"
	"encoding/json"
	"fmt"
	"time"
)

type ChatQuestionRequest struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
}

type ChatService interface {
	ChatQuestion(request ChatQuestionRequest) (map[string]interface{}, error)
}

type chatService struct {
	cfg *config.Config
}

func NewChatService(cfg *config.Config) ChatService {
	return &chatService{cfg}
}

func (svc *chatService) ChatQuestion(request ChatQuestionRequest) (map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Connect to the gRPC server
	conn, client, err := grpc.CreateGRPCConnection(svc.cfg.GRPC.LlamaCityPort)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// Create gRPC message
	messages := grpc.CreateGRPCMessages(request.Messages)

	// Send gRPC request
	response, err := client.SendMessage(ctx, &chat.ChatRequest{
		Model:    request.Model,
		Messages: messages,
	})
	if err != nil {
		return nil, err
	}
	fmt.Printf("Received ACK from city: %s\n", response.AckMessage)

	// Wait for the result from the channel
	result := <-constant.ResultChan

	var resultObject map[string]interface{}
	if err := json.Unmarshal([]byte(result), &resultObject); err != nil {
		return nil, err
	}

	return resultObject, nil
}
