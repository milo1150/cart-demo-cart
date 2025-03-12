package schemas

import pb "github.com/milo1150/cart-demo-proto/pkg/shop_product"

type GetCartResponse struct {
	BaseModelSchema
	UserId    uint       `json:"user_id"`
	CartItems []CartItem `json:"cart_items"`
}

type CartItem struct {
	BaseModelSchema
	Quantity uint                   `json:"quantity"`
	CartID   uint                   `json:"cart_id"`
	Products *pb.GetProductResponse `json:"product"`
	ShopId   uint                   `json:"shop_id"`
}
