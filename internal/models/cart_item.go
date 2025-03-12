package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartItem struct {
	gorm.Model
	Uuid uuid.UUID `gorm:"not null;type:uuid;unique;index"`

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

type CartItemJson struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Uuid      uuid.UUID `json:"uuid"`
	Name      string    `json:"name"`
	Price     float32   `json:"price"`
	Quantity  uint      `json:"quantity"`
	CartId    uint      `json:"cart_id"`
	ProductId uint      `json:"product_id"`
	ShopId    uint      `json:"shop_id"`
}
