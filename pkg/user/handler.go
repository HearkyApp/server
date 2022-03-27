package user

import (
	"github.com/gofiber/fiber/v2"
	"server/pkg/domain"
)

func HandleSignUp(userService *userService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var dto domain.SignUpDTO
		err := ctx.BodyParser(&dto)
		if err != nil {
			return fiber.ErrBadRequest
		}
		u, err := userService.SignUp(&dto)
		if err != nil {
			return err
		}
		return ctx.JSON(u)
	}
}
