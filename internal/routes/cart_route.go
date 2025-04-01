package routes

import (
	"cart-service/internal/api"
	"cart-service/internal/types"

	"github.com/labstack/echo/v4"
)

func CartRoutes(e *echo.Echo, appState *types.AppState) {
	cartGroup := e.Group("/cart")

	cartGroup.GET("/:id", func(c echo.Context) error {
		return api.GetCartHandler(c, appState)
	})
}
