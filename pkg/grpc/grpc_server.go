package grpc

import (
	"context"
	"datagram-intelligence/config"
	"datagram-intelligence/pkg/constant"
	chat "datagram-intelligence/proto"
	"fmt"
	"log"
	"net"
	"strconv" // Thêm package để chuyển đổi kiểu

	"google.golang.org/grpc"
)

// Server định nghĩa server struct cho gRPC server
type server struct {
	chat.UnimplementedChatServiceServer
}

// ReceiveOllamaResponse xử lý yêu cầu gRPC từ llama_citizen
func (s *server) ReceiveOllamaResponse(ctx context.Context,
	req *chat.OllamaResponseRequest) (*chat.OllamaResponse, error) {
	// Gửi kết quả cho kênh
	constant.ResultChan <- req.Response

	// Trả về phản hồi thành công
	return &chat.OllamaResponse{Response: "received successfully"}, nil
}

// StartGRPCServer khởi tạo gRPC server
func StartGRPCServer(config *config.Config) {
	// Chuyển đổi ServerPort từ string sang int
	serverPort, err := strconv.Atoi(config.GRPC.ServerPort)
	if err != nil {
		log.Fatalf("Lỗi khi chuyển đổi cổng: %v", err)
	}

	grpcServer := grpc.NewServer()
	chat.RegisterChatServiceServer(grpcServer, &server{})

	// Lắng nghe cổng từ cấu hình
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", serverPort))
	if err != nil {
		log.Fatalf("Không thể nghe trên cổng %d: %v", serverPort, err)
	}
	fmt.Printf("Server đang chạy trên cổng %d...\n", serverPort)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Không thể phục vụ gRPC: %v", err)
	}
}
