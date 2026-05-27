package http

import (
	"chat-platform-api/internal/auth/dto"
	"chat-platform-api/internal/auth/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler() *AuthHandler{
	return &AuthHandler{
		authService: service.NewAuthService(),
	}
}

func (h *AuthHandler) Register( c *gin.Context) {

	var body dto.RegisterRequest

	err:= c.ShouldBindJSON(&body);

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = h.authService.Register(
		body.Name,
		body.Email,
		body.Password,
	)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message":"register success",
	})
}

func (h *AuthHandler) Login(c *gin.Context){

	var body dto.LoginRequest

	err:= c.ShouldBindJSON(&body);
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	token, err := h.authService.Login(
		body.Email,
		body.Password,
	)
	if err !=nil{
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}