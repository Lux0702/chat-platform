package http

import (
	"chat-platform-api/internal/auth/dto"
	"chat-platform-api/internal/auth/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ConversationHandler struct {
	service *service.ConversationService
}

func NewConversationHandler() *ConversationHandler {
	return &ConversationHandler{
		service: service.NewConversationService(),
	}
}

func (h *ConversationHandler) CreatePrivateConversation(c *gin.Context) {

	var body dto.CreateConversationRequest

	err := c.ShouldBindJSON(&body)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		},
	)
		return
	}
	currentUserID := c.GetString("userId")
	err = h.service.CreatePrivateConversation(
		currentUserID,
		body.UserID,
	)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		},)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "conversation created",
	},
)
	

}