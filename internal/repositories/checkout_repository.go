package repositories

import (
	"cart-service/internal/models"
	"cart-service/internal/schemas"
	"cart-service/internal/utils"
	"fmt"

	"gorm.io/gorm"
)

type Checkout struct {
	DB *gorm.DB
}

func (c *Checkout) CreateCheckout(payload *schemas.CreateCheckoutItem, userId uint) (*models.Checkout, error) {
	checkout := models.Checkout{
		UserId: userId,
	}

	txErr := c.DB.Transaction(func(tx *gorm.DB) error {
		// Create checkout
		if err := tx.Create(&checkout).Error; err != nil {
			return err
		}

		// Create checkout_item
		checkoutItemRepository := CheckoutItem{DB: tx}
		for _, checkoutItem := range payload.CheckoutItems {
			uc := utils.CheckoutItemUtil{}
			uc.CalculateProductTotalPaidAmount(&checkoutItem) // Calculate and mutate total price
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

func (c *Checkout) UpdateCheckoutTotalPaidAmount(checkoutId, userId uint) error {
	checkout, err := c.GetCheckout(userId, checkoutId)
	if err != nil {
		return err
	}

	total := models.CalculateCheckoutTotalPaidAmount(&checkout.CheckoutItems)

	result := c.DB.Model(&models.Checkout{}).Where(checkout.ID).Update("total_paid_amount", total)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("failed to update total_paid_amount: %w", err)
	}

	return nil
}

func (c *Checkout) GetCheckout(userId, checkoutId uint) (*models.Checkout, error) {
	result := models.Checkout{}

	query := c.DB.Preload("CheckoutItems").Where("user_id = ? AND id = ?", userId, checkoutId).First(&result)

	if query.Error != nil {
		return nil, query.Error
	}

	return &result, nil
}

func (c *Checkout) GetCheckouts(userId uint) (*[]models.Checkout, error) {
	result := []models.Checkout{}

	query := c.DB.Preload("CheckoutItems").Order("updated_at DESC").Where("user_id = ?", userId).Find(&result)

	if query.Error != nil {
		return nil, query.Error
	}

	return &result, nil
}

func (c *Checkout) UpdateCheckoutPaymentId(checkoutid, paymentId uint) error {
	result := c.DB.Model(&models.Checkout{}).Where("id = ?", checkoutid).Update("payment_id", paymentId)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no checkout record updated")
	}

	return nil
}
