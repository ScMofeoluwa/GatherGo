package repository

import (
	"context"

	"github.com/ScMofeoluwa/GatherGo/internal/domain/user/entity"
	db "github.com/ScMofeoluwa/GatherGo/internal/infrastructure/database/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserRepository struct {
	Queries *db.Queries
}

var _ UserRepositoryInterface = &UserRepository{}

func NewUserRepository(queries *db.Queries) *UserRepository {
	return &UserRepository{Queries: queries}
}

func (r *UserRepository) Create(ctx context.Context, user *entity.User) error {
	params := db.CreateUserParams{
		Email:    user.Email,
		Password: user.Password,
	}

	err := r.Queries.CreateUser(ctx, params)
	return err
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*entity.User, error) {
	u, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	dbUser, err := r.Queries.GetUserByID(ctx, u)
	if err != nil {
		return nil, err
	}

	return &entity.User{
		ID:           dbUser.ID.String(),
		Email:        dbUser.Email,
		RegisteredAt: dbUser.RegisteredAt.Time,
		Verified:     dbUser.Verified.Bool,
	}, nil
}

func (r *UserRepository) Update(ctx context.Context, user *entity.User) error {
	params := db.UpdateUserParams{
		Email:    pgtype.Text{String: user.Email},
		Password: pgtype.Text{String: user.Password},
		Verified: pgtype.Bool{Bool: user.Verified},
	}

	err := r.Queries.UpdateUser(ctx, params)
	return err
}
