package schemas

type CreateCheckoutItem struct {
	Counpon      uint `json:"coupon_id" validate:"omitempty"`
	CartItemInfo []product
}

type product struct {
	Id          uint
	CreatedAt   string
	UpdatedAt   string
	Name        string
	Description string
	ImageUrl    string
	Price       uint64
	Stock       uint64
	quantity    uint64
}
