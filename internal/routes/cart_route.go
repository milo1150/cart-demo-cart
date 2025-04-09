package routes

import (
	"cart-service/internal/api"
	"cart-service/internal/types"

	"github.com/labstack/echo/v4"
)

func CartRoutes(e *echo.Echo, appState *types.AppState) {
	cartGroup := e.Group("")

	cartGroup.GET("/get-cart", func(c echo.Context) error {
		return api.GetCartUUIDHandler(c, appState)
	})

	cartGroup.GET("/:cart-uuid", func(c echo.Context) error {
		return api.GetCartHandler(c, appState)
	})
}
