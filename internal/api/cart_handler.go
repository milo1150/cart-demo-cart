package api

import (
	"cart-service/internal/repositories"
	"cart-service/internal/services"
	"cart-service/internal/types"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	cartpkg "github.com/milo1150/cart-demo-pkg/pkg"
)

func GetCartHandler(c echo.Context, appState *types.AppState) error {
	paramId := c.Param("id")

	// Validate param
	cartId, err := strconv.Atoi(paramId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, cartpkg.GetSimpleErrorMessage(err.Error()))
	}

	// Find Cart
	cart, err := repositories.GetCart(appState.DB, uint(cartId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, cartpkg.GetSimpleErrorMessage(err.Error()))
	}

	// Transform Cart Response
	res, err := services.GetCartItemsProducts(cart, appState)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}
