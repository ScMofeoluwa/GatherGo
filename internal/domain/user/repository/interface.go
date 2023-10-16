package repository

import (
	"context"

	"github.com/ScMofeoluwa/GatherGo/internal/domain/user/entity"
)

type UserRepositoryInterface interface {
	Create(ctx context.Context, user *entity.CreateUser) error
	GetByID(ctx context.Context, id string) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
}
