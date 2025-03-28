package schemas

import "cart-service/internal/models"

type CreateCheckoutItem struct {
	CouponId     *uint                    `json:"coupon_id" validate:"omitempty,gt=0"`
	CartItemInfo models.CartItemInfoSlice `json:"cart_item_info" validate:"required,dive"`
}
