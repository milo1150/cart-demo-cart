package api

import (
	"cart-service/internal/types"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateCartItemHandler(c echo.Context, appState *types.AppState) error {
	return c.JSON(http.StatusTeapot, "TEST TEST")
}
