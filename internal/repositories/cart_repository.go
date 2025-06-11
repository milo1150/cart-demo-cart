package repositories

import (
	"cart-service/internal/models"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Cart struct {
	DB *gorm.DB
}

func (c *Cart) CreateCart(userId uint) error {
	newCart := models.Cart{
		UserId: userId,
	}
	if err := c.DB.Create(&newCart).Error; err != nil {
		return err
	}
	return nil
}

func (c *Cart) GetCart(cartId uint) (*models.Cart, error) {
	cart := models.Cart{}
	if err := c.DB.Preload("CartItems").First(&cart, cartId).Error; err != nil {
		return nil, err
	}
	return &cart, nil
}

func (c *Cart) GetCartByUuid(cartUuid uuid.UUID) (*models.Cart, error) {
	cart := models.Cart{}
	if err := c.DB.Preload("CartItems").Where("uuid = ?", cartUuid).First(&cart).Error; err != nil {
		return nil, errors.New("cart by uuid not found")
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

func (c *Cart) GetCartIdByUserId(userId uint) (*uint, error) {
	cart := models.Cart{}
	query := c.DB.Where("user_id", userId).First(&cart)
	if query.Error != nil {
		return nil, query.Error
	}
	return &cart.ID, nil
}
