package handler

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"management/internal/api/models"
	"management/internal/api/utils"
	"management/internal/domain"
	"net/http"
)

func (h *Handler) getUserHandler(c echo.Context) error {
	userID, err := utils.GetUserIDFromToken(c)
	log.Info(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Message: "token is invalid",
			Error:   err.Error(),
		})
	}

	u, err := h.app.GetUserByID(&domain.User{ID: userID})
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

func (h *Handler) createUserHandler(c echo.Context) error {
	request := new(models.CreateUserRequest)
	if err := c.Bind(request); err != nil {
		return err
	}

	err := h.app.CreateUser(&domain.User{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Message: "failed to create user",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.OKResponse{
		Status:  "ok",
		Message: "user created",
	})
}

func (h *Handler) getTokenHandler(c echo.Context) error {
	request := new(models.GetTokenRequest)
	if err := c.Bind(request); err != nil {
		return nil
	}

	u := &domain.User{
		Email:    request.Email,
		Password: request.Password,
	}

	valid, err := h.app.ValidateUserPassword(u)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Message: "failed to authorize",
			Error:   err.Error(),
		})
	}
	if !valid {
		return c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Message: "password is invalid",
		})
	}

	token, err := h.app.GetUserToken(u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Message: "failed to get token",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.GetTokenResponse{
		Token: token,
	})
}
