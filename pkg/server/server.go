package server

import (
	"context"
	"encoding/base64"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"google.golang.org/api/option"
	"server/pkg/config"
	"server/pkg/domain"
)

type Server struct {
	app         *fiber.App
	fbApp       *firebase.App
	fbAuth      *auth.Client
	cfg         *config.Config
	userService domain.UserService
}

func New(cfg *config.Config, userService domain.UserService) *Server {
	creds, err := base64.StdEncoding.DecodeString(cfg.FirebaseCredentials)
	if err != nil {
		sentry.CaptureException(err)
		zap.L().Fatal("failed to decode firebase credentials", zap.Error(err))
	}
	fbApp, err := firebase.NewApp(context.Background(), &firebase.Config{}, option.WithCredentialsJSON(creds))
	if err != nil {
		sentry.CaptureException(err)
		zap.L().Fatal("failed to create firebase app", zap.Error(err))
	}
	fbAuth, err := fbApp.Auth(context.Background())
	if err != nil {
		sentry.CaptureException(err)
		zap.L().Fatal("failed to create firebase auth client", zap.Error(err))
	}
	app := fiber.New()

	s := &Server{
		app:         app,
		fbApp:       fbApp,
		fbAuth:      fbAuth,
		cfg:         cfg,
		userService: userService,
	}

	api := app.Group("/api")
	apiV1 := api.Group("/v1")

	apiV1.Get("/users/@me", s.HandleGetUserMe)
	apiV1.Post("/users/@me", s.HandleCreateUserMe)

	return s
}

func (s *Server) Start(bindAddress string) {
	err := s.app.Listen(bindAddress)
	if err != nil {
		sentry.CaptureException(err)
		zap.L().Fatal("failed to start server", zap.Error(err))
	}
}
