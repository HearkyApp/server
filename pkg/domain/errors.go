package domain

import "github.com/gofiber/fiber/v2"

var (
	ErrInvalidEmail    = fiber.NewError(fiber.StatusBadRequest, "invalid-email")
	ErrEmailInUse      = fiber.NewError(fiber.StatusBadRequest, "email-in-use")
	ErrInvalidUsername = fiber.NewError(fiber.StatusBadRequest, "invalid-username")
	ErrUsernameInUse   = fiber.NewError(fiber.StatusBadRequest, "username-in-use")
	ErrInvalidName     = fiber.NewError(fiber.StatusBadRequest, "invalid-name")
	ErrInvalidPassword = fiber.NewError(fiber.StatusBadRequest, "invalid-password")
)
