package http

import (
	"github.com/ScMofeoluwa/GatherGo/internal/domain/user/service"
	"github.com/ScMofeoluwa/GatherGo/internal/transport/http/handlers"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Router      *gin.Engine
	AuthHandler *handlers.AuthHandler
}

func NewHandler(userService *service.UserService) *Handler {
	router := gin.Default()

	authHandler := handlers.NewAuthHandler(userService)
	authHandler.RegisterAuthRoutes(router.Group("/api"))

	return &Handler{
		Router:      router,
		AuthHandler: authHandler,
	}
}

func (h *Handler) Serve() error {
	return h.Router.Run(":8080")
}
