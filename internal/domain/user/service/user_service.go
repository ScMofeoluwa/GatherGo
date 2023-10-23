package service

import (
	"context"
	"errors"
	"time"

	"github.com/ScMofeoluwa/GatherGo/internal/domain/user/entity"
	"github.com/ScMofeoluwa/GatherGo/internal/domain/user/repository"
	"github.com/ScMofeoluwa/GatherGo/internal/util"
)

type UserService struct {
	repo     *repository.UserRepository
	jwtMaker *util.JWTMaker
}

func NewUserService(repo *repository.UserRepository, jwtMaker *util.JWTMaker) *UserService {
	return &UserService{
		repo:     repo,
		jwtMaker: jwtMaker,
	}
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

func (u *UserService) SignIn(ctx context.Context, data *entity.CreateUser) (util.TokenPair, error) {
	user, err := u.repo.GetByEmail(ctx, data.Email)
	if err != nil {
		return util.TokenPair{}, err
	}
	if err := util.CheckPassword(user.Password, data.Password); err != nil {
		return util.TokenPair{}, errors.New("invalid email or password")
	}
	tokenPair, err := u.jwtMaker.CreateTokenPair(user.Email, time.Hour)
	if err != nil {
		return util.TokenPair{}, err
	}
	return tokenPair, nil
}
