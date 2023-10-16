package service

import (
	"context"
	"errors"

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

func (u *UserService) SignUp(ctx context.Context, data *entity.CreateUser) error {
	hashedPassword, err := util.HashPassword(data.Password)
	if err != nil {
		return err
	}

	data.Password = hashedPassword
	if err := u.repo.Create(ctx, data); err != nil {
		return err
	}
	return nil
}

func (u *UserService) SignIn(ctx context.Context, data *entity.CreateUser) (string, error) {
	user, err := u.repo.GetByEmail(ctx, data.Email)
	if err != nil {
		return "", err
	}
	if err := util.CheckPassword(user.Password, data.Password); err != nil {
		return "", errors.New("invalid email or password")
	}
	return "", nil
}
