package user

import (
	"github.com/UpMeetApp/server/pkg/domain"
	"github.com/UpMeetApp/server/pkg/domain/mock"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_userService_GetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)

	uid1 := "1"
	uid2 := "2"

	repo := mock.NewMockUserRepository(ctrl)
	repo.EXPECT().GetUserByID(gomock.Eq(uid1)).Return(&domain.User{ID: uid1}, nil)
	repo.EXPECT().GetUserByID(gomock.Eq(uid2)).Return(nil, fiber.ErrNotFound)

	s := NewUserService(repo)

	u, err := s.GetUserByID(uid1)
	assert.NoError(t, err)
	assert.NotNil(t, u)

	u, err = s.GetUserByID(uid2)
	assert.Error(t, err)
	assert.Nil(t, u)
}

func Test_userService_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := mock.NewMockUserRepository(ctrl)
	s := NewUserService(repo)

	uid := "1"

	// GetUserByID returns error
	repo.EXPECT().GetUserByID(gomock.Eq(uid)).Return(nil, fiber.ErrInternalServerError)
	dto := &domain.CreateUserDTO{}
	u, err := s.CreateUser(uid, dto)
	assert.ErrorIs(t, err, fiber.ErrInternalServerError)
	assert.Nil(t, u)

	// User already registered
	repo.EXPECT().GetUserByID(gomock.Eq(uid)).Return(&domain.User{}, nil)
	u, err = s.CreateUser(uid, dto)
	assert.ErrorIs(t, err, domain.ErrUserAlreadyExists)
	assert.Nil(t, u)

	// Username too short
	dto = &domain.CreateUserDTO{
		Username: "te",
	}
	repo.EXPECT().GetUserByID(gomock.Eq(uid)).Return(nil, fiber.ErrNotFound)
	u, err = s.CreateUser(uid, dto)
	assert.ErrorIs(t, err, domain.ErrInvalidUsername)
	assert.Nil(t, u)

	// Username too long
	dto = &domain.CreateUserDTO{
		Username: "testtesttesttesttesttestt",
	}
	repo.EXPECT().GetUserByID(gomock.Eq(uid)).Return(nil, fiber.ErrNotFound)
	u, err = s.CreateUser(uid, dto)
	assert.ErrorIs(t, err, domain.ErrInvalidUsername)
	assert.Nil(t, u)

	// Name too long
	dto = &domain.CreateUserDTO{
		Name:     "testtesttesttesttesttesttesttesttesttesttesttesttesttesttesttestt",
		Username: "test",
	}
	repo.EXPECT().GetUserByID(gomock.Eq(uid)).Return(nil, fiber.ErrNotFound)
	u, err = s.CreateUser(uid, dto)
	assert.ErrorIs(t, err, domain.ErrInvalidName)
	assert.Nil(t, u)

	// GetByUsername returns error
	dto = &domain.CreateUserDTO{
		Name:     "test",
		Username: "test",
	}
	repo.EXPECT().GetUserByID(gomock.Eq(uid)).Return(nil, fiber.ErrNotFound)
	repo.EXPECT().GetUserByUsername(gomock.Eq(dto.Username)).Return(nil, fiber.ErrInternalServerError)
	u, err = s.CreateUser(uid, dto)
	assert.ErrorIs(t, err, fiber.ErrInternalServerError)
	assert.Nil(t, u)

	// Username already taken
	dto = &domain.CreateUserDTO{
		Name:     "test",
		Username: "test",
	}
	repo.EXPECT().GetUserByID(gomock.Eq(uid)).Return(nil, fiber.ErrNotFound)
	repo.EXPECT().GetUserByUsername(gomock.Eq(dto.Username)).Return(&domain.User{}, nil)
	u, err = s.CreateUser(uid, dto)
	assert.ErrorIs(t, err, domain.ErrUsernameTaken)
	assert.Nil(t, u)

	// CreateUser success
	dto = &domain.CreateUserDTO{
		Name:     "test",
		Username: "test",
	}
	repo.EXPECT().GetUserByID(gomock.Eq(uid)).Return(nil, fiber.ErrNotFound)
	repo.EXPECT().GetUserByUsername(gomock.Eq(dto.Username)).Return(nil, fiber.ErrNotFound)
	repo.EXPECT().CreateUser(gomock.Any()).Return(nil)

	u, err = s.CreateUser(uid, dto)
	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.NotEmpty(t, u.ID)
	assert.Equal(t, u.Name, dto.Name)
	assert.Equal(t, u.Username, dto.Username)
	assert.NotNil(t, u.CreatedAt)
}

