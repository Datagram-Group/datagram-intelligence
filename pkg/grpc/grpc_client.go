package grpc

import (
	chat "datagram-intelligence/proto"
	"fmt"
	"log"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// CreateGRPCConnection tạo kết nối gRPC
func CreateGRPCConnection(port string) (*grpc.ClientConn, chat.ChatServiceClient, error) {
	llamaCityPort, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("Lỗi khi chuyển đổi cổng: %v", err)
	}

	conn, err := grpc.NewClient(fmt.Sprintf("localhost:%d", llamaCityPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, fmt.Errorf("không thể kết nối đến gRPC server: %v", err)
	}
	client := chat.NewChatServiceClient(conn)
	return conn, client, nil
}

// CreateGRPCMessages tạo các message gRPC từ dữ liệu nhận được
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
