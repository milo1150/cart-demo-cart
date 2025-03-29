package schemas

import "cart-service/internal/models"

type CreateCheckoutItem struct {
	CheckoutItems CheckoutItemSlice `json:"checkout_items" validate:"required,dive"`
	// CouponId      *uint         `json:"coupon_id" validate:"omitempty,gt=0"`
}
type CheckoutItemSlice []CheckoutItem

type CheckoutItem struct {
	Shop     models.CheckoutItemShopJson         `json:"shop" validate:"required"`
	Products models.CheckoutItemProductJsonSlice `json:"product" validate:"required,dive"`
}
