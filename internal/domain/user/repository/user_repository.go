package repository

import (
	"context"

	"github.com/ScMofeoluwa/GatherGo/internal/domain/user/entity"
	"github.com/ScMofeoluwa/GatherGo/internal/infrastructure/database/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserRepository struct {
	sqlc.BaseRepository
}

var _ UserRepositoryInterface = &UserRepository{}

func NewUserRepository(db *sqlc.Queries) *UserRepository {
	return &UserRepository{BaseRepository: sqlc.BaseRepository{DB: db}}
}

func (u *UserRepository) Create(ctx context.Context, user *entity.User) error {
	params := sqlc.CreateUserParams{
		Email:    user.Email,
		Password: user.Password,
	}

	err := u.DB.CreateUser(ctx, params)
	return err
}

func (u *UserRepository) GetByID(ctx context.Context, id string) (*entity.User, error) {
	uuid_, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	dbUser, err := u.DB.GetUserByID(ctx, uuid_)
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

func (u *UserRepository) Update(ctx context.Context, user *entity.User) error {
	params := sqlc.UpdateUserParams{
		Email:    pgtype.Text{String: user.Email},
		Password: pgtype.Text{String: user.Password},
		Verified: pgtype.Bool{Bool: user.Verified},
	}

	err := u.DB.UpdateUser(ctx, params)
	return err
}
