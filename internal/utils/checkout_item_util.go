package utils

import (
	"cart-service/internal/schemas"
)

type CheckoutItemUtil struct{}

func (c *CheckoutItemUtil) CalculateProductTotalPaidAmount(checkoutItem *schemas.CheckoutItem) {
	for index := range checkoutItem.Products {
		checkoutItem.Products[index].PaidAmount = checkoutItem.Products[index].Price * checkoutItem.Products[index].Quantity
	}
}
