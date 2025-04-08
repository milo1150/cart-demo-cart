package dto

import (
	"cart-service/internal/models"
	"cart-service/internal/schemas"

	pb "github.com/milo1150/cart-demo-proto/pkg/shop_product"
	"github.com/samber/lo"
)

func TransformCartDetail(cart models.Cart, products []*pb.GetProductResponse) schemas.GetCartResponse {
	result := schemas.GetCartResponse{
		CreatedAt: cart.CreatedAt,
		UpdatedAt: cart.UpdatedAt,
		Uuid:      cart.Uuid,
		UserId:    cart.UserId,
		CartItems: []schemas.CartItemResponse{},
	}

	hashProducts := lo.KeyBy(products, func(product *pb.GetProductResponse) uint64 {
		return product.Id
	})

	for _, cartItem := range cart.CartItems {
		productDetail := hashProducts[uint64(cartItem.ProductId)]
		cartItemDto := TransformCartItemDetail(cartItem, productDetail)
		result.CartItems = append(result.CartItems, cartItemDto)
	}

	return result
}

func TransformCartItemDetail(cartItem models.CartItem, productDetail *pb.GetProductResponse) schemas.CartItemResponse {
	result := schemas.CartItemResponse{
		BaseModelSchema: schemas.BaseModelSchema{
			ID:        cartItem.ID,
			CreatedAt: cartItem.CreatedAt,
			UpdatedAt: cartItem.UpdatedAt,
			Uuid:      cartItem.Uuid,
		},
		Quantity: cartItem.Quantity,
		CartID:   cartItem.CartID,
		ShopId:   cartItem.ShopId,
		Product:  productDetail,
	}
	return result
}
