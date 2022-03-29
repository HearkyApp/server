package user

import (
	"github.com/UpMeetApp/server/pkg/domain"
	"github.com/gofiber/fiber/v2"
	"time"
)

type userService struct {
	userRepository domain.UserRepository
}

// NewUserService creates a new user service instance.
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

	if len(dto.Username) < domain.UsernameMinLength || len(dto.Username) > domain.UsernameMaxLength {
		return nil, domain.ErrInvalidUsername
	}
	if len(dto.Name) > domain.UserNameMaxLength {
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

func (s *userService) UpdateUser(id string, dto *domain.UpdateUserDTO) (*domain.User, error) {
	u, err := s.userRepository.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	// Update Username
	if len(dto.Username) > 0 {
		if len(dto.Username) < domain.UsernameMinLength || len(dto.Username) > domain.UsernameMaxLength {
			return nil, domain.ErrInvalidUsername
		}
		_, err = s.userRepository.GetUserByUsername(dto.Username)
		if err != fiber.ErrNotFound {
			if err != nil {
				return nil, err
			}
			return nil, domain.ErrUsernameTaken
		}
		u.Username = dto.Username
	}

	// Update Name
	if len(dto.Name) > 0 {
		if len(dto.Name) > domain.UserNameMaxLength {
			return nil, domain.ErrInvalidName
		}
		u.Name = dto.Name
	}

	// Update Bio
	if len(dto.Bio) > 0 {
		if len(dto.Bio) > domain.UserBioMaxLength {
			return nil, domain.ErrInvalidBio
		}
		u.Bio = dto.Bio
	}

	// Update Age
	if dto.Age > 0 {
		if dto.Age > domain.UserMaxAge && dto.Age != 420 {
			return nil, domain.ErrInvalidAge
		}
		u.Age = dto.Age
	}

	// Update Age private
	if dto.AgePrivate != u.AgePrivate {
		u.AgePrivate = dto.AgePrivate
	}

	err = s.userRepository.UpdateUser(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}
