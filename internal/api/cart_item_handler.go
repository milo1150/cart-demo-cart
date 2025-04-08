package api

import (
	"cart-service/internal/schemas"
	"cart-service/internal/services"
	"cart-service/internal/types"
	"cart-service/internal/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	cartpkg "github.com/milo1150/cart-demo-pkg/pkg"
)

func AddCartItemHandler(c echo.Context, appState *types.AppState) error {
	payload := schemas.AddCartItemSlicesPayload{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, cartpkg.GetSimpleErrorMessage(err.Error()))
	}

	// Validate payload
	validate := validator.New()
	if errMap := cartpkg.ValidateJsonPayload(validate, payload); errMap != nil {
		return c.JSON(http.StatusBadRequest, errMap)
	}

	// Extract user id from request header
	userId, err := utils.GetUserIdFromRequestHeader(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, cartpkg.GetSimpleErrorMessage(err.Error()))
	}

	// Handle should create new cart item or update quantity
	cartItemService := services.CartItem{AppState: appState}
	if err := cartItemService.AddCartItemsToCart(payload, userId); err != nil {
		return c.JSON(http.StatusInternalServerError, cartpkg.GetSimpleErrorMessage(err.Error()))
	}

	return c.JSON(http.StatusOK, http.StatusOK)
}
