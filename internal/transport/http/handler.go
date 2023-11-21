package http

import (
	"github.com/ScMofeoluwa/GatherGo/internal/domain/user/service"
	"github.com/ScMofeoluwa/GatherGo/internal/transport/http/handlers/auth"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Router      *gin.Engine
	AuthHandler *auth.AuthHandler
}

// make this take in a struct of services instead
// in case the function grows larger
func NewHandler(userService service.UserServiceInterface) *Handler {
	router := gin.Default()

	authHandler := auth.NewAuthHandler(userService)
	authHandler.RegisterAuthRoutes(router.Group("/api"))

	return &Handler{
		Router:      router,
		AuthHandler: authHandler,
	}
}

func (h *Handler) Serve() error {
	return h.Router.Run(":8080")
}
