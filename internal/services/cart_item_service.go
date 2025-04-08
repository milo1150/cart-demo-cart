package services

import (
	"cart-service/internal/repositories"
	"cart-service/internal/schemas"
	"cart-service/internal/types"

	"gorm.io/gorm"
)

type CartItem struct {
	AppState *types.AppState
}

func (c *CartItem) AddCartItemToCart(db *gorm.DB, payload schemas.AddCartItemPayload, cartId uint) error {
	// Check if CartItem already existed
	cartItemExists, err := repositories.CartItemExists(db, payload.ShopId, payload.ProductId)
	if err != nil {
		return err
	}

	// Update CartItem quantity if item already exists in cart
	if cartItemExists {
		if err := repositories.UpdateCartItemQuantity(db, payload.ShopId, payload.ProductId, payload.Quantity); err != nil {
			return err
		}
	}

	// Create new CartItem
	if !cartItemExists {
		if err := repositories.CreateCartItem(db, payload, cartId); err != nil {
			return err
		}
	}

	return nil
}

func (c *CartItem) AddCartItemsToCart(payload schemas.AddCartItemSlicesPayload, userId uint) error {
	db := c.AppState.DB

	// Find Cart id
	cartRepo := repositories.Cart{DB: c.AppState.DB}
	cartId, err := cartRepo.GetCartIdByUserId(userId)
	if err != nil {
		return err
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		for _, cartItemPayload := range payload.CartItems {
			err := c.AddCartItemToCart(tx, cartItemPayload, *cartId)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return err
}
