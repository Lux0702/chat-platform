package routes

import (
	authHttp "chat-platform-api/internal/auth/delivery/http"
	"chat-platform-api/pkg/middleware"

	"github.com/gin-gonic/gin"
)
func SetupRoutes( r *gin.Engine){

	authHandler := authHttp.NewAuthHandler()

	auth:= r.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	conversationHandler := authHttp.NewConversationHandler()

	conversation := r.Group("/conversations")
	conversation.Use(middleware.JWTMiddleware(),)
	conversation.POST(
		"/private",
		conversationHandler.CreatePrivateConversation,
	)

}