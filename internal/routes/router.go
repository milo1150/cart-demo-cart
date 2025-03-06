package routes

import (
	"cart-service/internal/types"
	"fmt"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo/v4"
)

func RegisterAppRoutes(e *echo.Echo, appState *types.AppState) {
	e.GET("/", func(c echo.Context) error {
		spew.Dump(c.Request().Header)
		email := c.Request().Header.Get("X-User-email")
		fmt.Println("email:", email)
		return c.JSON(http.StatusOK, "Cart service default path")
	})
}
