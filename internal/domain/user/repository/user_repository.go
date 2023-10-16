package repository

import (
	"context"
	"errors"

	"github.com/ScMofeoluwa/GatherGo/internal/domain/user/entity"
	sqlc "github.com/ScMofeoluwa/GatherGo/internal/infrastructure/database/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserRepository struct {
	db *sqlc.Queries
}

var _ UserRepositoryInterface = &UserRepository{}

func NewUserRepository(db *sqlc.Queries) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) Create(ctx context.Context, user *entity.CreateUser) error {
	params := sqlc.CreateUserParams{
		Email:    user.Email,
		Password: user.Password,
	}

	err := u.db.CreateUser(ctx, params)
	if sqlc.ErrorCode(err) == sqlc.UniqueViolation {
		return errors.New("user already exists")
	}
	return err
}

func (u *UserRepository) GetByID(ctx context.Context, id string) (*entity.User, error) {
	uuid_, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	dbUser, err := u.db.GetUserByID(ctx, uuid_)
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

func (u *UserRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	user, err := u.db.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sqlc.ErrRecordNotFound) {
			return nil, errors.New("invalid email or password")
		}
		return nil, err
	}

	return &entity.User{
		ID:           user.ID.String(),
		Email:        user.Email,
		RegisteredAt: user.RegisteredAt.Time,
		Verified:     user.Verified.Bool,
		Password:     user.Password,
	}, nil
}

func (u *UserRepository) Update(ctx context.Context, user *entity.User) error {
	params := sqlc.UpdateUserParams{
		Email:    pgtype.Text{String: user.Email},
		Password: pgtype.Text{String: user.Password},
		Verified: pgtype.Bool{Bool: user.Verified},
	}

	err := u.db.UpdateUser(ctx, params)
	return err
}
