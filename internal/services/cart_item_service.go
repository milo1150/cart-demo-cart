package services

import (
	"cart-service/internal/repositories"
	"cart-service/internal/schemas"

	"gorm.io/gorm"
)

type CartItem struct {
	DB *gorm.DB
}

func (c *CartItem) AddCartItemToCart(payload schemas.AddCartItemPayload, cartId uint) error {
	cartItemRepo := repositories.CartItem{DB: c.DB}

	// Check if CartItem already existed
	cartItemExists, err := cartItemRepo.CartItemExists(payload.ShopId, payload.ProductId)
	if err != nil {
		return err
	}

	// Update CartItem quantity if item already exists in cart
	if cartItemExists {
		if err := cartItemRepo.UpdateCartItemQuantity(payload.ShopId, payload.ProductId, payload.Quantity); err != nil {
			return err
		}
	}

	// Create new CartItem
	if !cartItemExists {
		if err := cartItemRepo.CreateCartItem(payload, cartId); err != nil {
			return err
		}
	}

	return nil
}

func (c *CartItem) AddCartItemsToCart(payload schemas.AddCartItemSlicesPayload, userId uint) error {
	db := c.DB

	// Find Cart id
	cartRepo := repositories.Cart{DB: c.DB}
	cartId, err := cartRepo.GetCartIdByUserId(userId)
	if err != nil {
		return err
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		cartItemTx := CartItem{DB: tx}
		for _, cartItemPayload := range payload.CartItems {
			err := cartItemTx.AddCartItemToCart(cartItemPayload, *cartId)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return err
}

func (c *CartItem) RemoveCartItem(userId uint, payload schemas.RemoveCartItemPayload) error {
	cartItemRepo := repositories.CartItem{DB: c.DB}
	cartRepo := repositories.Cart{DB: c.DB}

	cartId, err := cartRepo.GetCartIdByUserId(userId)
	if err != nil {
		return err
	}

	if err := cartItemRepo.RemoveCartItem(payload.ShopId, payload.ProductId, *cartId); err != nil {
		return err
	}

	return nil
}
