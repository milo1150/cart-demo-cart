package repositories

import (
	"cart-service/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Cart struct {
	DB *gorm.DB
}

func (c *Cart) CreateCart(db *gorm.DB, userId uint) error {
	newCart := models.Cart{
		UserId: userId,
	}

	if err := db.Create(&newCart).Error; err != nil {
		return err
	}

	return nil
}

func (c *Cart) GetCart(db *gorm.DB, cartId uint) (*models.Cart, error) {
	cart := models.Cart{}

	if err := db.Preload("CartItems").First(&cart, cartId).Error; err != nil {
		return nil, err
	}

	return &cart, nil
}

func (c *Cart) GetCartByUuid(db *gorm.DB, cartUuid uuid.UUID) (*models.Cart, error) {
	cart := models.Cart{}

	if err := db.Preload("CartItems").Where("uuid = ?", cartUuid).First(&cart).Error; err != nil {
		return nil, err
	}

	return &cart, nil
}

func (c *Cart) GetCartUuidByUserId(userId uint) (*uuid.UUID, error) {
	cart := models.Cart{}
	query := c.DB.Where("user_id = ?", userId).First(&cart)
	if query.Error != nil {
		return nil, query.Error
	}
	return &cart.Uuid, nil
}
