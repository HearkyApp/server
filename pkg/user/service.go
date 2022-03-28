package user

import (
	"github.com/gofiber/fiber/v2"
	"server/pkg/domain"
	"time"
)

type userService struct {
	userRepository domain.UserRepository
}

func NewUserService(userRepository domain.UserRepository) domain.UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) GetUserByID(id string) (*domain.User, error) {
	return s.userRepository.GetUserByID(id)
}

func (s *userService) CreateUser(id string, dto *domain.CreateUserDTO) (*domain.User, error) {
	_, err := s.userRepository.GetUserByID(id)
	if err != fiber.ErrNotFound {
		if err != nil {
			return nil, err
		}
		return nil, domain.ErrUserAlreadyExists
	}

	if len(dto.Username) < 3 || len(dto.Username) > 24 {
		return nil, domain.ErrInvalidUsername
	}
	if len(dto.Name) > 64 {
		return nil, domain.ErrInvalidName
	}

	_, err = s.userRepository.GetUserByUsername(dto.Username)
	if err != fiber.ErrNotFound {
		if err != nil {
			return nil, err
		}
		return nil, domain.ErrUsernameTaken
	}

	u := &domain.User{
		ID:        id,
		Name:      dto.Name,
		Username:  dto.Username,
		CreatedAt: time.Now(),
	}

	err = s.userRepository.CreateUser(u)
	return u, err
}
