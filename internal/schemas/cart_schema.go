package schemas

import (
	"time"

	"github.com/google/uuid"
	pb "github.com/milo1150/cart-demo-proto/pkg/shop_product"
)

type GetCartResponse struct {
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	Uuid      uuid.UUID          `json:"uuid"`
	UserId    uint               `json:"user_id"`
	CartItems []CartItemResponse `json:"cart_items"`
}

type CartItemResponse struct {
	BaseModelSchema
	Quantity uint                   `json:"quantity"`
	CartID   uint                   `json:"cart_id"`
	Product  *pb.GetProductResponse `json:"product"`
	ShopId   uint                   `json:"shop_id"`
}
