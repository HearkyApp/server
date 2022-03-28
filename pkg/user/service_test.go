package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"server/pkg/domain"
	"server/pkg/domain/mock"
	"testing"
)

func Test_userService_GetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)

	u1 := &domain.User{
		ID: "1",
	}

	repo := mock.NewMockUserRepository(ctrl)
	repo.EXPECT().GetUserByID(gomock.Eq("1")).Return(u1, nil)
	repo.EXPECT().GetUserByID(gomock.Eq("2")).Return(nil, fiber.ErrNotFound)

	s := NewUserService(repo)

	u, err := s.GetUserByID("1")
	assert.NoError(t, err)
	assert.NotNil(t, u)

	u, err = s.GetUserByID("2")
	assert.Error(t, err)
	assert.Nil(t, u)
}

func Test_userService_CreateUser(t *testing.T) {
	t.Run("get user by ID error", func(t *testing.T) {
		uid := "1"
		dto := &domain.CreateUserDTO{
			Name:     "test",
			Username: "test",
		}

		ctrl := gomock.NewController(t)
		repo := mock.NewMockUserRepository(ctrl)
		repo.EXPECT().GetUserByID(gomock.Eq(uid)).Return(nil, fiber.ErrInternalServerError)

		s := NewUserService(repo)
		u, err := s.CreateUser(uid, dto)

		assert.ErrorIs(t, err, fiber.ErrInternalServerError)
		assert.Nil(t, u)
	})

	t.Run("user already registered", func(t *testing.T) {
		uid := "1"
		dto := &domain.CreateUserDTO{
			Name:     "test",
			Username: "test",
		}

		ctrl := gomock.NewController(t)
		repo := mock.NewMockUserRepository(ctrl)
		repo.EXPECT().GetUserByID(gomock.Eq(uid)).Return(&domain.User{}, nil)

		s := NewUserService(repo)
		u, err := s.CreateUser(uid, dto)

		assert.ErrorIs(t, err, domain.ErrUserAlreadyExists)
		assert.Nil(t, u)
	})

	t.Run("username too short", func(t *testing.T) {
		uid := "1"
		dto := &domain.CreateUserDTO{
			Name:     "test",
			Username: "te",
		}

		ctrl := gomock.NewController(t)
		repo := mock.NewMockUserRepository(ctrl)
		repo.EXPECT().GetUserByID(gomock.Eq(uid)).Return(nil, fiber.ErrNotFound)

		s := NewUserService(repo)
		u, err := s.CreateUser(uid, dto)

		assert.ErrorIs(t, err, domain.ErrInvalidUsername)
		assert.Nil(t, u)
	})

	t.Run("username too long", func(t *testing.T) {
		uid := "1"
		dto := &domain.CreateUserDTO{
			Name:     "test",
			Username: "testtesttesttesttesttestt",
		}

		ctrl := gomock.NewController(t)
		repo := mock.NewMockUserRepository(ctrl)
		repo.EXPECT().GetUserByID(gomock.Eq(uid)).Return(nil, fiber.ErrNotFound)

		s := NewUserService(repo)
		u, err := s.CreateUser(uid, dto)

		assert.ErrorIs(t, err, domain.ErrInvalidUsername)
		assert.Nil(t, u)
	})

	t.Run("name too long", func(t *testing.T) {
		uid := "1"
		dto := &domain.CreateUserDTO{
			Name:     "testtesttesttesttesttesttesttesttesttesttesttesttesttesttesttestt",
			Username: "test",
		}

		ctrl := gomock.NewController(t)
		repo := mock.NewMockUserRepository(ctrl)
		repo.EXPECT().GetUserByID(gomock.Eq(uid)).Return(nil, fiber.ErrNotFound)

		s := NewUserService(repo)
		u, err := s.CreateUser(uid, dto)

		assert.ErrorIs(t, err, domain.ErrInvalidName)
		assert.Nil(t, u)
	})

	t.Run("get by username error", func(t *testing.T) {
		uid := "1"
		dto := &domain.CreateUserDTO{
			Name:     "test",
			Username: "test",
		}

		ctrl := gomock.NewController(t)
		repo := mock.NewMockUserRepository(ctrl)
		repo.EXPECT().GetUserByID(gomock.Eq(uid)).Return(nil, fiber.ErrNotFound)
		repo.EXPECT().GetUserByUsername(gomock.Eq(dto.Username)).Return(nil, fiber.ErrInternalServerError)

		s := NewUserService(repo)
		u, err := s.CreateUser(uid, dto)

		assert.ErrorIs(t, err, fiber.ErrInternalServerError)
		assert.Nil(t, u)
	})

	t.Run("username already taken", func(t *testing.T) {
		uid := "1"
		dto := &domain.CreateUserDTO{
			Name:     "test",
			Username: "test",
		}

		ctrl := gomock.NewController(t)
		repo := mock.NewMockUserRepository(ctrl)
		repo.EXPECT().GetUserByID(gomock.Eq(uid)).Return(nil, fiber.ErrNotFound)
		repo.EXPECT().GetUserByUsername(gomock.Eq(dto.Username)).Return(&domain.User{}, nil)

		s := NewUserService(repo)
		u, err := s.CreateUser(uid, dto)

		assert.ErrorIs(t, err, domain.ErrUsernameTaken)
		assert.Nil(t, u)
	})

	t.Run("create user successful", func(t *testing.T) {
		uid := "1"
		dto := &domain.CreateUserDTO{
			Name:     "test",
			Username: "test",
		}

		ctrl := gomock.NewController(t)
		repo := mock.NewMockUserRepository(ctrl)
		repo.EXPECT().GetUserByID(gomock.Eq(uid)).Return(nil, fiber.ErrNotFound)
		repo.EXPECT().GetUserByUsername(gomock.Eq(dto.Username)).Return(nil, fiber.ErrNotFound)
		repo.EXPECT().CreateUser(gomock.Any()).Return(nil)

		s := NewUserService(repo)
		u, err := s.CreateUser(uid, dto)
		assert.NoError(t, err)
		assert.NotNil(t, u)
		assert.NotEmpty(t, u.ID)
		assert.Equal(t, u.Name, dto.Name)
		assert.Equal(t, u.Username, dto.Username)
		assert.NotNil(t, u.CreatedAt)
	})
}
