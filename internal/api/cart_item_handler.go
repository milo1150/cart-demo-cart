package api

import (
	"cart-service/internal/grpc"
	"cart-service/internal/schemas"
	"cart-service/internal/services"
	"cart-service/internal/types"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	cartpkg "github.com/milo1150/cart-demo-pkg/pkg"
)

func AddCartItemHandler(c echo.Context, appState *types.AppState) error {
	payload := schemas.AddCartItemPayload{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, cartpkg.GetSimpleErrorMessage(err.Error()))
	}

	// Validate payload
	validate := validator.New()
	if errMap := cartpkg.ValidateJsonPayload(validate, payload); errMap != nil {
		return c.JSON(http.StatusBadRequest, errMap)
	}

	// Validate product_id (gRPC)
	isExists, err := grpc.ProductExists(appState.Context, appState.GrpcClientConn, payload.ProductId)
	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, cartpkg.GetSimpleErrorMessage(err.Error()))
	}
	if !isExists {
		return c.JSON(http.StatusBadRequest, cartpkg.GetSimpleErrorMessage("invalid product id"))
	}

	// Handle should create new cart item or update quantity
	if err := services.AddCartItemToCart(appState, payload); err != nil {
		return c.JSON(http.StatusInternalServerError, cartpkg.GetSimpleErrorMessage(err.Error()))
	}

	return c.JSON(http.StatusOK, http.StatusOK)
}
