package api

import (
	"cart-service/internal/grpc"
	"cart-service/internal/repositories"
	"cart-service/internal/schemas"
	"cart-service/internal/types"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	cartpkg "github.com/milo1150/cart-demo-pkg/pkg"
)

// TODO: change to add CartItem and create logic for increase quantity.
func CreateCartItemHandler(c echo.Context, appState *types.AppState) error {
	payload := schemas.CreateCartItemPayload{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, cartpkg.GetSimpleErrorMessage(err.Error()))
	}

	validate := validator.New()
	if errMap := cartpkg.ValidateJsonPayload(validate, payload); errMap != nil {
		return c.JSON(http.StatusBadRequest, errMap)
	}

	// Validate product_id (gRPC)
	isExists, err := grpc.ProductExists(appState.Context, appState.GrpcClientConn, payload.ProductId)
	if !isExists {
		return c.JSON(http.StatusBadRequest, cartpkg.GetSimpleErrorMessage("invalid product id"))
	}
	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, cartpkg.GetSimpleErrorMessage(err.Error()))
	}

	// Create CartItem
	if err := repositories.CreateCartItem(appState.DB, payload); err != nil {
		return c.JSON(http.StatusInternalServerError, cartpkg.GetSimpleErrorMessage(err.Error()))
	}

	return c.JSON(http.StatusCreated, http.StatusCreated)
}
