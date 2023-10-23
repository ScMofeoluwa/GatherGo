package handlers

import (
	"net/http"

	"github.com/ScMofeoluwa/GatherGo/internal/domain/user/entity"
	"github.com/ScMofeoluwa/GatherGo/internal/domain/user/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *service.UserService
}

func NewAuthHandler(service *service.UserService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (a *AuthHandler) RegisterAuthRoutes(router *gin.RouterGroup) {
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/register", a.SignUp)
		authRoutes.POST("/login", a.SignIn)
	}
}

func (a *AuthHandler) SignUp(ctx *gin.Context) {
	var req entity.CreateUser
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, NewApiError(err.Error()))
		return
	}
	if err := a.service.SignUp(ctx, &req); err != nil {
		ctx.JSON(http.StatusBadRequest, NewApiError(err.Error()))
		return
	}
	ctx.JSON(http.StatusCreated, NewApiResponse(nil, "user created successfully"))
}

func (a *AuthHandler) SignIn(ctx *gin.Context) {
	var req entity.CreateUser
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, NewApiError(err.Error()))
	}
	data, err := a.service.SignIn(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewApiError(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, NewApiResponse(data, "login successful"))
}
