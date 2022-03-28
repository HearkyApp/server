package main

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"server/pkg/config"
	"server/pkg/domain"
	"server/pkg/server"
	"server/pkg/user"
)

func main() {
	cfg := config.LoadConfig()

	var l *zap.Logger
	var err error
	if cfg.Debug {
		l, err = zap.NewDevelopment()
	} else {
		l, err = zap.NewProduction()
	}
	if err != nil {
		sentry.CaptureException(err)
		log.Fatalf("failed to initialize logger: %v", err)
	}
	zap.ReplaceGlobals(l)

	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", cfg.PostgresHost, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDatabase, cfg.PostgresPort, cfg.PostgresSSLMode)), &gorm.Config{})
	if err != nil {
		sentry.CaptureException(err)
		zap.L().Fatal("failed to connect to database", zap.Error(err))
	}

	err = db.AutoMigrate(domain.User{})
	if err != nil {
		sentry.CaptureException(err)
		zap.L().Fatal("failed to migrate database", zap.Error(err))
	}

	userRepository := user.NewUserRepository(db)
	userService := user.NewUserService(userRepository)

	s := server.New(cfg, userService)
	s.Start(cfg.BindAddress)
}
