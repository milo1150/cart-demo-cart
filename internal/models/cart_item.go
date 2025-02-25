package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartItem struct {
	gorm.Model
	Uuid     uuid.UUID `gorm:"not null;type:uuid;unique;index"`
	Name     string
	Price    float32
	Quantity uint

	// Internal relation
	CartID uint

	// External relation
	ProductId uint
	ShopId    uint
}

func (c *CartItem) BeforeCreate(tx *gorm.DB) error {
	if c.Uuid == uuid.Nil {
		uuidV7, err := uuid.NewV7()
		if err != nil {
			return err
		}
		c.Uuid = uuidV7
	}
	return nil
}
