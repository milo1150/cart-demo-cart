package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type CheckoutItem struct {
	ID              uint                         `json:"id" gorm:"primarykey"`
	CreatedAt       time.Time                    `json:"created_at"`
	UpdatedAt       time.Time                    `json:"updated_at"`
	DeletedAt       gorm.DeletedAt               `json:"deleted_at" gorm:"index"`
	Shop            CheckoutItemShopJson         `json:"shop" gorm:"type:jsonb;not null"`
	Products        CheckoutItemProductJsonSlice `json:"products" gorm:"type:jsonb;not null"`
	TotalPaidAmount uint64                       `json:"total_paid_amount"`

	// Internal
	CheckoutID uint `json:"checkout_id"`
}

func (c *CheckoutItem) BeforeCreate(tx *gorm.DB) error {
	if c.TotalPaidAmount == 0 && len(c.Products) > 0 {
		c.TotalPaidAmount = calculateCheckoutItemPaidAmount(&c.Products)
	}
	return nil
}

type CheckoutItemShopJson struct {
	Id   uint   `json:"id" validate:"required,gt=0"`
	Name string `json:"name" validate:"required"`
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (s *CheckoutItemShopJson) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal CheckoutItemShopJson JSONB value:", value))
	}

	result := CheckoutItemShopJson{}
	err := json.Unmarshal(bytes, &result)
	*s = result // mutate ref value, this line is very important
	return err
}

// Value return json value, implement driver.Valuer interface
func (s CheckoutItemShopJson) Value() (driver.Value, error) {
	if s.Id == 0 {
		return nil, nil
	}
	return json.Marshal(s)
}

type CheckoutItemProductJson struct {
	Id          uint   `json:"id" validate:"required,gt=0"`
	CreatedAt   string `json:"created_at" validate:"required"`
	UpdatedAt   string `json:"updated_at" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	ImageUrl    string `json:"image_url" validate:"required"`
	Price       uint64 `json:"price" validate:"required"`
	Stock       uint64 `json:"stock" validate:"required"`
	Quantity    uint64 `json:"quantity" validate:"required"`
	PaidAmount  uint64 `json:"paid_amount"`
}

type CheckoutItemProductJsonSlice []CheckoutItemProductJson

// Scan scan value into Jsonb, implements sql.Scanner interface
func (p *CheckoutItemProductJsonSlice) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal CheckoutItemProductJsonSlice JSONB value:", value))
	}

	result := CheckoutItemProductJsonSlice{}
	err := json.Unmarshal(bytes, &result)
	*p = result // mutate ref value, this line is very important
	return err
}

// Value return json value, implement driver.Valuer interface
func (p CheckoutItemProductJsonSlice) Value() (driver.Value, error) {
	if len(p) == 0 {
		return nil, nil
	}
	return json.Marshal(p)
}

func calculateCheckoutItemPaidAmount(items *CheckoutItemProductJsonSlice) uint64 {
	var total uint64
	for _, item := range *items {
		total += item.PaidAmount
	}
	return total
}
