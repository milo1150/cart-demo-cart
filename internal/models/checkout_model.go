package models

import (
	"cart-service/internal/schemas"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Checkout struct {
	gorm.Model
	Uuid         uuid.UUID                 `gorm:"not null;type:uuid;unique;index"`
	CartItemInfo schemas.CartItemInfoJsonb `gorm:"type:jsonb;not null"`

	// External relation
	UserId    uint
	CouponId  *uint
	PaymentId uint // TODO: Queue for binding
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
