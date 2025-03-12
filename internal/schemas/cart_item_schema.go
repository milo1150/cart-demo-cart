package schemas

type CreateCartItemPayload struct {
	Quantity  uint `json:"quantity" validate:"required"`
	ProductId uint `json:"product_id" validate:"required"`
	ShopId    uint `json:"shop_id" validate:"required"`
	CartId    uint `json:"cart_id" validate:"required"`
}
