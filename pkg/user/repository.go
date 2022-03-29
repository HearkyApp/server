package user

import (
	"github.com/UpMeetApp/server/pkg/domain"
	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new user repository instance.
func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(u *domain.User) error {
	err := r.db.Create(u).Error
	if err != nil {
		sentry.CaptureException(err)
		zap.L().Error("failed to create user", zap.Error(err))
		return fiber.ErrInternalServerError
	}
	return nil
}

func (r *userRepository) GetUserByID(id string) (*domain.User, error) {
	u := &domain.User{}
	err := r.db.Where("id = ?", id).First(u).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fiber.ErrNotFound
		}
		sentry.CaptureException(err)
		zap.L().Error("failed to get user by id", zap.Error(err))
		return nil, fiber.ErrInternalServerError
	}
	return u, nil
}

func (r *userRepository) GetUserByEmail(email string) (*domain.User, error) {
	u := &domain.User{}
	err := r.db.Where("email = ?", email).First(u).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fiber.ErrNotFound
		}
		sentry.CaptureException(err)
		zap.L().Error("failed to get user by email", zap.Error(err))
		return nil, fiber.ErrInternalServerError
	}
	return u, nil
}

func (r *userRepository) GetUserByUsername(username string) (*domain.User, error) {
	u := &domain.User{}
	err := r.db.Where("username = ?", username).First(u).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fiber.ErrNotFound
		}
		sentry.CaptureException(err)
		zap.L().Error("failed to get user by username", zap.Error(err))
		return nil, fiber.ErrInternalServerError
	}
	return u, nil
}

func (r *userRepository) SearchUsersByName(name string) ([]*domain.User, error) {
	var users []*domain.User
	err := r.db.Where("name LIKE ?", "%"+name+"%").Find(&users).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fiber.ErrNotFound
		}
		sentry.CaptureException(err)
		zap.L().Error("failed to search users by name", zap.Error(err))
		return nil, fiber.ErrInternalServerError
	}
	return users, nil
}

func (r *userRepository) SearchUsersByUsername(username string) ([]*domain.User, error) {
	var users []*domain.User
	err := r.db.Where("username LIKE ?", "%"+username+"%").Find(&users).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fiber.ErrNotFound
		}
		sentry.CaptureException(err)
		zap.L().Error("failed to search users by username", zap.Error(err))
		return nil, fiber.ErrInternalServerError
	}
	return users, nil
}

func (r *userRepository) UpdateUser(u *domain.User) error {
	err := r.db.Save(u).Error
	if err != nil {
		sentry.CaptureException(err)
		zap.L().Error("failed to update user", zap.Error(err))
		return fiber.ErrInternalServerError
	}
	return nil
}

func (r *userRepository) DeleteUser(id string) error {
	err := r.db.Delete(&domain.User{}, "id = ?", id).Error
	if err != nil {
		sentry.CaptureException(err)
		zap.L().Error("failed to delete user", zap.Error(err))
		return fiber.ErrInternalServerError
	}
	return nil
}
