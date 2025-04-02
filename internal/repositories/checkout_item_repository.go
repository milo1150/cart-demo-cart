package repositories

import (
	"cart-service/internal/models"
	"cart-service/internal/schemas"

	"gorm.io/gorm"
)

type CheckoutItem struct {
	DB *gorm.DB
}

func (c *CheckoutItem) CreateCheckoutItem(payload *schemas.CheckoutItem, checkoutId uint) (*models.CheckoutItem, error) {
	checkoutItem := models.CheckoutItem{
		Shop:       payload.Shop,
		Products:   payload.Products,
		CheckoutID: checkoutId,
	}

	if err := c.DB.Create(&checkoutItem).Error; err != nil {
		return nil, err
	}

	return &checkoutItem, nil
}

func (c *CheckoutItem) GetCheckoutItem(id uint) (*models.CheckoutItem, error) {
	result := models.CheckoutItem{}

	query := c.DB.Where("id = ?", id).Find(&result)

	if query.Error != nil {
		return nil, query.Error
	}

	return &result, nil
}
