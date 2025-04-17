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

func (c *CheckoutService) syncCartItemQuantityWithStock(shopId uint, product models.CheckoutItemProductJson) error {
	cartItemRepo := repositories.CartItem{DB: c.AppState.DB}
	quantity := product.Stock - product.Quantity

	if quantity <= 0 {
		cartItem, err := cartItemRepo.FindCartItem(shopId, product.Id)
		if err != nil {
			return err
		}
		if err := cartItemRepo.RemoveCartItem(shopId, product.Id, cartItem.CartID); err != nil {
			return err
		}
	}

	if quantity > 0 {
		if err := cartItemRepo.UpdateCartItemQuantity(shopId, product.Id, uint(quantity)); err != nil {
			return err
		}
	}

	return nil
}

func (c *CheckoutService) CreateCheckout(ctx echo.Context, payload schemas.CreateCheckoutItem, userId uint) (*models.Checkout, error) {
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

	// update cart_items quantity
	for _, checkoutItem := range payload.CheckoutItems {
		for _, product := range checkoutItem.Products {
			if err := c.syncCartItemQuantityWithStock(checkoutItem.Shop.Id, product); err != nil {
				return nil, err
			}
		}
	}

	// Publish message to payment service
	nats.PublishCreateCheckoutHandler(c.AppState.JS, ctx.Request().Context(), checkout, c.AppState.Log)

	return checkout, nil
}

func (c *CheckoutService) GetCheckouts(ctx echo.Context, userId uint) (*schemas.CheckoutItemSliceResponse, error) {
	// Query checkout list
	rc := repositories.Checkout{DB: c.AppState.DB}
	checkouts, err := rc.GetCheckouts(userId)
	if err != nil {
		return nil, err
	}

	// Filtered payment for query
	paymentIds := lo.Map(*checkouts, func(checkout models.Checkout, index int) uint64 {
		return uint64(checkout.PaymentId)
	})

	// Query payment detail list from payment service
	payments, paymentRPCError := grpc.GetPayments(ctx.Request().Context(), c.AppState.GrpcPaymentClientConn, paymentIds)
	if paymentRPCError != nil {
		c.AppState.Log.Error("Query GetPayments", zap.Error(paymentRPCError))
	}

	// Transform - mapping payment detail into checkout list
	response := dto.TransformCheckoutSlice(*checkouts, payments.PaymentOrders)

	return &response, nil
}
