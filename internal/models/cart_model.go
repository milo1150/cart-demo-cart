package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Cart struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Uuid      uuid.UUID      `json:"uuid" gorm:"not null;type:uuid;unique;index"`

	UserId    uint       `json:"user_id" gorm:"not null"`
	CartItems []CartItem `json:"cart_items"`
}

func (c *Cart) BeforeCreate(tx *gorm.DB) error {
	if c.Uuid == uuid.Nil {
		uuidV7, err := uuid.NewV7()
		if err != nil {
			return err
		}
		c.Uuid = uuidV7
	}
	return nil
}
