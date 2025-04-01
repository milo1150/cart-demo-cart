package routes

import (
	"cart-service/internal/api"
	"cart-service/internal/types"

	"github.com/labstack/echo/v4"
)

func CheckoutRoutes(e *echo.Echo, appState *types.AppState) {
	checkoutGroups := e.Group("/checkout")

	checkoutGroups.POST("/create", func(c echo.Context) error {
		return api.CreateCheckoutHandler(c, appState)
	})
}
