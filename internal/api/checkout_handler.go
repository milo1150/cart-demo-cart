package api

import (
	"cart-service/internal/dto"
	"cart-service/internal/grpc"
	"cart-service/internal/models"
	"cart-service/internal/nats"
	"cart-service/internal/repositories"
	"cart-service/internal/schemas"
	"cart-service/internal/types"
	"cart-service/internal/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	cartpkg "github.com/milo1150/cart-demo-pkg/pkg"
	"github.com/samber/lo"
	"go.uber.org/zap"
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

	// Extract user id from request header
	userId, err := utils.GetUserIdFromRequestHeader(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, cartpkg.GetSimpleErrorMessage(err.Error()))
	}

	// Create checkout and checout_items
	checkoutRepository := repositories.Checkout{DB: appState.DB}
	res, err := checkoutRepository.CreateCheckout(&payload, uint(userId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, cartpkg.GetSimpleErrorMessage(err.Error()))
	}

	// Publish message to payment service
	nats.PublishCreateCheckoutHandler(appState.JS, c.Request().Context(), res, appState.Log)

	return c.JSON(http.StatusOK, res)
}

func GetCheckoutsHandler(c echo.Context, appState *types.AppState) error {
	// Extract user id from request header
	userId, err := utils.GetUserIdFromRequestHeader(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, cartpkg.GetSimpleErrorMessage(err.Error()))
	}

	// Query checkout list
	rc := repositories.Checkout{DB: appState.DB}
	checkouts, err := rc.GetCheckouts(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, cartpkg.GetSimpleErrorMessage(err.Error()))
	}

	// Query payment detail list from payment service
	paymentIds := lo.Map(*checkouts, func(data models.Checkout, index int) uint64 {
		return uint64(data.ID)
	})
	payments, paymentRPCError := grpc.GetPayments(c.Request().Context(), appState.GrpcPaymentClientConn, paymentIds)
	if paymentRPCError != nil {
		appState.Log.Error("Query GetPayments", zap.Error(paymentRPCError))
	}

	// Transform - mapping payment detail into checkout list
	response := dto.TransformCheckoutSlice(*checkouts, payments.PaymentOrders)

	if paymentRPCError == nil {
		return c.JSON(http.StatusOK, response)
	}
	return c.JSON(http.StatusOK, checkouts)
}
