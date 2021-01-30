package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func GetUserIDFromToken(c echo.Context) (string, error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["userID"].(string)
	if userID == "" {
		return "", errors.New("failed to get user id from token claims")
	}
	return userID, nil
}