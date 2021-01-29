package app

import (
	"management/internal/configs"
	"management/internal/domain"
	"management/internal/logger"
)

type (
	App struct {
		cfg    configs.Config
		logger *logger.Logger
		repo   Repository
	}
	Repository interface {
		// users
		GetUserByEmail(email string) (*domain.User, error)
	}
)

func NewApp(cfg configs.Config, logger *logger.Logger, repo Repository) *App {
	return &App{cfg: cfg, logger: logger, repo: repo}
}