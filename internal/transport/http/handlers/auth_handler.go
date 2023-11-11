package handlers

import (
	"net/http"

	"github.com/ScMofeoluwa/GatherGo/internal/domain/user/entity"
	"github.com/ScMofeoluwa/GatherGo/internal/domain/user/service"
	"github.com/gin-gonic/gin"
)

type SignUpRequest struct {
	Password string `json:"password" binding:"required,min=6,max=255"`
	Email    string `json:"email" binding:"required,email"`
}

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
	var req SignUpRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, NewApiError(err.Error()))
		return
	}

	user := createUserFromReq(req)
	if err := a.service.SignUp(ctx, &user); err != nil {
		ctx.JSON(http.StatusBadRequest, NewApiError(err.Error()))
		return
	}
	ctx.JSON(http.StatusCreated, NewApiResponse(nil, "user created successfully"))
}

func (a *AuthHandler) SignIn(ctx *gin.Context) {
	var req SignUpRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, NewApiError(err.Error()))
	}

	user := createUserFromReq(req)
	data, err := a.service.SignIn(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewApiError(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, NewApiResponse(data, "login successful"))
}

func createUserFromReq(r SignUpRequest) entity.CreateUser {
	return entity.CreateUser{
		Email:    r.Email,
		Password: r.Password,
	}
}
