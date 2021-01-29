package utils

import (
	"errors"

	"github.com/labstack/echo"
)

func GetUserIDFromToken(c echo.Context) (string, error) {
	id, ok := c.Get("user").(string)
	if !ok {
		return "", errors.New("failed to get user ID from token claims")
	}
	return id, nil
}
