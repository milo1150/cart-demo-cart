package schemas

import (
	"cart-service/internal/models"
	"time"

	pb "github.com/milo1150/cart-demo-proto/pkg/payment"
)

type CreateCheckoutItem struct {
	CheckoutItems CheckoutItemSlice `json:"checkout_items" validate:"required,dive"`
	// CouponId      *uint         `json:"coupon_id" validate:"omitempty,gt=0"`
}
type CheckoutItemSlice []CheckoutItem

type CheckoutItem struct {
	Shop     models.CheckoutItemShopJson         `json:"shop" validate:"required"`
	Products models.CheckoutItemProductJsonSlice `json:"products" validate:"required,dive"`
}

type CheckoutItemResponse struct {
	ID              uint                        `json:"id"`
	CreatedAt       time.Time                   `json:"created_at"`
	UpdatedAt       time.Time                   `json:"updated_at"`
	CheckoutItems   []models.CheckoutItem       `json:"checkout_items"`
	Payment         *pb.GetPaymentOrderResponse `json:"payment"`
	TotalPaidAmount uint64                      `json:"total_paid_amount"`
}

type CheckoutItemSliceResponse struct {
	Items []CheckoutItemResponse `json:"items"`
}
