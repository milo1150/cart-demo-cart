package database

import (
	"cart-service/internal/models"

	"gorm.io/gorm"
)

func migrate_cart_items_202503120840_drop_columns(db *gorm.DB) {
	if db.Migrator().HasColumn(&models.CartItem{}, "Name") {
		db.Migrator().DropColumn(&models.CartItem{}, "Name")
	}

	if db.Migrator().HasColumn(&models.CartItem{}, "Price") {
		db.Migrator().DropColumn(&models.CartItem{}, "Price")
	}

	if db.Migrator().HasColumn(&models.CartItem{}, "Quantity") {
		db.Migrator().DropColumn(&models.CartItem{}, "Quantity")
	}
}
