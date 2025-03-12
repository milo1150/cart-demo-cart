package repositories

import (
	"cart-service/internal/models"
	"cart-service/internal/schemas"

	"gorm.io/gorm"
)

func CreateCartItem(db *gorm.DB, payload schemas.CreateCartItemPayload) error {
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
