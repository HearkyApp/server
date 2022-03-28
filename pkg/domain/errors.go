package domain

import "github.com/gofiber/fiber/v2"

var (
	ErrInvalidUsername   = fiber.NewError(fiber.StatusBadRequest, "invalid-username")
	ErrInvalidName       = fiber.NewError(fiber.StatusBadRequest, "invalid-name")
	ErrUserAlreadyExists = fiber.NewError(fiber.StatusBadRequest, "user-already-exists")
	ErrUsernameTaken     = fiber.NewError(fiber.StatusBadRequest, "username-taken")
)
