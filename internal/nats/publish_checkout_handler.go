package nats

import (
	"cart-service/internal/models"
	"context"
	"encoding/json"

	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/zap"
)

func PublishCreateCheckoutHandler(js jetstream.JetStream, ctx context.Context, checkout *models.Checkout, log *zap.Logger) {
	// Parse publish message
	publishPayload := map[string]uint{
		"checkout_id": checkout.ID,
		"user_id":     checkout.UserId,
	}
	data, err := json.Marshal(publishPayload)
	if err != nil {
		log.Error("Failed to parse checkout.created message", zap.Error(err))
	}

	// Publish message to payment service
	_, err = js.Publish(ctx, "checkout.created", data)
	if err != nil {
		log.Error("Failed to publish checkout.created message", zap.Error(err))
	}
}
