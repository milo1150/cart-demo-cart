package routes

import (
	"cart-service/internal/api"
	"cart-service/internal/types"

	"github.com/labstack/echo/v4"
)

func CheckoutRoutes(e *echo.Echo, appState *types.AppState) {
	cartItemGroups := e.Group("/checkout")

	cartItemGroups.POST("/create", func(c echo.Context) error {
		return api.CreateCheckoutHandler(c, appState)
	})
}
