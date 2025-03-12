package service

import (
	"context"
	"datagram-intelligence/config"
	"datagram-intelligence/pkg/constant"
	"datagram-intelligence/pkg/grpc"
	chat "datagram-intelligence/proto"
	"encoding/json"
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
	// Tạo context với timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Kết nối đến gRPC server
	conn, client, err := grpc.CreateGRPCConnection(svc.cfg.GRPC.LlamaCityPort)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// Tạo message gRPC
	messages := grpc.CreateGRPCMessages(request.Messages)

	// Gửi yêu cầu gRPC
	_, err = client.SendMessage(ctx, &chat.ChatRequest{
		Model:    request.Model,
		Messages: messages,
	})
	if err != nil {
		return nil, err
	}

	// Chờ kết quả từ kênh
	result := <-constant.ResultChan

	var resultObject map[string]interface{}
	if err := json.Unmarshal([]byte(result), &resultObject); err != nil {
		return nil, err
	}

	return resultObject, nil
}
