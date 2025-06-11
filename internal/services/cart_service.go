package services

import (
	"cart-service/internal/dto"
	"cart-service/internal/grpc"
	"cart-service/internal/models"
	"cart-service/internal/schemas"
	"cart-service/internal/types"
	"context"
	"errors"

	"github.com/samber/lo"
	"go.uber.org/zap"
)

func GetProductIDsFromCartItems(carts []models.CartItem) []uint64 {
	productIds := lo.Map(carts, func(cart models.CartItem, index int) uint64 {
		return uint64(cart.ProductId)
	})
	return productIds
}

func GetCartItemsProducts(cart *models.Cart, appState *types.AppState) (*schemas.GetCartResponse, error) {
	// Query products from shop-product service
	productIds := GetProductIDsFromCartItems(cart.CartItems)
	res, err := grpc.GetProducts(context.Background(), appState.GrpcShopProductClientConn, productIds)
	if err != nil {
		msg := "failed to get cart item products"
		appState.Log.Error(msg, zap.Error(err))
		return nil, errors.New(msg)
	}

	// Transform cart detail
	cartDto := dto.TransformCartDetail(*cart, res.Products)

	return &cartDto, nil
}
