package service

import (
	"context"

	"github.com/ScMofeoluwa/GatherGo/internal/domain/user/entity"
	"github.com/ScMofeoluwa/GatherGo/internal/util"
)

type UserServiceInterface interface {
	SignUp(ctx context.Context, data *entity.CreateUser) error
	SignIn(ctx context.Context, data *entity.CreateUser) (util.TokenPair, error)
}
