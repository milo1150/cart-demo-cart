package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CheckoutItem struct {
	gorm.Model
	Uuid         uuid.UUID    `gorm:"not null;type:uuid;unique;index"`
	CartItemInfo CartItemJson `gorm:"type:jsonb;not null"`

	// External relation
	UserId    uint
	CouponId  *uint
	PaymentId uint
}

func (c *CheckoutItem) BeforeCreate(tx *gorm.DB) error {
	if c.Uuid == uuid.Nil {
		uuidV7, err := uuid.NewV7()
		if err != nil {
			return err
		}
		c.Uuid = uuidV7
	}
	return nil
}
