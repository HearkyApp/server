package domain

import "github.com/gofiber/fiber/v2"

var (
	//ErrInvalidUsername is returned when the provided username is invalid (too short or too long).
	ErrInvalidUsername = fiber.NewError(fiber.StatusBadRequest, "invalid-username")
	// ErrInvalidName is returned when the provided name is invalid (too short or too long).
	ErrInvalidName = fiber.NewError(fiber.StatusBadRequest, "invalid-name")
	// ErrUserAlreadyExists is returned when the user already exists.
	ErrUserAlreadyExists = fiber.NewError(fiber.StatusBadRequest, "user-already-exists")
	// ErrUsernameTaken is returned when the username is already taken.
	ErrUsernameTaken = fiber.NewError(fiber.StatusBadRequest, "username-taken")
	// ErrInvalidBio is returned when the provided bio is invalid (too long).
	ErrInvalidBio = fiber.NewError(fiber.StatusBadRequest, "invalid-bio")
	// ErrInvalidAge is returned when the provided age is invalid (too high).
	ErrInvalidAge = fiber.NewError(fiber.StatusBadRequest, "invalid-age")
)
