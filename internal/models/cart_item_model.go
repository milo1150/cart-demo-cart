package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartItem struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Uuid      uuid.UUID      `json:"uuid" gorm:"not null;type:uuid;unique;index"`

	Quantity uint `json:"quantity"`

	// Internal relation
	CartID uint `json:"cart_id"`

	// External relation
	ProductId uint `json:"product_id"`
	ShopId    uint `json:"shop_id"`
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
