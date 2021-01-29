package handler

import (
	"github.com/labstack/echo"
	"management/internal/app"
	"management/internal/configs"
	"management/internal/logger"
)

type Handler struct {
	cfg configs.Config
	logger *logger.Logger
	app *app.App
}

func NewHandler(e *echo.Echo, cfg configs.Config, logger *logger.Logger, app *app.App) *Handler {
	h := &Handler{cfg: cfg, logger: logger, app: app}
	h.link(e)
	return h
}

func (h *Handler) link(e *echo.Echo) {
	e.GET("/users", h.getUserHandler)
}