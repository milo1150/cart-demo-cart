package api

import (
	"cart-service/internal/repositories"
	"cart-service/internal/schemas"
	"cart-service/internal/types"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	cartpkg "github.com/milo1150/cart-demo-pkg/pkg"
)

func CreateCheckoutHandler(c echo.Context, appState *types.AppState) error {
	payload := schemas.CreateCheckoutItem{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, cartpkg.GetSimpleErrorMessage("Invalid payload"))
	}

	validate := validator.New()
	if errMap := cartpkg.ValidateJsonPayload(validate, payload); errMap != nil {
		return c.JSON(http.StatusBadRequest, errMap)
	}

	xUserId := c.Request().Header.Get("X-User-Id")
	userId, err := strconv.Atoi(xUserId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, cartpkg.GetSimpleErrorMessage("invalid UserId"))
	}

	checkoutRepository := repositories.Checkout{DB: appState.DB}
	res, err := checkoutRepository.CreateCheckout(&payload, uint(userId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, cartpkg.GetSimpleErrorMessage(err.Error()))
	}

	return c.JSON(http.StatusOK, res)
}
