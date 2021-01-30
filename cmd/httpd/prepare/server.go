package prepare

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"management/internal/api/handler"
	"management/internal/app"
	"management/internal/configs"
	"management/internal/logger"
)

func NewServer(cfg configs.Config, logger *logger.Logger, app *app.App) *echo.Echo {
	e := echo.New()

	// Middleware
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	if cfg.Env != "dev" {
		e.Use(middleware.Recover())
	}
	handler.NewHandler(e, cfg, logger, app)
	return e
}