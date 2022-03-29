package server

import (
	"github.com/UpMeetApp/server/pkg/domain"
	"github.com/gofiber/fiber/v2"
)

// HandleGetUserMe handles GET /users/@me
func (s *Server) HandleGetUserMe(ctx *fiber.Ctx) error {
	uid, err := s.FirebaseAuth(ctx)
	if err != nil {
		return err
	}
	u, err := s.userService.GetUserByID(uid)
	if err != nil {
		return err
	}
	return ctx.JSON(u)
}

// HandleCreateUserMe handles POST /users/@me
func (s *Server) HandleCreateUserMe(ctx *fiber.Ctx) error {
	uid, err := s.FirebaseAuth(ctx)
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

// HandleUpdateUserMe handles PATCH /users/@me
func (s *Server) HandleUpdateUserMe(ctx *fiber.Ctx) error {
	uid, err := s.FirebaseAuth(ctx)
	if err != nil {
		return err
	}
	var dto domain.UpdateUserDTO
	err = ctx.BodyParser(&dto)
	if err != nil {
		return fiber.ErrBadRequest
	}
	u, err := s.userService.UpdateUser(uid, &dto)
	if err != nil {
		return err
	}
	return ctx.JSON(u)
}
