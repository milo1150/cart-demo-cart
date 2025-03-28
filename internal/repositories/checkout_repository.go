package repositories

import (
	"cart-service/internal/models"
	"cart-service/internal/schemas"

	"gorm.io/gorm"
)

type Checkout struct {
	DB *gorm.DB
}

func (c *Checkout) CreateCheckout(payload *schemas.CreateCheckoutItem, userId uint) (*models.Checkout, error) {
	checkout := models.Checkout{
		CartItemInfo: payload.CartItemInfo,
		CouponId:     payload.CouponId,
		UserId:       userId,
	}

	query := c.DB.Create(&checkout)

	if query.Error != nil {
		return nil, query.Error
	}

	return &checkout, nil
}
