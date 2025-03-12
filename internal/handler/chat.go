package handler

import (
	"datagram-intelligence/internal/service"
	"datagram-intelligence/pkg/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ChatHandler interface {
	ChatQuestion(c *gin.Context)
}

type chatHandler struct {
	chatService service.ChatService
}

func NewChatHandler(chatService service.ChatService) ChatHandler {
	return &chatHandler{chatService}
}

func (h *chatHandler) ChatQuestion(c *gin.Context) {
	var request service.ChatQuestionRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if err := validator.ValidateStruct(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	chatQuestionResponse, err := h.chatService.ChatQuestion(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": chatQuestionResponse})
}
