package service

import (
	"context"
	"testing"
	"time"

	"github.com/ScMofeoluwa/GatherGo/internal/domain/user/entity"
	"github.com/ScMofeoluwa/GatherGo/internal/domain/user/repository/mocks"
	"github.com/ScMofeoluwa/GatherGo/internal/util"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserServiceSuite struct {
	suite.Suite
	ctx         context.Context
	mockRepo    *mocks.UserRepositoryInterface
	userService *UserService
	testUser    *entity.CreateUser
}

func (s *UserServiceSuite) SetupTest() {
	s.ctx = context.Background()
	s.mockRepo = &mocks.UserRepositoryInterface{}
	jwtMaker := &util.JWTMaker{}
	s.userService = NewUserService(s.mockRepo, jwtMaker)

	s.testUser = &entity.CreateUser{
		Email:    "test@example.com",
		Password: "password",
	}
}

func (s *UserServiceSuite) TestSignUp() {
	t := s.T()

	t.Run("success", func(t *testing.T) {
		s.mockRepo.On("Create", mock.Anything, s.testUser).Return(nil).Once()
		err := s.userService.SignUp(s.ctx, s.testUser)
		assert.NoError(t, err)
		s.mockRepo.AssertExpectations(t)
	})

	t.Run("error on signup", func(t *testing.T) {
		s.mockRepo.On("Create", mock.Anything, s.testUser).Return(errors.New("database error")).Once()
		err := s.userService.SignUp(s.ctx, s.testUser)
		assert.Error(t, err)
		s.mockRepo.AssertExpectations(t)
	})
}

func (s *UserServiceSuite) TestSignIn() {
	t := s.T()

	t.Run("success", func(t *testing.T) {
		hashedPwd, _ := util.HashPassword("password")

		mockUser := &entity.User{
			RegisteredAt: time.Now(),
			ID:           "1234",
			Email:        s.testUser.Email,
			Password:     hashedPwd,
			Verified:     true,
		}

		s.mockRepo.On("GetByEmail", mock.Anything, s.testUser.Email).Return(mockUser, nil).Once()

		tokenPair, err := s.userService.SignIn(s.ctx, s.testUser)
		assert.NoError(t, err)
		assert.NotEmpty(t, tokenPair.AccessToken, "access token should not be empty")
		assert.NotEmpty(t, tokenPair.RefreshToken, "refresh token should not be empty")
		s.mockRepo.AssertExpectations(t)
	})

	t.Run("user not found", func(t *testing.T) {
		s.mockRepo.On("GetByEmail", mock.Anything, s.testUser.Email).Return(nil, errors.New("invalid email or password")).Once()

		_, err := s.userService.SignIn(s.ctx, s.testUser)
		assert.Error(t, err)
		assert.Equal(t, "invalid email or password", err.Error())
		s.mockRepo.AssertExpectations(t)
	})

	t.Run("password mismatch", func(t *testing.T) {
		hashedPwd, _ := util.HashPassword("wrong-password")

		mockUser := &entity.User{
			RegisteredAt: time.Now(),
			ID:           "1234",
			Email:        s.testUser.Email,
			Password:     hashedPwd,
			Verified:     true,
		}

		s.mockRepo.On("GetByEmail", mock.Anything, s.testUser.Email).Return(mockUser, nil).Once()

		_, err := s.userService.SignIn(s.ctx, s.testUser)
		assert.Error(t, err)
		assert.Equal(t, "invalid email or password", err.Error())
		s.mockRepo.AssertExpectations(t)
	})
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(UserServiceSuite))
}
