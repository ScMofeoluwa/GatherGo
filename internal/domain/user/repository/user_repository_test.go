package repository

import (
	"context"
	"log"
	"testing"

	"github.com/ScMofeoluwa/GatherGo/internal/domain/user/entity"
	"github.com/ScMofeoluwa/GatherGo/internal/infrastructure/database/sqlc"
	"github.com/ScMofeoluwa/GatherGo/internal/infrastructure/testutils"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserRepoTestSuite struct {
	suite.Suite
	pgContainer *testutils.PostgresContainer
	repository  UserRepositoryInterface
	ctx         context.Context
}

func (suite *UserRepoTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	pgContainer, err := testutils.CreatePostgresContainer(suite.ctx)
	if err != nil {
		log.Fatal(err)
	}

	suite.pgContainer = pgContainer
	connPool, err := pgxpool.New(suite.ctx, suite.pgContainer.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}

	dbHandler := sqlc.NewSQLDbHandler(connPool)
	db := dbHandler.(*sqlc.SQLDbHandler).Queries
	suite.repository = NewUserRepository(db)
}

func (suite *UserRepoTestSuite) TearDownSuite() {
	if err := suite.pgContainer.Terminate(suite.ctx); err != nil {
		log.Fatalf("error terminating postgres container: %s", err)
	}
}

func (suite *UserRepoTestSuite) TestCreateCustomer() {
	t := suite.T()

	user := &entity.CreateUser{
		Email:    "test@example.com",
		Password: "password",
	}

	t.Run("success", func(t *testing.T) {
		err := suite.repository.Create(suite.ctx, user)
		assert.NoError(t, err)
	})

	t.Run("email already taken", func(t *testing.T) {
		err := suite.repository.Create(suite.ctx, user)
		assert.Error(t, err)
	})
}

func (suite *UserRepoTestSuite) TestGetCustomerByEmail() {
	t := suite.T()

	t.Run("success", func(t *testing.T) {
		user, err := suite.repository.GetByEmail(suite.ctx, "test@example.com")
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, "test@example.com", user.Email)
	})

	t.Run("user does not exist", func(t *testing.T) {
		_, err := suite.repository.GetByEmail(suite.ctx, "nonexistent@example.com")
		assert.Error(t, err)
	})
}

func TestUserRepoSuite(t *testing.T) {
	suite.Run(t, new(UserRepoTestSuite))
}
