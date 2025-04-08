package repositories

import (
	"cart-service/internal/models"
	"cart-service/internal/schemas"

	"gorm.io/gorm"
)

type CartItem struct {
	DB *gorm.DB
}

func (c *CartItem) CreateCartItem(payload schemas.AddCartItemPayload, cartId uint) error {
	newCartItem := models.CartItem{
		Quantity:  payload.Quantity,
		CartID:    cartId,
		ProductId: payload.ProductId,
		ShopId:    payload.ShopId,
	}

	if err := c.DB.Create(&newCartItem).Error; err != nil {
		return err
	}

	return nil
}

func (c *CartItem) CartItemExists(shopId, productId uint) (bool, error) {
	cartItem := &models.CartItem{}

	query := c.DB.Where("shop_id = ? AND product_id = ?", shopId, productId).Find(cartItem)
	if query.Error != nil {
		return false, query.Error
	}

	if cartItem.CartID == 0 {
		return false, nil
	}

	return true, nil
}

func (c *CartItem) FindCartItem(shopId, productId uint) (*models.CartItem, error) {
	cartItem := &models.CartItem{}

	query := c.DB.Where("shop_id = ? AND product_id = ?", shopId, productId).First(cartItem)
	if query.Error != nil {
		return nil, query.Error
	}

	return cartItem, nil
}

func (c *CartItem) UpdateCartItemQuantity(shopId, productId, amount uint) error {
	query := c.DB.Model(&models.CartItem{}).
		Where("shop_id = ? AND product_id = ?", shopId, productId).
		UpdateColumn("quantity", amount)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func (c *CartItem) RemoveCartItem(shopId, productId, cartId uint) error {
	result := c.DB.
		Unscoped(). // ! Hard delete
		Where("shop_id = ? AND product_id = ? AND cart_id = ?", shopId, productId, cartId).
		Delete(&models.CartItem{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
