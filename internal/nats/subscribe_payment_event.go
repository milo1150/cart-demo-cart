package nats

import (
	"context"
	"time"

	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func SubscribeCreatePaymentEvent(js jetstream.JetStream, log *zap.Logger, db *gorm.DB) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cons, err := js.CreateOrUpdateConsumer(ctx, "PAYMENT_ORDER", jetstream.ConsumerConfig{
		Durable:     "CONS_PAYMENT_ORDER_CREATED",
		Description: "",
		AckPolicy:   jetstream.AckExplicitPolicy,
	})
	if err != nil {
		log.Error("Failed to create payment_order.created consumer", zap.Error(err))
	}

	cons.Consume(func(msg jetstream.Msg) {
		err := SubscribeCreatePaymentHandler(db, msg)
		if err != nil {
			log.Error("Failed to ack create payment_order.created message", zap.Error(err))
		}
		if err == nil {
			msg.Ack()
		}
	})
}
