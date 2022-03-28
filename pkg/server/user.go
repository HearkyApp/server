package server

import (
	"github.com/gofiber/fiber/v2"
	"server/pkg/domain"
)

func (s *Server) HandleGetUserMe(ctx *fiber.Ctx) error {
	uid, err := s.JWTAuth(ctx)
	if err != nil {
		return err
	}
	u, err := s.userService.GetUserByID(uid)
	if err != nil {
		return err
	}
	return ctx.JSON(u)
}

func (s *Server) HandleCreateUserMe(ctx *fiber.Ctx) error {
	uid, err := s.JWTAuth(ctx)
	if err != nil {
		return err
	}
	var dto domain.CreateUserDTO
	err = ctx.BodyParser(&dto)
	if err != nil {
		return fiber.ErrBadRequest
	}
	u, err := s.userService.CreateUser(uid, &dto)
	if err != nil {
		return err
	}
	return ctx.JSON(u)
}
