package services

import (
	"cart-service/internal/dto"
	"cart-service/internal/grpc"
	"cart-service/internal/models"
	"cart-service/internal/schemas"
	"cart-service/internal/types"

	"github.com/samber/lo"
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
	res, err := grpc.GetProducts(appState.Context, appState.GrpcClientConn, productIds)
	if err != nil {
		return nil, err
	}

	// Transform cart detail
	cartDto := dto.TransformCartDetail(*cart, res.Products)

	return &cartDto, nil
}
