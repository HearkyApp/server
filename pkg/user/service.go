package user

import (
	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"server/pkg/domain"
	"time"
)

type userService struct {
	userRepository domain.UserRepository
}

func NewUserService(userRepository domain.UserRepository) *userService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) SignUp(dto *domain.SignUpDTO) (*domain.User, error) {
	if len(dto.Email) < 5 || len(dto.Email) > 320 {
		return nil, domain.ErrInvalidEmail
	}
	if len(dto.Username) < 2 || len(dto.Username) > 24 {
		return nil, domain.ErrInvalidUsername
	}
	if len(dto.Name) > 64 {
		return nil, domain.ErrInvalidName
	}

	u, err := s.userRepository.GetUserByEmail(dto.Email)
	if err != nil && err != fiber.ErrNotFound {
		return nil, err
	}
	if u != nil {
		return nil, domain.ErrEmailInUse
	}
	u, err = s.userRepository.GetUserByUsername(dto.Username)
	if err != nil && err != fiber.ErrNotFound {
		return nil, err
	}
	if u != nil {
		return nil, domain.ErrUsernameInUse
	}

	pw, err := hashPassword(dto.Password)
	if err != nil {
		return nil, err
	}

	u = &domain.User{
		ID:        uuid.New().String(),
		Email:     dto.Email,
		Name:      dto.Name,
		Username:  dto.Username,
		Password:  pw,
		CreatedAt: time.Now(),
	}

	err = s.userRepository.CreateUser(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func hashPassword(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		sentry.CaptureException(err)
		zap.L().Error("failed to hash password", zap.Error(err))
		return "", fiber.ErrInternalServerError
	}
	return string(b), nil
}

func comparePassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return domain.ErrInvalidPassword
		}
		sentry.CaptureException(err)
		zap.L().Error("failed to compare password", zap.Error(err))
		return err
	}
	return nil
}
