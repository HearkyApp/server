package server

import (
	"context"
	"encoding/base64"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/UpMeetApp/server/pkg/config"
	"github.com/UpMeetApp/server/pkg/domain"
	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"google.golang.org/api/option"
)

// Server is the main server struct.
type Server struct {
	app         *fiber.App
	fbApp       *firebase.App
	fbAuth      *auth.Client
	cfg         *config.Config
	userService domain.UserService
}

// New created a new (web) server instance.
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
	apiV1.Patch("/users/@me", s.HandleUpdateUserMe)
	apiV1.Delete("/users/@me", s.HandleDeleteUserMe)

	return s
}

// Start starts the (web) server.
func (s *Server) Start(bindAddress string) {
	err := s.app.Listen(bindAddress)
	if err != nil {
		sentry.CaptureException(err)
		zap.L().Fatal("failed to start server", zap.Error(err))
	}
}
