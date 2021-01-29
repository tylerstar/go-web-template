package handler

import (
	"github.com/labstack/echo"
	"management/internal/api/models"
	"management/internal/api/utils"
	"management/internal/domain"
	"net/http"
)

func (h *Handler) getUserHandler(c echo.Context) error {
	userID, err := utils.GetUserIDFromToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Message: "token is invalid",
			Error:   err.Error(),
		})
	}

	u, err := h.app.GetUser(&domain.User{ID: userID})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Message: "failed to get user",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, models.GetUserResponse{
		Email: u.Email,
		Username: u.Username,
	})
}
