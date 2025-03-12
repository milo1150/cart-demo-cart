package repositories

import (
	"cart-service/internal/models"
	"cart-service/internal/schemas"

	"gorm.io/gorm"
)

func CreateCartItem(db *gorm.DB, payload schemas.AddCartItemPayload) error {
	newCartItem := models.CartItem{
		Quantity:  payload.Quantity,
		CartID:    payload.CartId,
		ProductId: payload.ProductId,
		ShopId:    payload.ShopId,
	}

	if err := db.Create(&newCartItem).Error; err != nil {
		return err
	}

	return nil
}

func CartItemExists(db *gorm.DB, shopId uint, productId uint) (bool, error) {
	cartItem := &models.CartItem{}

	query := db.Debug().Where("shop_id = ? AND product_id = ?", shopId, productId).Find(cartItem)
	if query.Error != nil {
		return false, query.Error
	}

	if cartItem.CartID == 0 {
		return false, nil
	}

	return true, nil
}

func FindCartItem(db *gorm.DB, shopId uint, productId uint) (*models.CartItem, error) {
	cartItem := &models.CartItem{}

	query := db.Debug().Where("shop_id = ? AND product_id = ?", shopId, productId).First(cartItem)
	if query.Error != nil {
		return nil, query.Error
	}

	return cartItem, nil
}

func UpdateCartItemQuantity(db *gorm.DB, shopId uint, productId uint, amount uint) error {
	query := db.Model(&models.CartItem{}).Debug().
		Where("shop_id = ? AND product_id = ?", shopId, productId).
		UpdateColumn("quantity", amount)
	if query.Error != nil {
		return query.Error
	}
	return nil
}
