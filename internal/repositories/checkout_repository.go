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
	checkout := models.Checkout{UserId: userId}

	// Transaction Checkout and CheckoutItem
	txErr := c.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&checkout).Error; err != nil {
			return err
		}

		checkoutItemRepository := CheckoutItem{DB: tx}
		for _, checkoutItem := range payload.CheckoutItems {
			_, err := checkoutItemRepository.CreateCheckoutItem(&checkoutItem, checkout.ID)
			if err != nil {
				return err
			}
		}

		return nil
	})

	// Return transaction error
	if txErr != nil {
		return nil, txErr
	}

	return &checkout, nil
}

func (c *Checkout) GetCheckout(userId, checkoutId uint) (*models.Checkout, error) {
	result := models.Checkout{}

	query := c.DB.Preload("CheckoutItems").Where("user_id = ? AND id = ?", userId, checkoutId).Find(&result)

	if query.Error != nil {
		return nil, query.Error
	}

	return &result, nil
}

func (c *Checkout) GetCheckouts(userId uint) (*[]models.Checkout, error) {
	result := []models.Checkout{}

	query := c.DB.Preload("CheckoutItems").Where("user_id = ?", userId).Find(&result)

	if query.Error != nil {
		return nil, query.Error
	}

	return &result, nil
}
