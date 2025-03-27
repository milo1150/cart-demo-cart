package schemas

type CreateCheckoutItem struct {
	CouponId     *uint               `json:"coupon_id" validate:"omitempty,gt=0"`
	CartItemInfo []CartItemInfoJsonb `json:"cart_item_info" validate:"required,dive"`
}

type CartItemInfoJsonb struct {
	Id          uint              `json:"id" validate:"required,gt=0"`
	CreatedAt   string            `json:"created_at" validate:"required"`
	UpdatedAt   string            `json:"updated_at" validate:"required"`
	Name        string            `json:"name" validate:"required"`
	Description string            `json:"description" validate:"required"`
	ImageUrl    string            `json:"image_url" validate:"required"`
	Price       uint64            `json:"price" validate:"required"`
	Stock       uint64            `json:"stock" validate:"required"`
	Quantity    uint64            `json:"quantity" validate:"required"`
	Shop        cartItemShopJsonb `json:"shop" validate:"required"`
}

type cartItemShopJsonb struct {
	Id   uint   `json:"id" validate:"required,gt=0"`
	Name string `json:"name" validate:"required"`
}
