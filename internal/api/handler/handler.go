package handler

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
	jwtMiddleware := middleware.JWT([]byte(h.cfg.Auth.Secret))
	v1 := e.Group("/api")

	publicUsers := v1.Group("/users")
	publicUsers.POST("/token", h.getTokenHandler)
	publicUsers.POST("", h.createUserHandler)

	privateUsers := v1.Group("/users", jwtMiddleware)
	privateUsers.GET("", h.getUserHandler)
}