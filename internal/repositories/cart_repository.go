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
