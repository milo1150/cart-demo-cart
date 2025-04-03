package dto

import (
	"cart-service/internal/models"
	"cart-service/internal/schemas"

	pb "github.com/milo1150/cart-demo-proto/pkg/payment"
	"github.com/samber/lo"
)

func TransformCheckout(model models.Checkout, payment *pb.GetPaymentOrderResponse) schemas.CheckoutItemResponse {
	data := schemas.CheckoutItemResponse{
		// Checkout: model,
		ID:              model.ID,
		CreatedAt:       model.CreatedAt,
		UpdatedAt:       model.CreatedAt,
		CheckoutItems:   model.CheckoutItems,
		Payment:         payment,
		TotalPaidAmount: model.TotalPaidAmount,
	}
	return data
}

func TransformCheckoutSlice(models []models.Checkout, payment []*pb.GetPaymentOrderResponse) schemas.CheckoutItemSliceResponse {
	datas := schemas.CheckoutItemSliceResponse{
		Items: []schemas.CheckoutItemResponse{},
	}

	hashPayment := lo.KeyBy(payment, func(payment *pb.GetPaymentOrderResponse) uint {
		return uint(payment.Id)
	})

	for _, model := range models {
		datas.Items = append(datas.Items, TransformCheckout(model, hashPayment[model.PaymentId]))
	}

	return datas
}
