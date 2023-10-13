package service

import (
	"context"

	"github.com/ScMofeoluwa/GatherGo/internal/domain/user/entity"
	"github.com/ScMofeoluwa/GatherGo/internal/domain/user/repository"
	"github.com/ScMofeoluwa/GatherGo/internal/util"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) RegisterUser(ctx context.Context, user *entity.User) error {
	password, err := util.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = password
	return u.repo.Create(ctx, user)
}
