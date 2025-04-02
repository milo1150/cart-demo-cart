package utils

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetUserIdFromRequestHeader(c echo.Context) (uint, error) {
	xUserId := c.Request().Header.Get("X-User-Id")
	userId, err := strconv.Atoi(xUserId)
	if err != nil {
		return 0, err
	}
	return uint(userId), nil
}
