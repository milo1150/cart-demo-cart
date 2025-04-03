package services

import (
	"cart-service/internal/dto"
	"cart-service/internal/grpc"
	"cart-service/internal/models"
	"cart-service/internal/nats"
	"cart-service/internal/repositories"
	"cart-service/internal/schemas"
	"cart-service/internal/types"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type CheckoutService struct {
	AppState *types.AppState
}

func (c *CheckoutService) CreateCheckoutHandler(ctx echo.Context, payload schemas.CreateCheckoutItem, userId uint) (*models.Checkout, error) {
	// Create checkout and checkout_items
	checkoutRepository := repositories.Checkout{DB: c.AppState.DB}
	checkout, err := checkoutRepository.CreateCheckout(&payload, uint(userId))
	if err != nil {
		return nil, err
	}

	// Update total_paid_amount of checkout
	if err := checkoutRepository.UpdateCheckoutTotalPaidAmount(checkout.ID, userId); err != nil {
		c.AppState.Log.Error("Update total_paid_amount error", zap.Error(err))
	}

	// Publish message to payment service
	nats.PublishCreateCheckoutHandler(c.AppState.JS, ctx.Request().Context(), checkout, c.AppState.Log)

	// TODO: update stock, if payment got cancel then refund stock.

	return checkout, nil
}

func (c *CheckoutService) GetCheckoutsHandler(ctx echo.Context, userId uint) (*schemas.CheckoutItemSliceResponse, error) {
	// Query checkout list
	rc := repositories.Checkout{DB: c.AppState.DB}
	checkouts, err := rc.GetCheckouts(userId)
	if err != nil {
		return nil, err
	}

	// Query payment detail list from payment service
	paymentIds := lo.Map(*checkouts, func(data models.Checkout, index int) uint64 {
		return uint64(data.ID)
	})
	payments, paymentRPCError := grpc.GetPayments(ctx.Request().Context(), c.AppState.GrpcPaymentClientConn, paymentIds)
	if paymentRPCError != nil {
		c.AppState.Log.Error("Query GetPayments", zap.Error(paymentRPCError))
	}

	// Transform - mapping payment detail into checkout list
	response := dto.TransformCheckoutSlice(*checkouts, payments.PaymentOrders)

	return &response, nil
}
