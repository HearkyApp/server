package server

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"strings"
	"time"
)

// FirebaseAuth is a middleware that validates Firebase ID Tokens passed in the Authorization HTTP header.
func (s *Server) FirebaseAuth(ctx *fiber.Ctx) (uid string, err error) {
	h := ctx.Get("Authorization")
	parts := strings.Split(h, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", fiber.ErrBadRequest
	}

	c, ccl := context.WithTimeout(context.Background(), time.Second*10)
	defer ccl()
	t, err := s.fbAuth.VerifyIDToken(c, parts[1])
	if err != nil {
		return "", fiber.ErrUnauthorized
	}
	return t.UID, nil
}
