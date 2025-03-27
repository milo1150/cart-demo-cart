package schemas

import pb "github.com/milo1150/cart-demo-proto/pkg/shop_product"

type GetCartResponse struct {
	BaseModelSchema
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
