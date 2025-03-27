package routes

import (
	"cart-service/internal/types"

	"github.com/labstack/echo/v4"
)

func RegisterAppRoutes(e *echo.Echo, appState *types.AppState) {
	CartRoutes(e, appState)
	CartItemRoutes(e, appState)
}