func Test_userService_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := mock.NewMockUserRepository(ctrl)
	s := NewUserService(repo)

	uid := "1"

	// GetUserByID returns error
	dto := &domain.UpdateUserDTO{}
	repo.EXPECT().GetUserByID(gomock.Eq(uid)).Return(nil, fiber.ErrInternalServerError)
	u, err := s.UpdateUser(uid, dto)
	assert.ErrorIs(t, err, fiber.ErrInternalServerError)
	assert.Nil(t, u)

	// Username too short
	dto = &domain.UpdateUserDTO{
		Username: "a",
	}
	repo.EXPECT().GetUserByID(gomock.Eq(uid)).Return(&domain.User{}, nil)
	u, err = s.UpdateUser(uid, dto)
	assert.ErrorIs(t, err, domain.ErrInvalidUsername)
	assert.Nil(t, u)

	// Username too long
	dto = &domain.UpdateUserDTO{
		Username: "testtesttesttesttesttestt",
	}
	repo.EXPECT().GetUserByID(gomock.Eq(uid)).Return(&domain.User{}, nil)
	u, err = s.UpdateUser(uid, dto)
	assert.ErrorIs(t, err, domain.ErrInvalidUsername)
	assert.Nil(t, u)

	// GetByUsername returns error
	dto = &domain.UpdateUserDTO{
		Username: "test",
	}
	repo.EXPECT().GetUserByID(gomock.Eq(uid)).Return(&domain.User{}, nil)
	repo.EXPECT().GetUserByUsername(gomock.Eq(dto.Username)).Return(nil, fiber.ErrInternalServerError)
	u, err = s.UpdateUser(uid, dto)
	assert.ErrorIs(t, err, fiber.ErrInternalServerError)
	assert.Nil(t, u)

	// Username is already taken
	dto = &domain.UpdateUserDTO{
		Username: "test",
	}
	repo.EXPECT().GetUserByID(gomock.Eq(uid)).Return(&domain.User{}, nil)
	repo.EXPECT().GetUserByUsername(gomock.Eq(dto.Username)).Return(&domain.User{}, nil)
	u, err = s.UpdateUser(uid, dto)
	assert.ErrorIs(t, err, domain.ErrUsernameTaken)
	assert.Nil(t, u)

	// Username updated
	dto = &domain.UpdateUserDTO{
		Username: "test",
	}
	repo.EXPECT().GetUserByID(gomock.Eq(uid)).Return(&domain.User{}, nil)
	repo.EXPECT().GetUserByUsername(gomock.Eq(dto.Username)).Return(nil, fiber.ErrNotFound)
	repo.EXPECT().UpdateUser(gomock.Any()).Return(nil)
	u, err = s.UpdateUser(uid, dto)
	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, u.Username, dto.Username)

	// Name too long
	dto = &domain.UpdateUserDTO{
		Name: "testtesttesttesttesttesttesttesttesttesttesttesttesttesttesttestt",
	}
	repo.EXPECT().GetUserByID(gomock.Eq(uid)).Return(&domain.User{}, nil)
	u, err = s.UpdateUser(uid, dto)
	assert.ErrorIs(t, err, domain.ErrInvalidName)
	assert.Nil(t, u)

	// Name updated
	dto = &domain.UpdateUserDTO{
		Name: "test",
	}
	repo.EXPECT().GetUserByID(gomock.Eq(uid)).Return(&domain.User{}, nil)
	repo.EXPECT().UpdateUser(gomock.Any()).Return(nil)
	u, err = s.UpdateUser(uid, dto)
	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, u.Name, dto.Name)

	// Bio too long
	dto = &domain.UpdateUserDTO{
		Bio: "testtesttesttesttesttesttesttesttesttesttesttesttesttesttesttestttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttestttesttesttestttesttesttesttesttesttestftesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttestt",
	}
	repo.EXPECT().GetUserByID(gomock.Eq(uid)).Return(&domain.User{}, nil)
	u, err = s.UpdateUser(uid, dto)
	assert.ErrorIs(t, err, domain.ErrInvalidBio)
	assert.Nil(t, u)

	// Bio updated
	dto = &domain.UpdateUserDTO{
		Bio: "test",
	}
	repo.EXPECT().GetUserByID(gomock.Eq(uid)).Return(&domain.User{}, nil)
	repo.EXPECT().UpdateUser(gomock.Any()).Return(nil)
	u, err = s.UpdateUser(uid, dto)
	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, u.Bio, dto.Bio)

	// Age invalid
	dto = &domain.UpdateUserDTO{
		Age: 420,
	}
	repo.EXPECT().GetUserByID(gomock.Eq(uid)).Return(&domain.User{}, nil)
	u, err = s.UpdateUser(uid, dto)
	assert.ErrorIs(t, err, domain.ErrInvalidAge)
	assert.Nil(t, u)

	// Age updated
	dto = &domain.UpdateUserDTO{
		Age: 19,
	}
	repo.EXPECT().GetUserByID(gomock.Eq(uid)).Return(&domain.User{}, nil)
	repo.EXPECT().UpdateUser(gomock.Any()).Return(nil)
	u, err = s.UpdateUser(uid, dto)
	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, u.Age, dto.Age)

	// AgePrivate updated
	dto = &domain.UpdateUserDTO{
		AgePrivate: true,
	}
	repo.EXPECT().GetUserByID(gomock.Eq(uid)).Return(&domain.User{}, nil)
	repo.EXPECT().UpdateUser(gomock.Any()).Return(nil)
	u, err = s.UpdateUser(uid, dto)
	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, u.AgePrivate, dto.AgePrivate)

	// UpdateUser returns error
	dto = &domain.UpdateUserDTO{}
	repo.EXPECT().GetUserByID(gomock.Eq(uid)).Return(&domain.User{}, nil)
	repo.EXPECT().UpdateUser(gomock.Any()).Return(fiber.ErrInternalServerError)
	u, err = s.UpdateUser(uid, dto)
	assert.ErrorIs(t, err, fiber.ErrInternalServerError)
	assert.Nil(t, u)

	// UpdateUser successful
	dto = &domain.UpdateUserDTO{}
	repo.EXPECT().GetUserByID(gomock.Eq(uid)).Return(&domain.User{}, nil)
	repo.EXPECT().UpdateUser(gomock.Any()).Return(nil)
	u, err = s.UpdateUser(uid, dto)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func Test_userService_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := mock.NewMockUserRepository(ctrl)
	s := NewUserService(repo)

	uid := "1"

	// GetUserByID returns error
	repo.EXPECT().GetUserByID(gomock.Eq(uid)).Return(nil, fiber.ErrInternalServerError)
	err := s.DeleteUser(uid)
	assert.ErrorIs(t, err, fiber.ErrInternalServerError)

	// DeleteUser returns error
	repo.EXPECT().GetUserByID(gomock.Eq(uid)).Return(&domain.User{}, nil)
	repo.EXPECT().DeleteUser(gomock.Eq(uid)).Return(fiber.ErrInternalServerError)
	err = s.DeleteUser(uid)
	assert.ErrorIs(t, err, fiber.ErrInternalServerError)

	// DeleteUser successful
	repo.EXPECT().GetUserByID(gomock.Eq(uid)).Return(&domain.User{}, nil)
	repo.EXPECT().DeleteUser(gomock.Eq(uid)).Return(nil)
	err = s.DeleteUser(uid)
	assert.NoError(t, err)
}
