package nats

import (
	"cart-service/internal/repositories"
	"encoding/json"
	"fmt"

	schemas "github.com/milo1150/cart-demo-payment/pkg/schemas"
	"github.com/nats-io/nats.go/jetstream"
	"gorm.io/gorm"
)

func SubscribeCreatePaymentHandler(db *gorm.DB, msg jetstream.Msg) error {
	payload := schemas.PublishCreatedPaymentOrderPayload{}
	if err := json.Unmarshal(msg.Data(), &payload); err != nil {
		return fmt.Errorf("Failed to parse payment_order.created message payload: %w", err)
	}

	rc := repositories.Checkout{DB: db}
	if err := rc.UpdateCheckoutPaymentId(payload.CheckoutId, payload.PaymentId); err != nil {
		return err
	}

	return nil
}
