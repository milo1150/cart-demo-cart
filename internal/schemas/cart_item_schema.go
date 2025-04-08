package schemas

type AddCartItemPayload struct {
	Quantity  uint `json:"quantity" validate:"required"`
	ProductId uint `json:"product_id" validate:"required"`
	ShopId    uint `json:"shop_id" validate:"required"`
}

type AddCartItemSlicesPayload struct {
	CartItems []AddCartItemPayload `json:"cart_items" validate:"required,dive"`
}
