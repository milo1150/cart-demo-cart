package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	Uuid uuid.UUID `gorm:"not null;type:uuid;unique;index"`

	UserId    uint `gorm:"not null"`
	CartItems []CartItem
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
