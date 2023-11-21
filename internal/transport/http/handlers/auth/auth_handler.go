package auth

import (
	"errors"
	"net/http"

	"github.com/ScMofeoluwa/GatherGo/internal/domain/user/entity"
	"github.com/ScMofeoluwa/GatherGo/internal/domain/user/repository"
	"github.com/ScMofeoluwa/GatherGo/internal/domain/user/service"
	"github.com/ScMofeoluwa/GatherGo/internal/transport/http/handlers"
	"github.com/gin-gonic/gin"
)

type SignUpRequest struct {
	Password string `json:"password" binding:"required,min=6,max=255"`
	Email    string `json:"email" binding:"required,email"`
}

type AuthHandler struct {
	service service.UserServiceInterface
}

func NewAuthHandler(service service.UserServiceInterface) *AuthHandler {
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
	var req SignUpRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, handlers.NewApiError(err.Error()))
		return
	}

	user := createUserFromReq(req)
	if err := a.service.SignUp(ctx, &user); err != nil {
		if errors.Is(err, service.ErrEmailAlreadyExists) {
			ctx.JSON(http.StatusForbidden, handlers.NewApiError(err.Error()))
			return
		}
		ctx.JSON(http.StatusInternalServerError, handlers.NewApiError(err.Error()))
		return
	}
	ctx.JSON(http.StatusCreated, handlers.NewApiResponse(nil, "user created successfully"))
}

func (a *AuthHandler) SignIn(ctx *gin.Context) {
	var req SignUpRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, handlers.NewApiError(err.Error()))
	}

	user := createUserFromReq(req)
	data, err := a.service.SignIn(ctx, &user)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, handlers.NewApiError(err.Error()))
			return
		}
		if errors.Is(err, service.ErrInvalidDetails) {
			ctx.JSON(http.StatusUnauthorized, handlers.NewApiError(err.Error()))
			return
		}
		ctx.JSON(http.StatusInternalServerError, handlers.NewApiError(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, handlers.NewApiResponse(data, "login successful"))
}

func createUserFromReq(r SignUpRequest) entity.CreateUser {
	return entity.CreateUser{
		Email:    r.Email,
		Password: r.Password,
	}
}
