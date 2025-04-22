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

	streamConn := make(chan bool)
	ticker := time.NewTicker(2 * time.Second)
	var cons jetstream.Consumer

	go func() {
		defer ticker.Stop()

		for {
			select {
			case tick := <-ticker.C:
				// Create consumer
				log.Info("retry create PAYMENT_ORDER consumer", zap.Time("tick", tick))
				consumer, err := js.CreateOrUpdateConsumer(ctx, "PAYMENT_ORDER", jetstream.ConsumerConfig{
					Durable:     "CONS_PAYMENT_ORDER_CREATED",
					Description: "",
					AckPolicy:   jetstream.AckExplicitPolicy,
				})

				// Log error if failed to create consumer
				if err != nil {
					log.Error("Failed to create payment_order.created consumer", zap.Error(err))
					continue
				}

				// Create consumer is ok
				log.Info("create PAYMENT_ORDER consumer OK")
				cons = consumer
				close(streamConn)
				return

			case <-streamConn:
				return
			}
		}
	}()

	<-streamConn

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
