package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Checkout struct {
	ID           uint              `json:"id" gorm:"primarykey"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    time.Time         `json:"updated_at"`
	DeletedAt    gorm.DeletedAt    `json:"deleted_at" gorm:"index"`
	Uuid         uuid.UUID         `json:"uuid" gorm:"not null;type:uuid;unique;index"`
	CartItemInfo CartItemInfoSlice `json:"cart_item_info" gorm:"type:jsonb;not null"`

	// External relation
	UserId    uint  `json:"user_id"`
	CouponId  *uint `json:"coupon_id"`
	PaymentId uint  `json:"payment_id"` // TODO: Queue for binding
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

type CartItemInfoJson struct {
	Id          uint             `json:"id" validate:"required,gt=0"`
	CreatedAt   string           `json:"created_at" validate:"required"`
	UpdatedAt   string           `json:"updated_at" validate:"required"`
	Name        string           `json:"name" validate:"required"`
	Description string           `json:"description" validate:"required"`
	ImageUrl    string           `json:"image_url" validate:"required"`
	Price       uint64           `json:"price" validate:"required"`
	Stock       uint64           `json:"stock" validate:"required"`
	Quantity    uint64           `json:"quantity" validate:"required"`
	Shop        cartItemShopJson `json:"shop" validate:"required"`
}

type cartItemShopJson struct {
	Id   uint   `json:"id" validate:"required,gt=0"`
	Name string `json:"name" validate:"required"`
}

type CartItemInfoSlice []CartItemInfoJson

// Scan scan value into Jsonb, implements sql.Scanner interface
func (c *CartItemInfoSlice) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal CartItemInfoSlice JSONB value:", value))
	}
	return json.Unmarshal(bytes, c)
}

// Value return json value, implement driver.Valuer interface
func (c CartItemInfoSlice) Value() (driver.Value, error) {
	if len(c) == 0 {
		return nil, nil
	}
	return json.Marshal(c)
}
