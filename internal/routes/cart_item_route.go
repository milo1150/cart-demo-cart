package routes

import (
	"cart-service/internal/api"
	"cart-service/internal/types"

	"github.com/labstack/echo/v4"
)

func CartItemRoutes(e *echo.Echo, appState *types.AppState) {
	cartItemGroups := e.Group("/cart-item")

	cartItemGroups.POST("/add", func(c echo.Context) error {
		return api.AddCartItemHandler(c, appState)
	})

	cartItemGroups.POST("/remove", func(c echo.Context) error {
		return api.RemoveCartItemHandler(c, appState)
	})
}
