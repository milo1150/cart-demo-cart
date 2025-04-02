package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Checkout struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Uuid      uuid.UUID      `json:"uuid" gorm:"not null;type:uuid;unique;index"`

	// Internal
	CheckoutItems []CheckoutItem `json:"checkout_items"`

	// External
	UserId    uint `json:"-"`
	PaymentId uint `json:"payment_id"`
}

func (c *Checkout) BeforeCreate(tx *gorm.DB) error {
	if c.Uuid == uuid.Nil {
		uuidV7, err := uuid.NewV7()
		if err != nil {
			return err
		}
		c.Uuid = uuidV7
	}
	return nil
}
