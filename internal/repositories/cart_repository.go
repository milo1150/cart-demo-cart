package repositories

import (
	"cart-service/internal/models"

	"gorm.io/gorm"
)

func CreateCart(db *gorm.DB, userId uint) error {
	newCart := models.Cart{
		UserId: userId,
	}

	if err := db.Create(&newCart).Error; err != nil {
		return err
	}

	return nil
}

func GetCart(db *gorm.DB, cartId uint) (*models.Cart, error) {
	cart := &models.Cart{}

	if err := db.Preload("CartItems").First(cart, cartId).Error; err != nil {
		return nil, err
	}

	return cart, nil
}
