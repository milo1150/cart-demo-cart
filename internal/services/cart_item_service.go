package services

import (
	"cart-service/internal/repositories"
	"cart-service/internal/schemas"
	"cart-service/internal/types"
)

func AddCartItemToCart(appState *types.AppState, payload schemas.AddCartItemPayload) error {
	// Check if CartItem already existed
	cartItemExists, err := repositories.CartItemExists(appState.DB, payload.ShopId, payload.ProductId)
	if err != nil {
		return err
	}

	// Update CartItem quantity
	if cartItemExists {
		cartItem, err := repositories.FindCartItem(appState.DB, payload.ShopId, payload.ProductId)
		if err != nil {
			return err
		}

		increase := payload.Quantity + cartItem.Quantity

		if err := repositories.UpdateCartItemQuantity(appState.DB, payload.ShopId, payload.ProductId, increase); err != nil {
			return err
		}
	}

	// Create new CartItem
	if !cartItemExists {
		if err := repositories.CreateCartItem(appState.DB, payload); err != nil {
			return err
		}
	}

	return nil
}
