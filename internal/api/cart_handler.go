package api

import (
	"cart-service/internal/repositories"
	"cart-service/internal/services"
	"cart-service/internal/types"
	"cart-service/internal/utils"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	cartpkg "github.com/milo1150/cart-demo-pkg/pkg"
)

func GetCartHandler(c echo.Context, appState *types.AppState) error {
	cartUuidParam := c.Param("cart-uuid")

	// Validate uuid param
	cartUuid, err := uuid.Parse(cartUuidParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, cartpkg.GetSimpleErrorMessage(err.Error()))
	}

	// Find Cart
	cartRepo := repositories.Cart{DB: appState.DB}
	cart, err := cartRepo.GetCartByUuid(appState.DB, cartUuid)
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

func GetCartUUIDHandler(c echo.Context, appState *types.AppState) error {
	// Extract user id from request header
	userId, err := utils.GetUserIdFromRequestHeader(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, cartpkg.GetSimpleErrorMessage(err.Error()))
	}

	cartRepo := repositories.Cart{DB: appState.DB}
	cartUuid, err := cartRepo.GetCartUuidByUserId(uint(userId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, cartpkg.GetSimpleErrorMessage(err.Error()))
	}

	return c.JSON(http.StatusOK, cartUuid)
}
